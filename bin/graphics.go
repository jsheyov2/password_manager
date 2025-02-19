package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Graphic_start() {
main_graphic_start:
	PrintLang(5)
	action()
	PrintLang(7)
	n := NumChoose(2)
	var p string
	switch n {
	case 0:
		{
		}
	case 1:
		{
			ClearConsone()
			PrintLang(10)
			s := NotEmptyString()
			ClearConsone()
			l := Input(Check_language(13))
			PrintLang(6)
			PrintLang(11)
			g := NumChoose(2)
			switch g {
			case 0:
				p = Gen64()
				fmt.Printf("%v%v\n", Check_language(17), p)
			case 1:
				p = Gen32()
			case 2:
				{
					p = Input(Check_language(12))
					fmt.Printf("%v%v\n", Check_language(17), p)
				}
			default:
				Log("ERROR", "Graphic start: default passmode")
			}
			WritePassword(PROFILE, ENCR, l, p, s)
			PrintLang(14)
			Input("")
			ClearConsone()
			goto main_graphic_start
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

func NotEmptyString() string {
	var x string
	for {
		x = Input("")
		if x != "" && x != " " {
			return x
		}
		PrintLang(100)
	}
}
