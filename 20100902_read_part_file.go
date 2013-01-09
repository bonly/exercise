package main

import (
    "os"
    "bufio"
    "bytes"
    "fmt"
)

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err os.Error) {
    var (
        file *os.File
        part []byte
        prefix bool
    )
    if file, err = os.Open(path); err != nil {
        return
    }
    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 1024))
    for {
        if part, prefix, err = reader.ReadLine(); err != nil {
            break
        }
        buffer.Write(part)
        if !prefix {
            lines = append(lines, buffer.String())
            buffer.Reset()
        }
    }
    if err == os.EOF {
        err = nil
    }
    return
}

func main() {
    lines, err := readLines("foo.txt")
    if err != nil {
        fmt.Println("Error: %s\n", err)
        return
    }
    for _, line := range lines {
        fmt.Println(line)
    }
}