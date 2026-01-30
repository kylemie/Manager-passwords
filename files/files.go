package files

import (
	"os"
	"pet/project-1/output"

	"github.com/fatih/color"
)

type JsonGb struct {
	filename string
}

func NewJsonDb(name string) *JsonGb {
	return &JsonGb{
		filename: name,
	}
}

func (db *JsonGb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		output.PrintError(err)
		return nil, err
	}
	return data, nil
}

func (db *JsonGb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintError(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		output.PrintError("Неверный формат URL или Login")
		return
	}
	color.Green("Запись успешна")
}
