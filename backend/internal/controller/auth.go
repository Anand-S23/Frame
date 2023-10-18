package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Anand-S23/frame/internal/models"
	"github.com/Anand-S23/frame/internal/validators"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) SignUp(w http.ResponseWriter, r *http.Request) error {
    var userData models.UserDto
    err := json.NewDecoder(r.Body).Decode(&userData)
    if err != nil {
        errMsg := map[string]string {
            "error": "Could not parse sign up data",
        }
        return WriteJSON(w, http.StatusBadRequest, errMsg)
    }

    errs := validators.AuthValidator(userData, c.store)
    if len(errs) != 0 {
        errMsgs := make([]string, 0, 3)
        for _, err := range errs {
            errMsgs = append(errMsgs, err.Error())
        }

        return WriteJSON(w, http.StatusBadRequest, map[string][]string {
            "errors": errMsgs,
        })
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
        errMsg := map[string]string {
            "error": "Internal server error, please try again",
        }
        return WriteJSON(w, http.StatusBadRequest, errMsg)
	}

    // TODO: Store the user to database

    return WriteJSON(w, http.StatusOK, "Signed up")
}
