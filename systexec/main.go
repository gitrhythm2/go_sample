package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

// syscall.Execのサンプル
// 参考: Go by Example: Exec'ing Processes
// https://oohira.github.io/gobyexample-jp/execing-processes.html
func main() {
	execLs()
	//execDockerLogin()
}

func execLs() {
	path, err := exec.LookPath("ls")
	fmt.Printf("path: %q\n", path)
	if err != nil {
		log.Fatal(err)
	}

	args := []string{"ls", "-alF"}
	env := os.Environ()
	fmt.Printf("env: %q\n", env)

	err = syscall.Exec(path, args, env)
	if err != nil {
		log.Fatal(err)
	}
}

func execDockerLogin() {
	path, err := exec.LookPath("docker")
	if err != nil {
		log.Fatal(err)
	}

	args := []string{"docker", "exec", "-it", "dwrapper", "/usr/bin/bash", "--login"}
	env := os.Environ()
	err = syscall.Exec(path, args, env)
	if err != nil {
		log.Fatal(err)
	}
}
