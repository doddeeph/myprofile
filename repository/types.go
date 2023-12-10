// This file contains types that are used in the repository layer.
package repository

type CreateUserParams struct {
	FullName    string `json:"full_name"`
	Password    string `json:"password"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
}
