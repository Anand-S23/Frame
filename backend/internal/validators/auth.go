package validators

import (
	"errors"
	"net/mail"
	"regexp"

	"github.com/Anand-S23/frame/internal/models"
	"github.com/Anand-S23/frame/internal/storage"
)

func AuthValidator(userData models.UserDto, store *storage.MongoStore) error {
    err := validateUsername(userData.Username, store)
    if err != nil {
        return err
    }

    err = validateEmail(userData.Email, store)
    if err != nil {
        return err
    }

    err = validatePassword(userData.Password)
    if err != nil {
        return err
    }

    return nil
}

func validateUsername(username string, store *storage.MongoStore) error {
    if len(username) < 5 || len(username) > 15 {
        return errors.New("Username must be between 5-15 characters long")
    }

    if !regexp.MustCompile("^[a-zA-Z0-9_\\-.]+$").MatchString(username) {
        return errors.New("Username can only contain letters, numbers, hyphens(-), underscores(_) and periods(.)")
    }

    user := store.FindUserByUsername(username)
    if user != nil {
        return errors.New("User already exsits with that username")
    }
    return nil
}

func validateEmail(email string, store *storage.MongoStore) error {
    _, err := mail.ParseAddress(email)
    if err != nil {
        return errors.New("Email entered is not valid")
    }

    user := store.FindUserByEmail(email)
    if user != nil {
        return errors.New("User already exsits with that email")
    }

    return nil
}

func validatePassword(password string) error {
    if len(password) < 8 || len(password) > 30 {
        return errors.New("Password must be between 8 and 30 characters long")
    }

    return nil
}

