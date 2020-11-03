package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var root = "ead-files"

func main() {
	//delete the root directory if it exists
	if _, err := os.Stat("root"); !os.IsNotExist(err) {
		e := os.Remove(root)
		if e != nil {
			panic(e)
		}
	}

	//create the directories
	err := makeDirectories()
	if err != nil {
		panic(err)
	}

	//open and iterate through the tsv file
	tsv, err := os.Open("sample-set.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(tsv)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		coll := fmt.Sprintf("/repositories/%s%s", line[1], line[0])
		fmt.Println(coll)
	}
}

func makeDirectories() error {

	err := os.Mkdir(root, 0777)
	if err != nil {
		return err
	}
	repos := []string{"archives", "fales", "tamwag"}

	for _, repo := range repos {
		fp := filepath.Join(root, repo)
		os.Mkdir(fp, 0777)
	}

	return nil
}
