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

func AocInput(year int, day int) string {
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

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Failed reading %v: %v", url, err))
	}

	return string(b)
}

var newlinesRegexp = regexp.MustCompile(`\n+`)

func AocInputParagraphs(year int, day int) [][]string {
	var paragraphs [][]string
	input := strings.TrimSpace(AocInput(year, day))
	for _, paragraph := range newlinesRegexp.Split(input, -1) {
		if len(paragraph) > 0 {
			paragraphs = append(paragraphs, strings.Split(paragraph, "\n"))
		}
	}
	if len(paragraphs) == 0 {
		panic("No paragraphs found")
	}
	return paragraphs
}

func AocInputLines(year int, day int) []string {
	var lines []string
	input := strings.TrimSpace(AocInput(year, day))
	for _, line := range newlinesRegexp.Split(input, -1) {
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}
	if len(lines) == 0 {
		panic("No lines found")
	}
	return lines
}
