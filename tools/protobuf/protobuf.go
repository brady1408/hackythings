package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path := flag.String("path", "", "Path to the file you want to parse")
	typeMatching := flag.Bool("typematch", false, "Show duplicate types in output")
	typeSort := flag.Bool("typesort", false, "Sort the proto and write to out.proto")

	flag.Parse()

	if *path != "" {
		if *typeSort {
			sortProto(*path)
		}
		if *typeMatching {
			findDupTypes(*path)
		}
	} else {
		fmt.Print("You must include the -path parameter")
	}

}

func findDupTypes(path string) {
	dat, err := ioutil.ReadFile(path)
	check(err)

	file := string(dat)
	// Clean file
	if strings.Index(file, "\r\n") != -1 {
		file = strings.ReplaceAll(file, "\r\n", "\n")
	}
	if strings.Index(file, "\r") != -1 {
		file = strings.ReplaceAll(file, "\r", "\n")
	}
	// File should be clean with only \n EOL delimiters
	file = strings.ReplaceAll(file, "\nmessage", "\n<replaceMe>message")
	file = strings.ReplaceAll(file, "\nenum", "\n<replaceMe>enum")
	fileSlice := strings.Split(file, "<replaceMe>")
	fileSlice = fileSlice[1:] // get rid of header
	sort.Slice(fileSlice, func(i int, j int) bool { return fileSlice[i] < fileSlice[j] })
	m := make(map[string]string)
	dup := []string{}
	for _, value := range fileSlice {
		s := strings.Split(value, "\n")
		s = s[1 : len(s)-2]
		var index string
		for _, v := range s {
			index = index + v + "\n"
		}
		if _, ok := m[index]; !ok {
			m[index] = value
		} else {
			if strings.Index(value, "currency") != -1 {
				dup = append(dup, value)
			}
		}
	}
	fmt.Printf("%v", dup)
}

func sortProto(path string) {
	dat, err := ioutil.ReadFile(path)
	check(err)

	file := string(dat)
	file = strings.ReplaceAll(file, "\nmessage", "\n<replaceMe>message")
	file = strings.ReplaceAll(file, "\nenum", "\n<replaceMe>enum")
	fileSlice := strings.Split(file, "<replaceMe>")
	fileOut := fileSlice[0]
	fileSlice = fileSlice[1:]
	sort.Slice(fileSlice, func(i int, j int) bool { return fileSlice[i] < fileSlice[j] })
	for _, v := range fileSlice {
		fileOut += v
	}
	err = ioutil.WriteFile("out.proto", []byte(fileOut), 0644)
	check(err)
	fmt.Print("complete")
}
