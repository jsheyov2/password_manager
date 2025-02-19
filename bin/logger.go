package main

import (
	"fmt"
	"os"
)

const LOG bool = true // false / true
var START_TIME string

func Log(h, s string) {
	if LOG {
		to_log := fmt.Sprintf("[%v] (%v) %v\n", START_TIME, h, s)
		write_log([]byte(to_log))
	}
}

func LogErr(e error) {
	Log("ERROR", e.Error())
	os.Exit(1)
}

func write_log(d []byte) {
	ReturnToBase()
	os.WriteFile(START_TIME+".log", d, 0775)
}
