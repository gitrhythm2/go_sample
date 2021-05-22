package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

type Bookmark struct {
	Key         string
	Uri         string
	Description string
}

type ArrayConfig struct {
	Bookmark []Bookmark
}

func tomlArray() {
	fmt.Println("tomlArray in")
	saveConfig()
	loadConfig()
}

func saveConfig() {
	file, err := os.Create("config_array.toml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	config := ArrayConfig{}
	config.Bookmark = []Bookmark{
		{Key: "goo", Uri: "https://www.google.co.jp", Description: "Google site"},
		{Key: "am", Uri: "https://www.google.co.jp", Description: "Amazon site"},
		{Key: "gh", Uri: "https://www.github.com", Description: "GitHub site"},
	}

	encoder := toml.NewEncoder(file)
	encoder.Encode(config)
}

func loadConfig() {
	file, err := os.Open("config_array.toml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	config := ArrayConfig{}
	toml.Unmarshal(data, &config)
	fmt.Printf("%#v\n", config)
}
