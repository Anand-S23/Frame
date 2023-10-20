package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Anand-S23/frame/internal/models"
	"github.com/Anand-S23/frame/internal/validators"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func createToken(secretKey string, userID string, expDuration time.Duration) (string, error) {
    token := jwt.New(jwt.GetSigningMethod("HS256"))
    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = userID
    claims["exp"] = time.Now().Add(expDuration).Unix()

    return token.SignedString([]byte(secretKey))
}

func createJWTCookie(token string, expDuration time.Duration) http.Cookie {
    return http.Cookie {
        Name: "jwt",
        Value: token,
        Expires: time.Now().Add(expDuration),
        HttpOnly: true,
        Secure: false, // TODO: Need to dynamically decide this depending on development mode or production
        Path: "/",
    }
}

func (c *Controller) SignUp(w http.ResponseWriter, r *http.Request) error {
    var userData models.UserDto
    err := json.NewDecoder(r.Body).Decode(&userData)
    if err != nil {
        errMsg := map[string]string {
            "error": "Could not parse sign up data",
        }
        return WriteJSON(w, http.StatusBadRequest, errMsg)
    }

    err = validators.AuthValidator(userData, c.store)
    if err != nil {
        errMsg := map[string]string {
            "error": err.Error(),
        }
        return WriteJSON(w, http.StatusBadRequest, errMsg)
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
        return InternalServerError(w)
	}
    userData.Password = string(hashedPassword)

    user := models.NewUser(userData)
    insertResult, err := c.store.CreateUser(user)
    if err != nil {
        return InternalServerError(w)
    }
    insertedID := insertResult.InsertedID.(primitive.ObjectID).Hex()

    expDuration := time.Hour * 24
    token, err := createToken(c.JwtSecretKey, insertedID, expDuration)
    if err != nil {
        return InternalServerError(w)
    }

    cookie := createJWTCookie(token, expDuration)
    http.SetCookie(w, &cookie)

    successMsg := map[string]string {
        "message": "User created successfully",
        "userID": insertedID,
    }

    return WriteJSON(w, http.StatusOK, successMsg)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) error {
    var loginData models.LoginDto
    err := json.NewDecoder(r.Body).Decode(&loginData)
    if err != nil {
        errMsg := map[string]string {
            "error": "Could not parse login data",
        }
        return WriteJSON(w, http.StatusBadRequest, errMsg)
    }

    user := c.store.FindUserByEmail(loginData.Email)
    if user == nil {
        errMsg := map[string]string {
            "error": "Incorrect email or password, please try again",
        }
        return WriteJSON(w, http.StatusBadRequest, errMsg)
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
        errMsg := map[string]string {
            "error": "Incorrect email or password, please try again",
        }
        return WriteJSON(w, http.StatusBadRequest, errMsg)
	}

    expDuration := time.Hour * 24
    token, err := createToken(c.JwtSecretKey, user.User_ID, expDuration)
    if err != nil {
        return InternalServerError(w)
    }

    cookie := createJWTCookie(token, expDuration)
    http.SetCookie(w, &cookie)

    successMsg := map[string]string {
        "message": "User logged in successfully",
    }
    return WriteJSON(w, http.StatusOK, successMsg)
}

