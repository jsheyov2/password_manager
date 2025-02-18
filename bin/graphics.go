package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Graphic_start() {
	PrintLang(5)
	action()
	PrintLang(7)
	n := NumChoose(2)
	switch n {
	case 0:
		{

		}
	case 1:
		{
		}
	case 2:
		{
		}
	}
}

func action() {
	PrintLang(6)
}

func Printl(s string) {
	fmt.Println(s)
}

func PrintLang(i int) {
	fmt.Println(Check_language(i))
}

func ClearConsone() {
	cmd := exec.Command("clear")
	if _, err := exec.LookPath("cmd.exe"); err == nil {
		cmd = exec.Command("cmd.exe", "/c", "cls")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
