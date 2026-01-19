package objects

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Contact string `json:"contact"`
	Name    string `json:"name"`
	jwt.RegisteredClaims
}

type AuthResponse struct {
	Status bool   `json:"status"`
	Token  string `json:"token"`
}

type UserProfileData struct {
	Name         string `json:"name"`
	UserName     string `json:"user_name"`
	ProfilePhoto string `json:"profile_photo"`
}

// Define a struct for the response since we are returning more than just strings now
type BlockedUserDetails struct {
	Name         string `json:"name"`
	UserName     string `json:"user_name"`
	ProfilePhoto string `json:"profile_photo"`
}

type ResponseWithBlockedUserList struct {
	Status bool                 `json:"status"`
	Data   []BlockedUserDetails `json:"data"`
}
