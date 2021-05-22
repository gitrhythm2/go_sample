package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

// https://github.com/pelletier/go-toml
func tomlBasic() {
	//basic()
	//marshaling()
	writeToFile()
	readFromFile()
}

func basic() {
	fmt.Println("toml in")
	config, _ := toml.Load(string(configData()))
	user := config.Get("postgres.user").(string)
	fmt.Println(user)

	postgresConfig := config.Get("postgres").(*toml.Tree)
	password := postgresConfig.Get("password").(string)
	fmt.Println(password)
	fmt.Printf("%#v\n", postgresConfig)
}

type Db struct {
	User     string
	Password string
}

type Config struct {
	Postgres Db
}

func marshaling() {
	config := Config{}
	toml.Unmarshal(configData(), &config)
	fmt.Println("user:", config.Postgres.User)
	fmt.Printf("%#v\n", config)
}

func writeToFile() {
	file, err := os.Create("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	config := Config{}
	toml.Unmarshal(configData(), &config)
	encoder := toml.NewEncoder(file)
	encoder.Encode(config)
}

func readFromFile() {
	file, err := os.Open("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	config := Config{}
	toml.Unmarshal(data, &config)
	fmt.Printf("%#v\n", config)
}

func configData() []byte {
	return []byte(`[postgres]
	user = "pelletier"
	password = "mypassword"
	`)
}
