package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	English language = iota
	Russian
	Week encryption_metod = iota + 11
	Medium
	Strong
)

type Settings struct {
	Profile          profile          `json : "profile"`
	Master           string           `json: "mp"`
	Language         language         `json : "language"`
	Encryption_metod encryption_metod `json : "method"`
}

type language int
type encryption_metod int
type profile string

var LANG language
var ENCR encryption_metod
var SETS Settings
var PROFILE profile
var LAUNCH string
var MASTER string

func Change_language(i int) {
	s := read_settngs()
	s.Language = language(i)
	write_settings(s)
	LANG = language(i)
}

func read_settngs() *Settings {
	ReturnToAssets()
	b, err := os.ReadFile("settings.json")
	if err != nil {
		LogErr(err)
	}
	var s Settings
	json.Unmarshal(b, &s)
	return &s
}

func write_settings(s *Settings) {
	ReturnToAssets()
	b, err := json.Marshal(s)
	if err != nil {
		LogErr(err)
	}
	os.Remove("settings.json")
	os.WriteFile("settings.json", b, 0775)
}

func Change_profile(p string) {
	s := read_settngs()
	s.Profile = profile(p)
	PROFILE = profile(p)
	write_settings(s)

}
func Change_mode(i int) {
	s := read_settngs()
	s.Encryption_metod = encryption_metod(i)
	ENCR = s.Encryption_metod
	write_settings(s)
}
func Startup() {
	GetLaunch()
	START_TIME = fmt.Sprintf("%v", time.Now().Format("15:04:05 02.01.2006"))
	SETS = *read_settngs()
	LANG = SETS.Language
	ENCR = SETS.Encryption_metod
	PROFILE = SETS.Profile
	MASTER = SETS.Master
	check_password_on_login()
}

func WriteMasterPassword(m string) {
	ReturnToAssets()
	s := read_settngs()
	s.Master = m
	write_settings(s)
	MASTER = m
}

func WriteMetod(m encryption_metod) {
	ReturnToAssets()
	s := read_settngs()
	s.Encryption_metod = m
	write_settings(s)
}

func check_password_on_login() {
	p := Input(Check_language(15))
	if Sha512(p) == MASTER {
		return
	}
	PrintLang(16)
	Input("")
	os.Exit(0)
}
