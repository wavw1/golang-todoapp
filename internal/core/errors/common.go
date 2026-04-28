package core_errors

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrInvalidArgument    = errors.New("invalid argument")
	ErrConflict           = errors.New("conflict")
	ErrViolatesForeignKey = errors.New("violates foreign key")
	// ErrUnknown            = errors.New("err unknown")
)

const (
	ForeignKeyViolation = "23503"
)
