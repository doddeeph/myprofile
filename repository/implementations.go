package repository

import (
	"context"
)

const createUser = `
INSERT INTO users (
    full_name,
    password,
    country_code,
    phone_number
) VALUES (
    $1, $2, $3, $4
) RETURNING id, full_name, password, country_code, phone_number, successful_login, created_at, updated_at
`

func (r *Repository) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := r.Db.QueryRowContext(ctx, createUser,
		arg.FullName,
		arg.Password,
		arg.CountryCode,
		arg.PhoneNumber,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Password,
		&i.CountryCode,
		&i.PhoneNumber,
		&i.SuccessfulLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByPhoneNumber = `
SELECT id, full_name, password, country_code, phone_number, successful_login, created_at, updated_at FROM users
WHERE phone_number = $1 LIMIT 1
`

func (r *Repository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (User, error) {
	row := r.Db.QueryRowContext(ctx, getUserByPhoneNumber, phoneNumber)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Password,
		&i.CountryCode,
		&i.PhoneNumber,
		&i.SuccessfulLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `
SELECT id, full_name, password, country_code, phone_number, successful_login, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (r *Repository) GetUser(ctx context.Context, id int64) (User, error) {
	row := r.Db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Password,
		&i.CountryCode,
		&i.PhoneNumber,
		&i.SuccessfulLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateFullNameUser = `
UPDATE users
SET full_name = $2, updated_at = $3
WHERE id = $1
RETURNING id, full_name, password, country_code, phone_number, successful_login, created_at, updated_at
`

func (r *Repository) UpdateFullNameUser(ctx context.Context, arg UpdateFullNameUserParams) (User, error) {
	row := r.Db.QueryRowContext(ctx, updateFullNameUser, arg.ID, arg.FullName, arg.UpdatedAt)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Password,
		&i.CountryCode,
		&i.PhoneNumber,
		&i.SuccessfulLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePhoneNumberUser = `
UPDATE users
SET country_code = $2, phone_number = $3, updated_at = $4
WHERE id = $1
RETURNING id, full_name, password, country_code, phone_number, successful_login, created_at, updated_at
`

func (r *Repository) UpdatePhoneNumberUser(ctx context.Context, arg UpdatePhoneNumberUserParams) (User, error) {
	row := r.Db.QueryRowContext(ctx, updatePhoneNumberUser,
		arg.ID,
		arg.CountryCode,
		arg.PhoneNumber,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Password,
		&i.CountryCode,
		&i.PhoneNumber,
		&i.SuccessfulLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
