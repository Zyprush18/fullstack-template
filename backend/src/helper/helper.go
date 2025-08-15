package helper

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)


func IsDuplicate(err error) bool  {
	var pgErr  = &pgconn.PgError{
		Code: "23505",
	}
	return errors.As(err, &pgErr)
}