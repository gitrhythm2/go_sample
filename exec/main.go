package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

// os/execのサンプル
// 参考: Go by Example: Spawning Processes
// https://oohira.github.io/gobyexample-jp/spawning-processes
//
// Go の exec.Commandを調査する
// https://qiita.com/TsuyoshiUshio@github/items/22cafc8a4dc097add73b
func main() {
	fmt.Println("exec sample in")
	//execOutput()
	//execRun()
	//execStart()
	//execStartWait()
	//execPipe()
}

func execOutput() {
	fmt.Println("execOutput():")
	cmd := exec.Command("ls", "-laF")
	fmt.Printf("%#v\n\n", cmd)

	// Output()は処理結果を返す
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

// Run()は処理結果を返さない。且つ処理をwaitする
func execRun() {
	fmt.Println("execRun():")
	fmt.Println("Run(sleep 3) before")
	cmd := exec.Command("sleep", "3")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("sleep after")
}

// Startはwaitしない(値も返さない)
func execStart() {
	fmt.Println("execStart():")
	fmt.Println("Run(sleep 3) before")
	cmd := exec.Command("sleep", "3")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("sleep after")
}

// Startで処理結果を待つにはWait()する
func execStartWait() {
	fmt.Println("execStart():")
	fmt.Println("Run(sleep 3) before")
	cmd := exec.Command("sleep", "3")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	cmd.Wait()
	fmt.Println("sleep after")
}

// Pipeを使って標準出力をコマンドの標準入力に渡す方法
func execPipe() {
	cmd := exec.Command("wc", "-w")
	stdin, _ := cmd.StdinPipe()
	io.WriteString(stdin, "hoge fuga chome")
	stdin.Close()
	out, _ := cmd.Output()
	fmt.Printf("result: %s", out)
}
