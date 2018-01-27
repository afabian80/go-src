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
			fmt.Printf("\n\nAlbum %d - %v\n\n", folderIndex, fixAlbumName(line))
		}
		fileMatch, _ := regexp.MatchString("^[^/]+", line)
		if fileMatch {
			if isStarted {
				fileIndex++
				fmt.Printf("\t%5d: %v\n", fileIndex, line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "error: ", err)
		os.Exit(1)
	}
}

func fixAlbumName(in string) string {
	return strings.Replace(strings.Title(in), "/", "", -1)
}
