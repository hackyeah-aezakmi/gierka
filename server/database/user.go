package database

type User struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}

func (database *Database) CreateUser(id string, data string) (User, error) {
	_, err := database.DB.Exec(`INSERT INTO users VALUES (?, ?)`,
		id, data)
	if err != nil {
		return User{}, err
	}

	return User{
		Id:   id,
		Data: data,
	}, nil
}

func (database *Database) GetUser(id string) (User, error) {
	row := database.DB.QueryRow(`SELECT data FROM users WHERE id = ?`, id)
	var data string
	err := row.Scan(&data)
	if err != nil {
		return User{}, err
	}
	return User{
		Id:   id,
		Data: data,
	}, nil
}

func (database *Database) UpdateUser(id string, data string) (User, error) {
	_, err := database.DB.Exec(`UPDATE users SET data = ? WHERE id = ?`, data, id)
	if err != nil {
		return User{}, err
	}
	return User{
		Id:   id,
		Data: data,
	}, nil
}

func (database *Database) DeleteUser(id string) error {
	_, err := database.DB.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) GetAllUsers() ([]string, error) {
	rows, err := database.DB.Query(`SELECT id FROM users`)
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
