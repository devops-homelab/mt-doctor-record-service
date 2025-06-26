package handler

import (
	"doctor-record-service/db"
	"doctor-record-service/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type DoctorHandler struct{}

// CreateDoctorHandler handles the creation of a new doctor
func (h *DoctorHandler) CreateDoctorHandler(w http.ResponseWriter, r *http.Request) {
	var doctor model.Doctor
	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := db.AddDoctor(doctor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	doctor.ID = id
	message := map[string]string{
		"message": "Doctor " + id + " is added successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}

// GetDoctorHandler handles retrieving a doctor by ID
func (h *DoctorHandler) GetDoctorHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	doctor, err := db.GetDoctor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(doctor)
}

// UpdateDoctorHandler handles updating a doctor by ID
func (h *DoctorHandler) UpdateDoctorHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var doctor model.Doctor
	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	doctor.ID = id
	if err := db.UpdateDoctor(doctor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := map[string]string{
		"message": "Doctor " + id + " is updated successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}

// GetAllDoctorsHandler handles retrieving all doctors
func (h *DoctorHandler) GetAllDoctorsHandler(w http.ResponseWriter, r *http.Request) {
	doctors, err := db.GetAllDoctors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(doctors)
}

// DeleteDoctorHandler handles deleting a doctor by ID
func (h *DoctorHandler) DeleteDoctorHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := db.DeleteDoctor(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := map[string]string{
		"message": "Doctor " + id + " is deleted successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}
