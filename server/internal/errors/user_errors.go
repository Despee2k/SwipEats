package errors

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrEmailAlreadyInUse   = errors.New("email already in use")
	ErrPasswordsDoNotMatch = errors.New("passwords do not match")

	ErrFileCouldNotSave = errors.New("file could not be saved")
	ErrFileCountNotBeDeleted = errors.New("file could not be deleted")
)