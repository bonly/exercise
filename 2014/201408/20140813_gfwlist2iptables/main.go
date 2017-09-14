package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"text/template"
)

func main() {
	gfwlist := getGFWList();
	content := decodeGFWList(gfwlist);
	domains := parseGFWList(content);
	iplist := resolveDomains(domains);
	generate(iplist, "./gfw-iptable.sh");
}

func getGFWList() []byte {
	// resp, err := http.Get("http://autoproxy-gfwlist.googlecode.com/svn/trunk/gfwlist.txt")
	resp, err := http.Get("https://raw.githubusercontent.com/gfwlist/gfwlist/master/gfwlist.txt");
	defer resp.Body.Close();
	if err != nil {
		log.Fatal(err);
	}
	gfwlist, _ := ioutil.ReadAll(resp.Body);
	return gfwlist;
}

func decodeGFWList(gfwlist []byte) []byte {
	result, _ := base64.StdEncoding.DecodeString(string(gfwlist));
	return result;
}

func parseGFWList(content []byte) []string {
	var domains []string;

	scanner := bufio.NewScanner(bytes.NewReader(content));
	for scanner.Scan() {
		text := scanner.Text();
		regComment := regexp.MustCompile("(^!|^@@|^\\[|!--)[\\s\\w-:,+\\.\\]\\*#\\(\\)\\|/@]*");
		textIgnoreComment := regComment.ReplaceAllString(text, "");
		if len(textIgnoreComment) == 0 {
			continue;
		}

		regDomainSpec := regexp.MustCompile("^\\|\\||^http(s?)://|^\\|http(s?)://|^\\.|^\\w+\\*\\.|(\\.?)[/*]\\S+|^%\\S+|/|^\\w+$");
		domainSpec := regDomainSpec.ReplaceAllString(textIgnoreComment, "");

		if strings.ContainsAny(domainSpec, ".") {
			domains = append(domains, domainSpec);
		}
	}

	removeDuplicate(&domains);

	return domains;
}

func removeDuplicate(slis *[]string) {
	found := make(map[string]bool);
	j := 0;
	for i, val := range *slis {
		if _, ok := found[val]; !ok {
			found[val] = true;
			(*slis)[j] = (*slis)[i];
			j++;
		}
	}
	*slis = (*slis)[:j];
}

func resolveDomains(domains []string) []string {
	var iplist []string;

	const pool int = 50;
	domainSenter := make(chan string, pool);
	subDone := make(chan bool, pool);
	done := make(chan bool);

	go func() {
		var subDoneCount int;
		totalDoneCount := len(domains);
		for i, domain := range domains {
			if i >= pool {
				<-subDone;
				subDoneCount++;
			}
			domainSenter <- domain;
			log.Printf("开始解析第 %d 个，总共 %d 个，已解析 %d 个", i+1, totalDoneCount, subDoneCount);
		}
		close(domainSenter);
		for i := 0; i < totalDoneCount-subDoneCount; i++ {
			<-subDone;
			log.Printf("等待解析完成，还剩 %d 个", totalDoneCount-subDoneCount-i-1);
		}
		done <- true;
	}()

	go func() {
		for {
			domain, more := <-domainSenter;
			if more {
				go lookupHost(domain, subDone, &iplist);
			} else {
				return;
			}
		}
	}()

	<-done;
	removeDuplicate(&iplist);
	return iplist;
}

func lookupHost(host string, domainRes chan<- bool, iplist *[]string) {
	addrs, err := net.LookupHost(host);
	if err != nil {
		log.Println(err);
	}
	for _, ip := range addrs {
		*iplist = append(*iplist, ip);
	}
	domainRes <- true;
}

func generate(iplist []string, filename string) {
	t := template.Must(template.ParseFiles("template.tpl"));
	var b bytes.Buffer;
	t.ExecuteTemplate(&b, "template.tpl", iplist);
	ioutil.WriteFile(filename, b.Bytes(), 0755);
}