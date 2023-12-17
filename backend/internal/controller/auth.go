package controller

import (
	"encoding/json"
	"log"
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

func createJWTCookie(token string, expDuration time.Duration, secure bool) http.Cookie {
    return http.Cookie {
        Name: "jwt",
        Value: token,
        Expires: time.Now().Add(expDuration),
        HttpOnly: true,
        Secure: secure,
        Path: "/",
        SameSite: http.SameSiteStrictMode,
    }
}

func createExpiredJWTCookie(secure bool) http.Cookie {
    return http.Cookie {
        Name: "jwt",
        Value: "",
        Expires: time.Unix(0, 0),
        HttpOnly: true,
        Secure: secure,
        Path: "/",
        SameSite: http.SameSiteStrictMode,
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

    authErrs := validators.AuthValidator(userData, c.store)
    if len(authErrs) != 0 {
        return WriteJSON(w, http.StatusBadRequest, authErrs)
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
        log.Println("Error hashing the password")
        return InternalServerError(w)
	}
    userData.Password = string(hashedPassword)

    user := models.NewUser(userData)
    insertResult, err := c.store.CreateUser(user)
    if err != nil {
        log.Println("Error storing the password in the database")
        return InternalServerError(w)
    }
    insertedID := insertResult.InsertedID.(primitive.ObjectID).Hex()

    expDuration := time.Hour * 24
    token, err := createToken(c.JwtSecretKey, insertedID, expDuration)
    if err != nil {
        log.Println("Error creating token")
        return InternalServerError(w)
    }

    cookie := createJWTCookie(token, expDuration, c.production)
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
        log.Println("Error creating token")
        return InternalServerError(w)
    }

    cookie := createJWTCookie(token, expDuration, c.production)
    http.SetCookie(w, &cookie)

    successMsg := map[string]string {
        "message": "User logged in successfully",
    }
    return WriteJSON(w, http.StatusOK, successMsg)
}

func (c *Controller) Logout(w http.ResponseWriter, r *http.Request) error {
    cookie := createExpiredJWTCookie(c.production)
    http.SetCookie(w, &cookie)
    return WriteJSON(w, http.StatusOK, "")
}

// TODO: Only for test purposes
func (c *Controller) GetAllUsers(w http.ResponseWriter, r *http.Request) error {
    users := c.store.GetAllUsers()
    return WriteJSON(w, http.StatusOK, users)
}

func (c *Controller) GetAuthenticatedUser(w http.ResponseWriter, r *http.Request) error {
    cookie, err := r.Cookie("jwt")
    if err != nil {
        return WriteJSON(w, http.StatusUnauthorized, "Authentication token not found")
    }

    tokenString := cookie.Value
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(c.JwtSecretKey), nil
    })
    if err != nil || !token.Valid {
        return WriteJSON(w, http.StatusUnauthorized, "Invalid authentication token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return WriteJSON(w, http.StatusUnauthorized, "Invalid token claims")
    }

    userID, ok := claims["user_id"].(string)
    if !ok {
        return WriteJSON(w, http.StatusUnauthorized, "User not found")
    }

    user := c.store.FindUserByID(userID)
    if user == nil {
        return WriteJSON(w, http.StatusUnauthorized, "User not found")
    }

    return WriteJSON(w, http.StatusOK, user)
}

