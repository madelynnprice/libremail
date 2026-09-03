package main

import (
	"encoding/json"
	"errors"
	"os"
)

type Mail struct {
	Title    string `json:"title"`
	Contents string `json:"contents"`
}

var mailPath = "/home/maddie/Code/libremail/mailbox.json"

func registerMailForUser(mail Mail, username string) error {
	if taken, err := isUsernameTaken(username); taken == false || err != nil {
		if err != nil {
			return err
		}

		return errors.New("user does not exist")
	}

	list, err := readMail()

	if err != nil {
		return err
	}

	list[username] = append(list[username], mail)

	writeMail(list)
	return nil
}

func readMailForUser(username string) ([]Mail, error) {
	list, err := readMail()

	if err != nil {
		return []Mail{}, err
	}

	return list[username], nil
}

func readMail() (map[string][]Mail, error) {
	data, err := os.ReadFile(mailPath)

	parse := map[string][]Mail{}

	if err != nil {
		return parse, err
	}

	err = json.Unmarshal(data, &parse)

	if err != nil {
		return parse, err
	}

	return parse, nil
}

func writeMail(mail map[string][]Mail) error {
	data, err := json.Marshal(mail)

	if err != nil {
		return err
	}

	err = os.WriteFile(mailPath, data, 0755)

	if err != nil {
		return err
	}

	return nil
}
