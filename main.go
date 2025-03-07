package main

import (
	"doctor-record-service/db"
	"doctor-record-service/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection
	db.InitDB()
	log.Println("Database connected")

	r := mux.NewRouter()

	doctorHandler := &handler.DoctorHandler{}

	// Create a new doctor
	r.HandleFunc("/doctors", doctorHandler.CreateDoctorHandler).Methods("POST")

	// Retrieve a doctor by ID
	r.HandleFunc("/doctors/{id}", doctorHandler.GetDoctorHandler).Methods("GET")

	// Update a doctor by ID
	r.HandleFunc("/doctors/{id}", doctorHandler.UpdateDoctorHandler).Methods("PUT")

	// Retrieve all doctors
	r.HandleFunc("/doctors", doctorHandler.GetAllDoctorsHandler).Methods("GET")

	// Delete a doctor by ID
	r.HandleFunc("/doctors/{id}", doctorHandler.DeleteDoctorHandler).Methods("DELETE")

	// Start the server on port 8084
	log.Println("Starting server on :8084")
	log.Fatal(http.ListenAndServe(":8084", r))
}
