package main

import (
	"bufio"
	"fmt"
	"github.com/nyudlts/go-aspace"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var root = "ead-files"
var client *aspace.ASClient
var repositories = map[string]string{
	"2": "tamwag",
	"3": "fales",
	"6": "archives",
}

func main() {
	//delete the root directory if it exists
	if _, err := os.Stat(root); !os.IsNotExist(err) {
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

	//create an aspace client
	client, err := aspace.NewClient("fade", 50)
	if err != nil {
		panic(err)
	}
	client.GetAspaceInfo()

	//open and iterate through the tsv file
	tsv, err := os.Open("sample-set.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(tsv)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		repositoryId, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}
		resourceId, err := strconv.Atoi(strings.Split(line[0], "/")[2])
		if err != nil {
			panic(err)
		}
		fmt.Println(repositoryId, resourceId)

		//generate the eadfiles
		loc := filepath.Join(root, repositories[line[1]])
		err = client.SerializeEAD(repositoryId, resourceId, loc,true, false, false, false, false, line[2])
		if err != nil {
			panic(err)
		}
	}
}

func writeEAD() error {
	return nil
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
