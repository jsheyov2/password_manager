package main

import (
	"encoding/json"
	"log"
	"os"
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
	os.Chdir("assets")
	os.Chdir("str")
	b, err := json.Marshal(j)
	if err != nil {
		log.Fatal(err)
	}
	os.Remove("prm.es")
	os.WriteFile("prm.es", b, 0775)
}

func WritePassword(u profile, metod encryption_metod, login, password, site_name string) {
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
