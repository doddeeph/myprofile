// This file contains types that are used in the repository layer.
package repository

import "time"

type CreateUserParams struct {
	FullName    string `json:"full_name"`
	Password    string `json:"password"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateFullNameUserParams struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"full_name"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatePhoneNumberUserParams struct {
	ID          int64     `json:"id"`
	CountryCode string    `json:"country_code"`
	PhoneNumber string    `json:"phone_number"`
	UpdatedAt   time.Time `json:"updated_at"`
}
