package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func Quit(msg string) {
	if msg != "" {
		fmt.Print(msg)
	}
	fmt.Scanln()
	os.Exit(0)
}

func Input(s string) string {
	var i string
	if s != "" {
		fmt.Println(s)
	}
	fmt.Scanln(&i)
	return i
}

func Check_language(i int) string {
	for k, v := range TRANSLATIONS {
		if i == k {
			if LANG == 0 {
				return v.ENG
			} else if LANG == 1 {
				return v.RUS
			}
		}
	}
	return "ERROR CODE 107"
}

func Sha512(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	hib := h.Sum(nil)
	return hex.EncodeToString(hib)
}

func ReturnToBase() {
	os.Chdir(LAUNCH)
}

func ReturnToAssets() {
	ReturnToBase()
	os.Chdir("assets")
}

func NumChoose(max int) int {
	array := []string{}
	for i := 0; i <= max; i++ {
		array = append(array, str(i))
	}
	var input string
	for {
		input = Input("")
		if slices.Contains(array, input) {
			x, _ := strconv.Atoi(input)
			return x
		}
		PrintLang(100)
	}
}
