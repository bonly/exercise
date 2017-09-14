// golang-searchable-scrollback: interactively search in a command's output.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
)

var (
	debug            = flag.Bool("debug", false, "enable debugging output")
	trimLongLines    = flag.Bool("trimLongLines", false, "trim lines wider than terminal")
	showOmittedLines = flag.Bool("showOmittedLines", true, "draw a [...] when lines are omitted")
	showLineNumbers  = flag.Bool("showLineNumbers", false, "add a prefix for line numbers")
)

const size = 9999 // number of lines in the scrollback buffer.

var prefixWidth int
var ellipsisPrefix string

// simple circular buffer
type Circular struct {
	x [size]string
	i int
}

func (c *Circular) Add(s string) {
	c.x[c.i%size] = s
	c.i += 1
}

func (c *Circular) Get() []string {
	if c.i < size {
		return c.x[:c.i]
	}
	i := c.i % size
	return append(c.x[i:], c.x[:i]...)
}

func (c *Circular) WriteTo(w io.Writer) {
	//io.WriteString(w, "\n\nBeginning of searchable scrollback:\n")
	clear()
	for _, v := range c.Get() {
		io.WriteString(w, v)
	}
}

func (c *Circular) WriteIfContains(w io.Writer, s string, scrollOffset int) {
	clear()
	count := 0
	var lastMatched int

	var outLines []string

	for i, v := range c.Get() {
		ok, err := regexp.MatchString("(?i:"+s+")", v)
		if err != nil {
			log.Print(err)
			continue
		}
		if ok {
			if lastMatched > 0 {
				outLines = append(outLines, ellipsisPrefix+"["+strings.Repeat(".", lastMatched)+"]\n")
			}
			if *showLineNumbers {
				outLines = append(outLines, fmt.Sprintf("%4d: %s", i, v))
			} else {
				outLines = append(outLines, v)
			}
			count += 1
			lastMatched = 0
		} else {
			lastMatched += 1
		}
	}

	for i, v := range outLines {
		if scrollOffset+i >= len(outLines) {
			break
		}
		io.WriteString(os.Stdout, v)
	}

	fmt.Fprintf(os.Stdout, "[%d lines matched]\n", count)
}

func outReader(rr io.Reader, ch chan string) {
	r := bufio.NewReader(rr)
	for {
		x, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				runtime.Goexit()
			}
			log.Fatal(err)
		}
		ch <- x
	}
}

var b = Circular{}

type winsize struct {
	ws_row, ws_col       uint16
	ws_xpixel, ws_ypixel uint16
}

func getwinsize() winsize {
	ws := winsize{}
	syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(0), uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&ws)))
	return ws
}

func termSize() (rows, cols uint16) {
	ws := getwinsize()
	return ws.ws_row, ws.ws_col
}

func main() {
	flag.Parse()

	if *showLineNumbers {
		prefixWidth = 4 + 2 // (len(size) + 2)
		ellipsisPrefix = "       "
	} else {
		ellipsisPrefix = "  "
	}

	args := flag.Args()

	cmd := exec.Command(args[0], args[1:]...)
	p1, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	p2, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	lines := make(chan string)
	chars := make(chan byte)

	go outReader(p1, lines)
	go outReader(p2, lines)

	//rawtty.Raw()
	rawttyCbreak()
	defer rawttyCook()

	go func() {
		for {
			c := rawttyGetCh()
			chars <- c
		}
	}()

	cmd.Start()

	for {
		select {
		case c := <-chars:
			switch c {
			case '\n':
				fmt.Println("scrollback-search is running. To get started, type the letters you want to search for.")
			default:
				searchMode(c, chars)
			}
		case l := <-lines:
			if *trimLongLines {
				_, colsu := termSize()
				cols := int(colsu)
				if cols > 10 && len(l) > cols {
					l = l[:cols-6-prefixWidth] + " [...]\n"
				}
			}

			io.WriteString(os.Stdout, l)
			b.Add(l)
		}
	}
}

func out(s string) {
	io.WriteString(os.Stdout, s)
}

func clear() {
	out(`[2J`)
}
func searchMode(c byte, ch chan byte) {
	var s []byte
	var scrollOffset int

	print := func() {
		b.WriteIfContains(os.Stdout, string(s), scrollOffset)
		out("\nScrollback search: ")
		os.Stdout.Write(s)
	}
	s = append(s, c)
	print()
	for {
		c := <-ch
		switch c {
		case 8, 127: // should catch all of ^H, Del, Backspace
			if len(s) > 0 {
				s = s[:len(s)-1]
			}
		case '', '\n':
			b.WriteTo(os.Stdout)
			return
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'L', 'M', 'N', 'O',
			'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
		case 'K':
			scrollOffset += 1
		case 'J':
			scrollOffset -= 1
			if scrollOffset < 0 {
				scrollOffset = 0
			}
		default:
			s = append(s, c)
		}
		print()
	}
}

func init() {
	log.SetPrefix(os.Args[0] + ":")
	log.SetFlags(log.Lshortfile)
}

func usage() {
	fmt.Fprintln(os.Stderr, os.Args[0])
	fmt.Fprintln(os.Stderr)
	fmt.Fprintf(os.Stderr, "USAGE:\n  %s\n", os.Args[0])
	flag.PrintDefaults()
}

func ttyCmd(arg string) error {
	cmd := exec.Command("stty", arg)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

// cbreak mode allows getch and ^C, but doesnt alter \r behavior
func rawttyCbreak() error {
	return ttyCmd("cbreak")
}

func rawttyCook() error {
	return ttyCmd("cooked")
}

func rawttyGetCh() byte {
	b := make([]byte, 1)

	os.Stdin.Read(b)
	return b[0]
}

func exitUsage() {
	usage()
	os.Exit(1)
}