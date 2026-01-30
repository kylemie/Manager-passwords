package main

import (
	"fmt"
	"pet/project-1/account"
	"pet/project-1/files"
	"pet/project-1/output"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("___Менеджер паролей___")
	vault := account.NewVault(files.NewJsonDb("data.json"))
Menu:
	for {
		res := scandan([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		})
		switch res {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func findAccount(vault *account.VaultWithDb) {
	url := scandan([]string{"Введите URL: "})
	accounts := vault.FindAccountsByURL(url)
	if len(accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.OutputAccount()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := scandan([]string{"Введите URL: "})
	isDeleted := vault.DeleteAccountByURL(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := scandan([]string{"Введите логин: "})
	password := scandan([]string{"Введите пароль: "})
	url := scandan([]string{"Введите URL: "})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL или Login")
		return
	}
	vault.AddAccount(*myAccount)
}

func scandan[T any](promt []T) string {
	for index, value := range promt {
		if index == len(promt)-1 {
			fmt.Printf("%v: ", value)
		} else {
			fmt.Println(value)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
