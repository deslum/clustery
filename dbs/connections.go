package dbs

import (
	"database/sql"
	"fmt"

	"github.com/kshvakov/clickhouse"
)

func Connect(driver string, str string) (*sql.DB, error) {
	connect, err := sql.Open(driver, str)
	if err != nil {
		return nil, err
	}

	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			err := fmt.Errorf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
			return nil, err
		}

		return nil, err
	}

	return connect, nil
}
