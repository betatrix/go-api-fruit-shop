package errors

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

func IsDuplicateEntry(err error) bool {
	if err == nil {
		return false
	}

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mysqlErr.Number == 1062
	}

	return false
}
