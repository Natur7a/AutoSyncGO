package sync

import (
	"database/sql"
	"fmt"
)

func Detector(db *sql.DB, lastVersion []byte, tableName string) ([]map[string]interface{}, error) {
	query := fmt.Sprintf(`
		SELECT *
		FROM %s
		WHERE rowversion > ?
		ORDER BY rowversion ASC
	`, tableName)

	rows, err := db.Query(query, lastVersion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(cols))
		valuePtrs := make([]interface{}, len(cols))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if errors := rows.Scan(valuePtrs...); errors != nil {
			return nil, errors
		}

		rowMap := make(map[string]interface{}, len(cols))
		for i, c := range cols {
			rowMap[c] = values[i]
		}
		result = append(result, rowMap)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
