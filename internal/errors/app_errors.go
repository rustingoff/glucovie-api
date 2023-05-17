package apperrors

import "errors"

var ErrInvalidUserCredentials = errors.New("invalid user credentials")
var ErrUserAlreadyExists = errors.New("user already exists")
