package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	v2 "math/rand/v2"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
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
func GetLaunch() {
	l, err := os.Executable()
	if err != nil {
		Quit("ERROR in launch")
	}
	l = filepath.ToSlash(l)
	LAUNCH = strings.TrimSuffix(l, APP_NAME)
	//Printl(APP_NAME)
	//Quit(LAUNCH)
}

func ReturnToBase() {
	err := os.Chdir(LAUNCH)

	if err != nil {
		print(err.Error())
		LogErr(err)
	}
}

func ReturnToAssets() {
	ReturnToBase()
	err := os.Chdir("assets")
	if err != nil {
		LogErr(err)
	}
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

// Works with any table
func passgen(length int, table *map[string]int) string {
	var output string
	for i := 0; i <= length; i++ {
		r := v2.IntN(len(*table))
		for k, v := range *table {
			if v == r {
				output += k
			}
		}
	}
	return output
}

// Works with default table
func Gen32() string {
	return passgen(32, &DEFAULT_TABLE)
}

// Works with default table
func Gen64() string {
	return passgen(64, &DEFAULT_TABLE)
}

// To any type in future
func To_one_lenght(a, b *[]int) {
	if len(*a) > len(*b) {
		for len(*a) > len(*b) {
			*b = append(*b, *b...)
		}
		*b = (*b)[:len(*a)]
	}
	if len(*b) > len(*a) {
		for len(*b) > len(*a) {
			*a = append(*a, *a...)
		}
		*a = (*a)[:len(*b)]
	}
}

func Xor(a, b []int) []int {
	c := []int{}
	for i, v := range a {
		c = append(c, v^b[i])
	}
	return c
}

func Curcular_shift(n []int, positions int, left bool) []int {
	if !left {
		slices.Reverse(n)
	}
	for i := 0; i < int(positions); i++ {
		corner := n[0]
		for num := range n {
			if num+1 >= len(n) {
				n[num] = corner
				break
			}
			n[num] = n[num+1]
		}
	}
	if !left {
		slices.Reverse(n)
	}
	return n
}
