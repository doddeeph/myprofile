package repository

import (
	"database/sql"
	"time"
)

type User struct {
	ID              int64         `json:"id"`
	FullName        string        `json:"full_name"`
	Password        string        `json:"password"`
	CountryCode     string        `json:"country_code"`
	PhoneNumber     string        `json:"phone_number"`
	SuccessfulLogin sql.NullInt32 `json:"successful_login"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
