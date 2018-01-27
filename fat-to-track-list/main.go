package main

/*
My first line filter Go program to enumerate tracks
on a FAT formatted pendrive.
The input if from the command "sudo fatsort -fl /dev/sdc1",
output is stdout
*/

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

import "github.com/mitchellh/go-wordwrap"

const width uint = 34

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fileIndex := 0
	folderIndex := 0
	isStarted := false

	for scanner.Scan() {
		line := scanner.Text()
		folderMatch, _ := regexp.MatchString("^/\\w+.*", line)
		if folderMatch {
			isStarted = true
			folderIndex++
			fmt.Printf("\n\n===========  Fol # %d  =============\n", folderIndex)
			fmt.Printf("  %v\n", indent(fixAlbumName(line), 2))
		}
		fileMatch, _ := regexp.MatchString("^[^/]+", line)
		if fileMatch {
			if isStarted {
				fileIndex++
				fmt.Printf("%4d: ", fileIndex)
				fmt.Println(indent(fixFileName(line), 6))
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "error: ", err)
		os.Exit(1)
	}
}

func fixAlbumName(in string) string {
	res := strings.Replace(strings.Title(in), "/", "", -1)
	res = strings.Replace(res, "-", " ", -1)
	wrapped := wordwrap.WrapString(res, width)
	return wrapped
}

func fixFileName(in string) string {
	res := strings.TrimLeft(strings.TrimLeft(in, "0123456789"), " .-")
	res = strings.TrimSuffix(res, ".mp3")
	res = strings.Replace(res, "_", "", -1)
	res = strings.Title(res)
	wrapped := wordwrap.WrapString(res, width)
	return wrapped
}

func indent(in string, num int) string {
	return strings.Replace(in, "\n", "\n"+strings.Repeat(" ", num), -1)
}
