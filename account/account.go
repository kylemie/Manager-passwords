package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letter = []rune("qazwsxedcrfvtgbyhnujmikolpQAZWSXEDCRFVTGBYHNUJMIKOLP1234567890-*!")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"UpdateAt"`
}

func (acc *Account) generatePassword(n int) {
	result := make([]rune, n)
	for i := range result {
		result[i] = letter[rand.IntN(len(letter))]
	}
	acc.Password = string(result)
}

func (acc *Account) OutputAccount() {
	color.Cyan(acc.Login)
	color.Cyan(acc.Password)
	color.Cyan(acc.Url)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("Неверный логин")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Неверный URL")
	}
	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
		Url:       urlString,
		Login:     login,
		Password:  password,
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
