package sync

import (
	"AutoSyncGO/code/models"
	"database/sql"
	"time"
)

func Detector(db *sql.DB, lastSync time.Time) ([]models.Change, error) {
	query := `
		SELECT id, file_path, change_type, changed_at
		FROM sync_changes
		WHERE changed_at > ?
		ORDER BY changed_at ASC
	`

	rows, err := db.Query(query, lastSync)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var changes []Change
	for rows.Next() {
		var c Change
		if err := rows.Scan(&c.ID, &c.FilePath, &c.ChangeType, &c.ChangedAt); err != nil {
			return nil, err
		}
		changes = append(changes, c)
	}

	return changes, nil
}
