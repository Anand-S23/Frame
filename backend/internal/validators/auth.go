package validators

import (
	"github.com/Anand-S23/frame/internal/models"
	"github.com/Anand-S23/frame/internal/storage"
)

func AuthValidator(userData models.UserDto, store *storage.MongoStore) []error {
    errs := make([]error, 0, 3)

    err := validateUsername(userData.Username, store)
    if err != nil {
        errs = append(errs, err)
    }

    err = validateEmail(userData.Email, store)
    if err != nil {
        errs = append(errs, err)
    }

    err = validatePassword(userData.Password)
    if err != nil {
        errs = append(errs, err)
    }

    return errs
}

func validateUsername(username string, store *storage.MongoStore) error {
    return nil
}

func validateEmail(email string, store *storage.MongoStore) error {
    return nil
}

func validatePassword(password string) error {
    return nil
}
