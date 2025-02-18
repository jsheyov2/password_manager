package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func Test_enviroment() {
	dir, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	var valid bool
	for _, d := range dir {
		if !slices.Contains(RESERVED_NAMES, d.Name()) {
			valid = false
			break
		}
		valid = true
	}
	if !valid {
		Quit(Greeting_error)
	}
	c := os.Chdir("assets")
	if os.IsNotExist(c) {
		fmt.Println(Not_installed)
		fmt.Scanln()
		install()
		ClearConsone()
		Quit(Install_completed)

	}
	os.Chdir("assets")
	f, err := os.Stat("settings.json")
	if err != nil {
		Quit(err.Error())
	}
	if f.Size() == 0 {
		first_settings()
	}

}

func install() {
	os.Mkdir("assets", 0775)
	os.Chdir("assets")
	os.Create("settings.json")
	os.Mkdir("str", 0775)
	os.Chdir("str")
	for _, v := range COMPLETE_PROGRAMM {
		os.Create(v)
	}
}

func first_settings() {
	// Choose language
	func() {

		fmt.Println(Choose_language)
		l := NumChoose(1)
		Change_language(l)
	}()
	//Name
	func() {
		PrintLang(1)
		for {
			Change_profile(Input(""))
			if PROFILE == "" {
				PrintLang(2)
			} else {
				break
			}
		}
	}()
	//password
	password_maker()
	//Method
	func() {
		PrintLang(3)
		WriteMetod(encryption_metod(NumChoose(2)))
	}()
	Quit(Check_language(4))

}

func password_maker() {
	PrintLang(8)
	for {
		pass := Input("")
		if len(strings.Split(pass, "")) > 10 && strings.ContainsAny(pass, "*/!@#%^&()=+_") {
			func() {
				for {
					ClearConsone()
					PrintLang(9)
				pass_in18418:
					in := Input("")
					if pass == in {
						break
					}
					PrintLang(109)
					goto pass_in18418
				}
			}()
			WriteMasterPassword(Sha512(pass))
			return
		}
		PrintLang(108)
	}
}
