package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalln(err)
	}

	textFiles := make([]string, 0, 5)
	for _, fileInfo := range files {
		fileName := fileInfo.Name()
		if strings.HasSuffix(fileName, ".txt") {
			textFiles = append(textFiles, fileName)
		}
	}

	for _, textFileName := range textFiles {
		err := searchAndWriteFiles(textFileName)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func searchAndWriteFiles(fileName string) error {
	inputFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	outputFile, err := os.Create(fileName + "_" + time.Now().Format("20060102-0304") + ".csv")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	bytes, err := ioutil.ReadAll(inputFile)
	if err != nil {
		return err
	}

	targets := searchAddress(string(bytes))
	log.Println(targets)

	fmt.Fprintln(outputFile, time.Now().Format("20060102"))
	for _, target := range targets {
		fmt.Fprintln(outputFile, target)
	}

	return nil
}

func searchAddress(s string) []string {
	targets := make([]string, 0, 40)
	for _, word := range strings.Split(strings.Replace(s, "\n", " ", -1), " ") {
		if strings.Contains(word, "市") || strings.Contains(word, "町") || strings.Contains(word, "村") || strings.Contains(word, "郡") {
			targets = append(targets, word)
		}
	}
	return targets
}
