package database

// Go 实现
func AddToFavorites(userID, poemID string) error {
	query := `
		INSERT INTO favorites (user_id, poem_id)
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE user_id = user_id;`
	_, err := DB.Exec(query, userID, poemID)
	return err
}

// Go 实现
func RemoveFromFavorites(userID, poemID string) error {
	query := `
		DELETE FROM favorites
		WHERE user_id = $1 AND poem_id = $2;`
	_, err := DB.Exec(query, userID, poemID)
	return err
}

// Go 实现
func GetUserFavorites(userID string) ([]string, error) {
	query := `
		SELECT poem_id
		FROM favorites
		WHERE user_id = ?;`
	rows, err := DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var poemIDs []string
	for rows.Next() {
		var poemID string
		if err := rows.Scan(&poemID); err != nil {
			return nil, err
		}
		poemIDs = append(poemIDs, poemID)
	}
	return poemIDs, nil
}
