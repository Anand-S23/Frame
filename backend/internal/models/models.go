package models

type UserDto struct {
    Username string
    Email    string
    Password string
}

type User struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"-"`
    Avatar   string `json:"avatar"`
}



