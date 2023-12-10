package repository

import "context"

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
