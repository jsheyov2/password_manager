package main

import "fmt"

const APP_NAME = "password_manager.exe"
const VESION = "0.22 alpha"

var RESERVED_NAMES = []string{"main.go", "Comments", "installer.go", "settings_manager.go", "types.go", "utils.go", "main.exe", "go.mod", APP_NAME, "assets", "encryptor.go", "go.sum", "graphics.go", "onefile.go"}

var COMPLETE_PROGRAMM = []string{"pmx.es"}

type trns struct {
	ENG string
	RUS string
}

// Strings

const Greeting_error string = "It looks like you`ve run this program in non-empty directory!\nPlease, put an exe-file in empty folder and restart it from there\n\n\nПохоже, вы запускаете программу в не пустой папке!\nПожалуйста, поместите exe-файл в свободную директорию и перезапустите его оттуда!"
const Not_installed string = "Program is uninstalled or files are broken. Press any key to repair.\n\n\nПрограма не установленна или её файлы повреждены. Нажмите, чтобы запустить установку/починку"
const Install_completed string = "Install is complited. Now restart this program\n\n\nПрограмма установлена. Теперь перезапустите её"
const Choose_language string = "Please, choose a language/Пожалуйста, выберете язык:\n0. English\n1. Русский"

//const MakePassword string =
//const Password_Too_Weak string =

// English / Russian
var TRANSLATIONS = map[int]trns{
	1:   {"Enter an username", "Введите имя для работы с программой"},
	2:   {"Please, enter a valid name. It may be unreal, but you should recognize it", "Пожалуйста, введите правильное имя. Оно может быть ненастоящим, но вам нужно его запомнить"},
	3:   {"Please, choose an encryption method for this profile:\n0. Weak (no encryption)\n1.Medium encryption(recomended)\n2. Strong encryption (it`s almost unable to get passwords unless you know a code)", "Пожалуйста, выберете метод шифрования для этого профиля:\n0. Слабое (без шифрования)\n1. Среднее(рекомендованно)\n2. Сильное (без ключа восстановить данные практически невозможно)"},
	4:   {"Setup is done. Now restart the app", "Настройка завершена. Перезапустите приложение"},
	5:   {"Password manager", "Менеджер паролей"},
	6:   {"Choose an apropriate option", "Выберете нужное действие"},
	7:   {"0. Get a password\n1. Write a password\n2. Additional options", "0. Получить пароль\n1. Записать пароль\n2. Дополнительные опции"},
	8:   {"Please, make a master-password to this user. Write it down to some paper, otherwise if you don`t remember it,\nALL YOUR PASSWORDS WILL BE LOST", "Пожалуйста, придумайте надежный мастер-пароль. Запишите его на бумагу, иначе можете\nПОТЕРЯТЬ ДОСТУП КО ВСЕМ СВОИМ ПАРОЛЯМ"},
	9:   {"Repeat your password", "Повторите пароль"},
	10:  {},
	100: {"Please, choose possible option", "Пожалуйста, выберете допустимое значение"},
	101: {"Please, enter a valid numer!", "Пожалуйста, введите допустимое число"},
	108: {"Your password is too week. Requirements:\n@ At least 10 characters long\n@ Contains at least 1 number and 1 upper letter\n@ Contains at least one of these simbols (*/!@#%^&()=+_)", "Ваш пароль сликом слабый. Требования:\n@ Длинна не менее 10 символов\n@ Содержит хотя бы 1 цифру и 1 заглавную букву\n@ Содержит хотя бы 1 спец. символ  (*/!@#%^&()=+_)"},
	109: {"Passwords do not match. Repeat again", "Пароли не совподают. Повторите попытку"},
}

func str(i any) string {
	return fmt.Sprintf("%v", i)
}
