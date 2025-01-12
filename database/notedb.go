package database

import (
	"BackendAndroid/models"
	"database/sql"
	"fmt"
	"log"
)

func SaveNote(db *sql.DB, note models.Note) error {
	query := `
		INSERT INTO notes (id, title, content, image_base64, username)
		VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			title = VALUES(title),
			content = VALUES(content),
			image_base64 = VALUES(image_base64),
			username = VALUES(username)
	`
	log.Println(note.Id, note.Title)
	_, err := db.Exec(query, note.Id, note.Title, note.Content, note.ImageBase64, note.Username)
	if err != nil {
		return fmt.Errorf("failed to save note: %v", err)
	}
	return nil
}

func DeleteNoteByID(id int) error {
	db := DB
	query := "DELETE FROM notes WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete note: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no note found with id %d", id)
	}

	fmt.Printf("Successfully deleted note with id %d\n", id)
	return nil
}

func GetNotes() ([]models.Note, error) {
	db := DB
	query := "SELECT id, title, content, image_base64, username FROM notes"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query notes: %v", err)
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		err := rows.Scan(&note.Id, &note.Title, &note.Content, &note.ImageBase64, &note.Username)
		if err != nil {
			return nil, fmt.Errorf("failed to scan note: %v", err)
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after scanning rows: %v", err)
	}

	return notes, nil
}

func GetNotesByUsername(username string) ([]models.Note, error) {
	db := DB
	query := "SELECT id, title, content, image_base64, username FROM notes WHERE username = ?"
	rows, err := db.Query(query, username)
	if err != nil {
		return nil, fmt.Errorf("failed to query notes: %v", err)
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		err := rows.Scan(&note.Id, &note.Title, &note.Content, &note.ImageBase64, &note.Username)
		if err != nil {
			return nil, fmt.Errorf("failed to scan note: %v", err)
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after scanning rows: %v", err)
	}
	return notes, nil
}
