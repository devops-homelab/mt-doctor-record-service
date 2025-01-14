package service

import (
	"doctor-record-service/model"
	"patient-record-service/db"
)

// DoctorService provides doctor-related services
type DoctorService struct{}

// AddDoctor adds a new doctor and returns the doctor ID
func (s *DoctorService) AddDoctor(doctor model.Doctor) (string, error) {
	return db.AddDoctor(doctor)
}

// GetDoctor retrieves a doctor by ID
func (s *DoctorService) GetDoctor(id string) (model.Doctor, error) {
	return db.GetDoctor(id)
}

// UpdateDoctor updates an existing doctor's record
func (s *DoctorService) UpdateDoctor(doctor model.Doctor) error {
	return db.UpdateDoctor(doctor)
}

// GetAllDoctors retrieves all doctors
func (s *DoctorService) GetAllDoctors() ([]model.Doctor, error) {
	return db.GetAllDoctors()
}

// DeleteDoctor deletes a doctor by ID
func (s *DoctorService) DeleteDoctor(id string) error {
	return db.DeleteDoctor(id)
}
