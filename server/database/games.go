package database

type Game struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}

func (database *Database) CreateGame(id string, data string) (Game, error) {
	_, err := database.DB.Exec(`INSERT INTO games VALUES (?, ?)`,
		id, data)
	if err != nil {
		return Game{}, err
	}

	return Game{
		Id:   id,
		Data: data,
	}, nil
}

func (database *Database) GetGame(id string) (Game, error) {
	row := database.DB.QueryRow(`SELECT data FROM games WHERE id = ?`, id)
	var data string
	err := row.Scan(&data)
	if err != nil {
		return Game{}, err
	}
	return Game{
		Id:   id,
		Data: data,
	}, nil
}

func (database *Database) UpdateGame(id string, data string) (Game, error) {
	_, err := database.DB.Exec(`UPDATE games SET data = ? WHERE id = ?`, data, id)
	if err != nil {
		return Game{}, err
	}
	return Game{
		Id:   id,
		Data: data,
	}, nil
}

func (database *Database) DeleteGame(id string) error {
	_, err := database.DB.Exec(`DELETE FROM games WHERE id = ?`, id)
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) GetAllGames() ([]string, error) {
	rows, err := database.DB.Query(`SELECT id FROM games`)
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
