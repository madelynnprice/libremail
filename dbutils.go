package main

import (
	"errors"
)

func getUser(name string) (User, error) {
	list, err := readUsers()

	if err != nil {
		return User{}, err
	}

	for _, u := range list {
		if u.IDHash == usernameID(name) {
			return u, nil
		}
	}

	return User{}, errors.New("user does not exist")
}

// must take UNHASHED name, not ID
func isUsernameTaken(name string) (bool, error) {
	_, err := getUser(name)

	if err != nil {
		if err.Error() != "user does not exist" {
			return true, err
		} else {
			return false, nil
		}
	}

	return true, nil
}
