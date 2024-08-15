package main

import (
	"bufio"
	"log"
	"os"
)

// rename file == file list with new names

/*
if rename file does not exists
	if argument present
		create filtered rename file
	else
		create general rename file
else
	move files
	delete rename file
*/

func main() {
	info, err := os.Stat("names")
	if err == nil { // path exists
		if info.IsDir() {
			log.Fatalf("ERROR: 'names' is a directory")
		}
		log.Print("'names' file exists, moving files now")
		linesOld := linesFromFile(".names.old")
		linesNew := linesFromFile("names")
		for k := range linesOld {
			if linesOld[k] != "" && linesNew[k] != "" {
				os.Rename(linesOld[k], linesNew[k])
			}
		}
		os.Remove("names")
		os.Remove(".names.old")
	} else {
		items, err := os.ReadDir(".")
		if err != nil {
			log.Fatalf("Error when reading directory contents: %v", err)
		}
		fOld, err := os.OpenFile(".names.old", os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatalf("Error when opening '.names.old' file for writing: %v", err)
		}
		defer fOld.Close()
		fNew, err := os.OpenFile("names", os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatalf("Error when opening 'names' file for writing: %v", err)
		}
		defer fNew.Close()
		for _, item := range items {
			if !item.IsDir() {
				if _, err = fOld.WriteString(item.Name() + "\n"); err != nil {
					log.Fatalf("Error when writing line to file: %v", err)
				}
				if _, err = fNew.WriteString(item.Name() + "\n"); err != nil {
					log.Fatalf("Error when writing line to file: %v", err)
				}
			}
		}
		fOld.Chmod(0400)
	}
}

func linesFromFile(file string) []string {
	input, err := os.Open(file)
	if err != nil {
		log.Fatalf("Error when opening '%s' file for reading: %v", file, err)
	}
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanLines)

	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
