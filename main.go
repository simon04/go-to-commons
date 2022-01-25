package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"cgt.name/pkg/go-mwclient"
)

const useragent = "go-to-commons"

var api = flag.String("api", "https://commons.wikimedia.org/w/api.php", "MediaWiki API URL")
var username = flag.String("username", os.Getenv("MEDIAWIKI_USERNAME"), "Wikimedia username")
var password = flag.String("password", os.Getenv("MEDIAWIKI_PASSWORD"), "Wikimedia password")
var comment = flag.String("comment", "Uploaded with "+useragent, "Upload comment")
var file = flag.String("file", "", "Media file to upload")
var filename = flag.String("filename", "", "Filename on Wikimedia Commons")
var text = flag.String("text", "", "Wikitext of media file on Wikimedia Commons or (if specified as @file.txt, the text is read from file.txt)")

func main() {
	flag.Parse()
	if *file == "" {
		panic("-file needs to be specified!")
	} else if *text == "" {
		panic("-text needs to be specified!")
	}
	if *filename == "" {
		_, *filename = path.Split(*file)
	}

	fmt.Printf("Reading file %s\n", *file)
	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	w, err := mwclient.New(*api, useragent)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Logging in as %s...\n", *username)
	err = w.Login(*username, *password)
	var warnings mwclient.APIWarnings
	if err != nil && !errors.As(err, &warnings) {
		panic(err)
	}

	fmt.Printf("Retrieving CSRF token...\n")
	token, err := w.GetToken(mwclient.CSRFToken)
	if err != nil {
		panic(err)
	}

	if strings.HasPrefix(*text, "@") {
		base := (*text)[len("@"):]
		fmt.Printf("Reading description text from file %s...\n", base)
		bytes, err := ioutil.ReadFile(base)
		if err != nil {
			panic(err)
		}
		*text = string(bytes)
	} else if strings.HasPrefix(*text, "base64:") {
		base := (*text)[len("base64:"):]
		bytes, err := base64.StdEncoding.DecodeString(base)
		if err != nil {
			panic(err)
		}
		*text = string(bytes)
	}

	fmt.Printf("Uploading %s as %s to %s...\n", *file, *filename, *api)
	_, err = w.Post(map[string]string{
		"action":   "upload",
		"comment":  *comment,
		"file":     string(bytes),
		"filename": *filename,
		"format":   "json",
		"text":     *text,
		"token":    token,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Finished successfully :-)\n")
}
