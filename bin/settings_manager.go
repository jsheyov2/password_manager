package main

import (
	"encoding/json"
	"os"
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
	b, _ := os.ReadFile("settings.json")
	var s Settings
	json.Unmarshal(b, &s)
	return &s
}

func write_settings(s *Settings) {
	ReturnToAssets()
	b, _ := json.Marshal(s)
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
	SETS = *read_settngs()
	LANG = SETS.Language
	ENCR = SETS.Encryption_metod
	PROFILE = SETS.Profile
	LAUNCH, _ = os.Executable()
	MASTER = SETS.Master
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
