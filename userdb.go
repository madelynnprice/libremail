package main

import (
	"encoding/json"
	"errors"
	"os"
)

type User struct {
	IDHash    string `json:"idHash"`
	SecretKey string `json:"secretKey"`
}

var userPath string = "/home/maddie/Code/libremail/users.json"

type UnregisteredUser struct {
	Username string `json:"username"`
}

func registerUser(user UnregisteredUser) (string, error) {
	if taken, err := isUsernameTaken(user.Username); taken == true || err != nil {
		if err != nil {
			return "", err
		}

		return "", errors.New("username is taken")
	}

	key, err := generatePassword(16)
	if err != nil {
		return "", err
	}

	newUser := User{IDHash: usernameID(user.Username), SecretKey: key}

	list, err := readUsers()

	if err != nil {
		return "", err
	}

	list = append(list, newUser)

	err = writeUsers(list)

	if err != nil {
		return "", err
	}

	return key, nil
}

func verifySecretKey(name string, key string) (bool, error) {
	if exists, err := isUsernameTaken(name); exists == false || err != nil {
		if err != nil {
			return false, err
		}

		return false, errors.New("User does not exist")
	}

	// we do not need to validate error as doesUserExist runs the same calls and has just been validated (maybe just best to validate this instead but oh well)
	user, _ := getUser(name)

	return key == user.SecretKey, nil
}

func readUsers() ([]User, error) {
	data, err := os.ReadFile(userPath)

	if err != nil {
		return []User{}, err
	}

	parse := []User{}

	err = json.Unmarshal(data, &parse)

	if err != nil {
		return []User{}, err
	}

	return parse, nil
}

func writeUsers(users []User) error {
	data, err := json.Marshal(users)

	if err != nil {
		return err
	}

	err = os.WriteFile(userPath, data, 0755)

	if err != nil {
		return err
	}

	return nil
}
