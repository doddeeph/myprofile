package handler

import (
	"context"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (s *Server) UserRegistration(ctx echo.Context) error {
	var request generated.RegistrationRequest
	var errors []interface{}
	if err := ctx.Bind(&request); err != nil {
		errors = append(errors, err.Error())
		return ctx.JSON(http.StatusBadRequest, generated.RegistrationResponse{
			Code: string(ERROR), Errors: errors,
		})
	}
	regReqValidator := util.NewRegistrationRequestValidator()
	errors = regReqValidator.Validate(request)
	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, generated.RegistrationResponse{
			Code: string(ERROR), Errors: errors,
		})
	}
	user, err := s.Repository.CreateUser(context.Background(), repository.CreateUserParams{
		FullName:    request.FullName,
		CountryCode: request.CountryCode,
		PhoneNumber: request.PhoneNumber,
		Password:    util.HashAndSaltPassword(request.Password),
	})
	if err != nil {
		errors = append(errors, err.Error())
		return ctx.JSON(http.StatusInternalServerError, generated.RegistrationResponse{
			Code: string(ERROR), Errors: errors,
		})
	}
	data := struct {
		Id *int64 `json:"id,omitempty"`
	}{Id: &user.ID}
	return ctx.JSON(http.StatusCreated, generated.RegistrationResponse{
		Code: string(SUCCESS), Data: data,
	})
}

func (s *Server) GetUser(ctx echo.Context, id int64) error {
	var errors []interface{}
	user, err := s.Repository.GetUser(context.Background(), id)
	if err != nil {
		errors = append(errors, "User not found")
		return ctx.JSON(http.StatusNotFound, generated.GetUserResponse{
			Code: string(ERROR), Errors: errors,
		})
	}
	phoneNumber := user.CountryCode + user.PhoneNumber
	data := struct {
		FullName    *string `json:"fullName,omitempty"`
		Id          *int64  `json:"id,omitempty"`
		PhoneNumber *string `json:"phoneNumber,omitempty"`
	}{Id: &user.ID, FullName: &user.FullName, PhoneNumber: &phoneNumber}
	return ctx.JSON(http.StatusCreated, generated.GetUserResponse{
		Code: string(SUCCESS), Data: data,
	})
}

func (s *Server) UpdateUser(ctx echo.Context, id int64) error {
	var request generated.UpdateUserRequest
	var errors []interface{}
	if err := ctx.Bind(&request); err != nil {
		errors = append(errors, err.Error())
		return ctx.JSON(http.StatusBadRequest, generated.UpdateUserResponse{
			Code: string(ERROR), Errors: errors,
		})
	}
	regReqValidator := util.NewUpdateRequestValidator()
	errors = regReqValidator.Validate(request)
	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, generated.UpdateUserResponse{
			Code: string(ERROR), Errors: errors,
		})
	}
	var user repository.User
	var err error
	if request.FullName != "" {
		user, err = s.Repository.UpdateFullNameUser(context.Background(), repository.UpdateFullNameUserParams{
			ID:        id,
			FullName:  request.FullName,
			UpdatedAt: time.Now(),
		})
	} else {
		user, err = s.Repository.UpdatePhoneNumberUser(context.Background(), repository.UpdatePhoneNumberUserParams{
			ID:          id,
			CountryCode: request.CountryCode,
			PhoneNumber: request.PhoneNumber,
			UpdatedAt:   time.Now(),
		})
	}
	if err != nil {
		errors = append(errors, err.Error())
		return ctx.JSON(http.StatusInternalServerError, generated.UpdateUserResponse{
			Code: string(ERROR), Errors: errors,
		})
	}
	phoneNumber := user.CountryCode + user.PhoneNumber
	data := struct {
		FullName    *string `json:"fullName,omitempty"`
		Id          *int64  `json:"id,omitempty"`
		PhoneNumber *string `json:"phoneNumber,omitempty"`
	}{Id: &user.ID, FullName: &user.FullName, PhoneNumber: &phoneNumber}
	return ctx.JSON(http.StatusOK, generated.UpdateUserResponse{
		Code: string(SUCCESS), Data: data,
	})
}
