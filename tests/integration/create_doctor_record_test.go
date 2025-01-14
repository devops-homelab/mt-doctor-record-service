package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestAddDoctor(t *testing.T) {
	baseURL := "http://preview.doctor-records.apps.meditrack-app.me/doctors"

	// Request payload for adding a new doctor
	newDoctor := map[string]interface{}{
		"first_name":   "Dr. Emily",
		"last_name":    "Smith",
		"specialty":    "Pediatrician",
		"phone_number": "+1-234-567-8901",
		"email":        "emilysmith@example.com",
	}

	payload, err := json.Marshal(newDoctor)
	if err != nil {
		t.Fatalf("Failed to marshal request payload: %v", err)
	}

	// Create a new doctor
	resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to make POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		t.Fatalf("Expected status code 201 or 200, got %d. Response body: %s", resp.StatusCode, string(bodyBytes))
	}

	// Verify the added doctor's details
	var createdDoctor map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&createdDoctor); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

}
