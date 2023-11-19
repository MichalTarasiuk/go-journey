package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

var sessionPath = path.Join(os.Getenv("HOME"), ".advent-of-code-session")

func aocInput(year int, day int) string {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(fmt.Sprintf("Failed creating request to %v: %v", url, err))
	}

	session, err := ioutil.ReadFile(sessionPath)
	if err != nil {
		panic(fmt.Sprint("Failed reading session ID: ", err))
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: strings.TrimSpace(string(session))})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(fmt.Sprintf("Failed requesting %v: %v", url, err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Getting %v failed: %v", url, resp.Status))
	}

	if err != nil {
		panic(fmt.Sprintf("Failed reading %v: %v", url, err))
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Failed reading %v: %v", url, err))
	}

	return string(b)
}

var newlinesRegexp = regexp.MustCompile(`\n\n+`)

func AocInputParagraphs(year int, day int) [][]string {
	var pgs [][]string
	all := strings.Trim(aocInput(year, day), "\n")
	for _, pg := range newlinesRegexp.Split(all, -1) {
		if len(pg) > 0 {
			pgs = append(pgs, strings.Split(pg, "\n"))
		}
	}
	if len(pgs) == 0 {
		panic("No paragraphs found")
	}
	return pgs
}
