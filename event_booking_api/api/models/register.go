package models

import (
	"fmt"

	"github.com/dreking/event-booking-api/db"
)

type Registration struct {
	ID      int64
	UserID  int64 `binding:"required"`
	EventID int64 `binding:"required"`
}

func (r *Registration) Save() error {
	query := `
	INSERT INTO registrations(user_id, event_id)
	VALUES (?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(r.UserID, r.EventID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	r.ID = id

	return err
}

func FindUserRegistration(userId, eventId int64) (*Registration, error) {
	query := `
	SELECT * FROM registrations
	WHERE user_id=? AND event_id=?
	`

	row := db.DB.QueryRow(query, userId, eventId)
	fmt.Println(row)

	var registration Registration
	err := row.Scan(&registration.ID, &registration.UserID, &registration.EventID)
	if err != nil {
		return nil, err
	}

	return &registration, nil
}

func CancelRegistration(userId, eventId int64) error {
	query := `
	DELETE FROM registrations
	WHERE user_id=? AND event_id=?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId, eventId)
	return err
}
