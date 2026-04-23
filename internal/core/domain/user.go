package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/wavw1/golang-todoapp/internal/core/errors"
)

type User struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func NewUser(
	id int,
	version int,
	fullName string,
	phoneNumber *string,
) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

func NewUserUninitialized(
	fullName string,
	phoneNumber *string,
) User {
	return NewUser(
		UninitializedId,
		UninitializedVersion,
		fullName,
		phoneNumber,
	)
}

func (u *User) Validate() error {
	fullnameLength := len([]rune(u.FullName))
	if fullnameLength < 3 || fullnameLength > 100 {
		return fmt.Errorf(
			"invalid `FullName` len: %d: %w",
			fullnameLength,
			core_errors.ErrInvalidArgument,
		)
	}

	if u.PhoneNumber != nil {
		phoneNumberLen := len([]rune(*u.PhoneNumber))
		if phoneNumberLen < 10 || phoneNumberLen > 15 {
			return fmt.Errorf(
				"invalid `PhoneNumber` len: %d: %w",
				phoneNumberLen,
				core_errors.ErrInvalidArgument,
			)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf(
				"invalid `PhoneNumber` format: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}
