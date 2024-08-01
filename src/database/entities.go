package database

import "gorm.io/gorm"

type OAuth2 struct {
	gorm.Model
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	gorm.Model
	UserId string `json:"user_id"`
}

type Guild struct {
	gorm.Model
	GuildId string `json:"guild_id"`
	Prefix  string `json:"prefix"`
}
