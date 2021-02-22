package usecase

import "errors"

var (
	InvalidEmailSessionCredential = errors.New("invalid email credential")
	EmailContentToBig             = errors.New("email content exceeding allowed capacity")
	InvalidTypeErr                = errors.New("invalid type")
)
