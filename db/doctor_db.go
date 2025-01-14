package db

import (
	"database/sql"
	"doctor-record-service/model"
	"fmt"

	"github.com/google/uuid"
)

// AddDoctor adds a new doctor to the PostgreSQL database and returns the doctor ID
func AddDoctor(doctor model.Doctor) (string, error) {
	doctor.ID = uuid.New().String()
	query := `INSERT INTO doctors (id, first_name, last_name, specialty, phone_number, email) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := GetDB().Exec(query, doctor.ID, doctor.FirstName, doctor.LastName, doctor.Specialty, doctor.PhoneNumber, doctor.Email)
	if err != nil {
		return "", err
	}
	return doctor.ID, nil
}

// GetDoctor retrieves a doctor by ID from the PostgreSQL database
func GetDoctor(id string) (model.Doctor, error) {
	var doctor model.Doctor
	query := `SELECT id, first_name, last_name, specialty, phone_number, email FROM doctors WHERE id = $1`
	row := GetDB().QueryRow(query, id)
	err := row.Scan(&doctor.ID, &doctor.FirstName, &doctor.LastName, &doctor.Specialty, &doctor.PhoneNumber, &doctor.Email)
	if err == sql.ErrNoRows {
		return model.Doctor{}, fmt.Errorf("doctor not found")
	} else if err != nil {
		return model.Doctor{}, err
	}
	return doctor, nil
}

// UpdateDoctor updates an existing doctor's record in the PostgreSQL database
func UpdateDoctor(doctor model.Doctor) error {
	query := `UPDATE doctors SET first_name = $1, last_name = $2, specialty = $3, phone_number = $4, email = $5 WHERE id = $6`
	_, err := GetDB().Exec(query, doctor.FirstName, doctor.LastName, doctor.Specialty, doctor.PhoneNumber, doctor.Email, doctor.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetAllDoctors retrieves all doctors from the PostgreSQL database
func GetAllDoctors() ([]model.Doctor, error) {
	query := `SELECT id, first_name, last_name, specialty, phone_number, email FROM doctors`
	rows, err := GetDB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []model.Doctor
	for rows.Next() {
		var doctor model.Doctor
		if err := rows.Scan(&doctor.ID, &doctor.FirstName, &doctor.LastName, &doctor.Specialty, &doctor.PhoneNumber, &doctor.Email); err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)
	}
	return doctors, nil
}

// DeleteDoctor deletes a doctor by ID from the PostgreSQL database
func DeleteDoctor(id string) error {
	query := `DELETE FROM doctors WHERE id = $1`
	_, err := GetDB().Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
