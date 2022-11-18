package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
	urlRoot = "http://ruliweb.com"
)

var wg sync.WaitGroup

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func parseMainNodes(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent != nil {
		return scrape.Attr(n.Parent, "class") == "row"
	}
	return false
}

func scrapContents(url string, filename string) {
	defer wg.Done()

	resp, err := http.Get(urlRoot)
	errCheck(err)

	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)
	errCheck(err)

	// Response data
	matchNode := func(n *html.Node) bool {
		return n.DataAtom == atom.A && scrape.Attr(n, "class") == "deco"
	}

	file, err := os.OpenFile("files/"+filename+".txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	errCheck(err)

	defer file.Close()

	w := bufio.NewWriter(file)

	for _, g := range scrape.FindAll(root, matchNode) {
		fmt.Println(scrape.Text(g))
		w.WriteString(scrape.Text(g) + "\r\n")
	}
	w.Flush()
}

func main() {
	resp, err := http.Get(urlRoot)
	errCheck(err)

	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)
	errCheck(err)

	//Parse main Nodes
	urlList := scrape.FindAll(root, parseMainNodes)

	for _, link := range urlList {
		// fmt.Println(link)
		// fmt.Println("target URL", scrape.Attr(link, "href"))
		fileName := strings.Replace(scrape.Attr(link, "href"), "https://bbs.ruliweb.com/family/", "", 1)

		// fmt.Println(fileName)
		wg.Add(1)
		// GOROUTINE start

		go scrapContents(scrape.Attr(link, "href"), fileName)
	}
	wg.Wait()
}
