package util

import (
	"errors"
	"fmt"
	"github.com/SawitProRecruitment/UserService/generated"
	passwordValidator "github.com/go-passwd/validator"
	"github.com/go-playground/validator/v10"
)

type RegistrationRequestValidator struct {
	validate       *validator.Validate
	validatePasswd *passwordValidator.Validator
}

func NewRegistrationRequestValidator() *RegistrationRequestValidator {
	return &RegistrationRequestValidator{
		validate: validator.New(),
		validatePasswd: passwordValidator.New(
			passwordValidator.ContainsAtLeast("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1,
				errors.New("Password must contains at least 1 capital character")),
			passwordValidator.ContainsAtLeast("1234567890", 1,
				errors.New("Password must contains at least 1 number character")),
			passwordValidator.ContainsAtLeast("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~", 1,
				errors.New("Password must contains at least 1 special character"))),
	}
}

func (v *RegistrationRequestValidator) Validate(request generated.RegistrationRequest) (errors []interface{}) {
	err1 := v.validate.Struct(request)
	err2 := v.validatePasswd.Validate(request.Password)
	if err1 != nil || err2 != nil {
		if err1 != nil {
			errs := err1.(validator.ValidationErrors)
			if len(errs) != 0 {
				for _, fieldErr := range errs {
					errMsg := fmt.Sprintf("%s %s", fieldErr.Field(), fieldErr.Tag())
					if "required" != fieldErr.Tag() {
						errMsg = fmt.Sprintf("%s must be at %s %s characters", fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
					}
					errors = append(errors, errMsg)
				}
			}
		}
		if err2 != nil {
			errors = append(errors, err2.Error())
		}
	}
	return
}

type UpdateRequestValidator struct {
	validate *validator.Validate
}

func NewUpdateRequestValidator() *UpdateRequestValidator {
	return &UpdateRequestValidator{
		validate: validator.New(),
	}
}

func (v *UpdateRequestValidator) Validate(request generated.UpdateUserRequest) (errors []interface{}) {
	err1 := v.validate.Struct(request)
	if err1 != nil {
		errs := err1.(validator.ValidationErrors)
		if len(errs) != 0 {
			for _, fieldErr := range errs {
				errMsg := fmt.Sprintf("%s %s", fieldErr.Field(), fieldErr.Tag())
				if "required" != fieldErr.Tag() {
					errMsg = fmt.Sprintf("%s must be at %s %s characters", fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
				}
				errors = append(errors, errMsg)
			}
		}
	}
	return
}
