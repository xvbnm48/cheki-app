package entity

var jwtKey = []byte("your_secret")

type LoginResponse struct {
	Token string `json:"token"`
}
