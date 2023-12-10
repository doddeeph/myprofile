package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/utils"
	passwordValidator "github.com/go-passwd/validator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) UserRegistration(ctx echo.Context) error {
	var request generated.RegistrationRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Code: strconv.Itoa(http.StatusBadRequest), Message: err.Error(),
		})
	}
	validate := validator.New()
	validatePasswd := passwordValidator.New(
		passwordValidator.ContainsAtLeast("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1,
			errors.New("Password must contains at least 1 capital character")),
		passwordValidator.ContainsAtLeast("1234567890", 1,
			errors.New("Password must contains at least 1 number character")),
		passwordValidator.ContainsAtLeast("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~", 1,
			errors.New("Password must contains at least 1 special character")))
	err1 := validate.Struct(request)
	err2 := validatePasswd.Validate(request.Password)
	if err1 != nil || err2 != nil {
		var message string
		if err1 != nil {
			errs := err1.(validator.ValidationErrors)
			if len(errs) != 0 {
				errMsg := make([]string, len(errs))
				for i, fieldErr := range errs {
					errMsg[i] = fmt.Sprintf("%s must be at %s %s characters", fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
				}
				message = strings.Join(errMsg, ", ")
			}
		}
		if err2 != nil {
			if message != "" {
				message += ", "
			}
			message += err2.Error()
		}
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Code: strconv.Itoa(http.StatusBadRequest), Message: message,
		})
	}
	user, err := s.Repository.CreateUser(context.Background(), repository.CreateUserParams{
		FullName:    request.FullName,
		CountryCode: request.CountryCode,
		PhoneNumber: request.PhoneNumber,
		Password:    utils.HashAndSaltPassword(request.Password),
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{
			Code: strconv.Itoa(http.StatusInternalServerError), Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, generated.RegistrationResponse{
		Id: &user.ID,
	})
}
