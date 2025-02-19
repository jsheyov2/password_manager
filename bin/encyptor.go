package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"
)

type Json_passwords struct {
	User profile `json : "User"`
	Data []Site  `json : "Data"`
}

type Site struct {
	name      string     `json : "name"`
	Site_data []Password `json : "site data"`
}

type Password struct {
	Method   encryption_metod `json : method`
	Login    string           `json : "login"`
	password string           `json : "password"`
	datatime time.Time        `json : datatime`
}

func read_passwords_from_json() *Json_passwords {
	os.Chdir(LAUNCH)
	os.Chdir("assets")
	os.Chdir("str")
	b, err := os.ReadFile("prm.es")
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	var jspw Json_passwords
	json.Unmarshal(b, &jspw)
	os.Chdir(LAUNCH)
	return &jspw
}

func write_to_json(j *Json_passwords) {
	ReturnToBase()
	err := os.Chdir("assets")
	if err != nil {
		Log("ERROR", err.Error())
		os.Exit(1)
	}
	err = os.Chdir("str")
	if err != nil {
		LogErr(err)
	}
	b, err := json.Marshal(j)
	if err != nil {
		LogErr(err)
	}
	os.Remove("prm.es")
	os.WriteFile("prm.es", b, 0775)
}

func WritePassword(u profile, metod encryption_metod, login, password, site_name string) {
	switch metod {
	case encryption_metod(13):

	case encryption_metod(14):
		password = medium_encryption(password)
	case encryption_metod(15):
		//Strong `ll be added soon
	}
	p := Password{
		metod, login, password, time.Now(),
	}
	ps := []Password{p}
	s := Site{
		site_name, ps,
	}
	c := read_passwords_from_json()
	c.Data = append(c.Data, s)
	write_to_json(c)

}

func medium_encryption(p string) string {
	/*
		Medium encryption:
		password, master sha-512 -> []int, to one length
		password xor master sha-512
		password << (circular) on nums of consonants in original password
		password to text(using table)
	*/
	ipwd, c := string_to_int(p, &DEFAULT_TABLE)
	imaster, _ := string_to_int(MASTER, &DEFAULT_TABLE)
	To_one_lenght(&ipwd, &imaster)
	r := Xor(ipwd, ipwd)
	res := Curcular_shift(r, c, true)
	result, _ := int_to_string(res)
	return result
}
func strong_encryption(p string) string {
	return ""
}

func string_to_int(s string, _ *map[string]int) ([]int, int) {
	var output []int
	for _, v := range s {
		output = append(output, DEFAULT_TABLE[string(v)])
		r := recover()
		if r != nil {
			return nil, -1
		}
	}
	return output, count_cons(s)
}
func int_to_string(a []int) (string, int) {
	out := ""
	for _, v := range a {
		for k, v1 := range DEFAULT_TABLE {
			if v == v1 {
				out += k
			}
		}
	}
	return out, count_cons(out)
}

func count_cons(s string) int {
	var i = 0
	for _, v := range s {
		if strings.ContainsAny(string(v), "AaEeYyUuIiOo") {
			i++
		}
	}
	return i
}
