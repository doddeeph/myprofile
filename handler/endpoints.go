package handler

import (
	"context"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) UserRegistration(ctx echo.Context) error {
	var request generated.RegistrationRequest
	if err := ctx.Bind(&request); err != nil {
		errors := append(make([]interface{}, 1), err.Error())
		return ctx.JSON(http.StatusBadRequest, generated.RegistrationResponse{
			Code: string(ERROR), Errors: &errors,
		})
	}
	regReqValidator := util.NewRegistrationRequestValidator()
	errors := regReqValidator.Validate(request)
	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, generated.RegistrationResponse{
			Code: string(ERROR), Errors: &errors,
		})
	}
	user, err := s.Repository.CreateUser(context.Background(), repository.CreateUserParams{
		FullName:    request.FullName,
		CountryCode: request.CountryCode,
		PhoneNumber: request.PhoneNumber,
		Password:    util.HashAndSaltPassword(request.Password),
	})
	if err != nil {
		errors = append(make([]interface{}, 1), err.Error())
		return ctx.JSON(http.StatusInternalServerError, generated.RegistrationResponse{
			Code: string(ERROR), Errors: &errors,
		})
	}
	data := struct {
		Id *int64 `json:"id,omitempty"`
	}{Id: &user.ID}
	return ctx.JSON(http.StatusCreated, generated.RegistrationResponse{
		Code: string(SUCCESS), Data: &data, Errors: nil,
	})
}
