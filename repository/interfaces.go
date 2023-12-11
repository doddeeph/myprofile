// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)

	GetUser(ctx context.Context, id int64) (User, error)

	UpdateFullNameUser(ctx context.Context, arg UpdateFullNameUserParams) (User, error)

	UpdatePhoneNumberUser(ctx context.Context, arg UpdatePhoneNumberUserParams) (User, error)
}
