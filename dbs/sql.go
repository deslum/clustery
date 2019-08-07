package dbs

import (
	"database/sql"

	"github.com/golang/glog"
)

type SQL struct {
	*sql.DB
}

func NewSQL(db *sql.DB) *SQL {
	return &SQL{
		db,
	}
}

func (o *SQL) GetData(query string) ([][]float64, error) {
	var result = make([][]float64, 0)

	rows, err := o.Query(query)
	if err != nil {
		return make([][]float64, 0), err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			field1 float64
			field2 float64
		)
		if err := rows.Scan(&field1, &field2); err != nil {
			return make([][]float64, 0), err
		}

		if field1 == 0 && field2 == 0 {
			glog.Warningf("field1 %v and field 2 %v equal 0", field1, field2)
		}

		result = append(result, []float64{field1, field2})
	}

	return result, nil
}
