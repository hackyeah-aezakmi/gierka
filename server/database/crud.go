package database

import "database/sql"

func Create(db *sql.DB, data string) error {
	_, err := db.Exec(`INSERT INTO games (data) VALUES (?)`,
		data)
	if err != nil {
		return err
	}
	return nil
}

func Read(db *sql.DB, id string) (string, error) {
	row := db.QueryRow(`SELECT data FROM games WHERE id = ?`, id)
	var data string
	err := row.Scan(&data)
	if err != nil {
		return "", err
	}
	return data, nil
}

func Update(db *sql.DB, id string, data string) error {
	_, err := db.Exec(`UPDATE games SET data = ? WHERE id = ?`, data, id)
	if err != nil {
		return err
	}
	return nil
}

func Delete(db *sql.DB, id string) error {
	_, err := db.Exec(`DELETE FROM games WHERE id = ?`, id)
	if err != nil {
		return err
	}
	return nil
}

func List(db *sql.DB) ([]string, error) {
	rows, err := db.Query(`SELECT id FROM games`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}
