package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestUpdateDoctor(t *testing.T) {
	baseURL := "http://preview.doctor-records.apps.meditrack-app.me/doctors"
	doctorID := fetchRandomDoctor(t, baseURL)

	// Request payload for updating the doctor
	updatedDoctor := map[string]interface{}{
		"first_name":   "Dr. Emily",
		"last_name":    "Smith",
		"specialty":    "Pediatrician",
		"phone_number": "+1-234-567-8901",
		"email":        "emilysmith@example.com",
	}

	payload, err := json.Marshal(updatedDoctor)
	if err != nil {
		t.Fatalf("Failed to marshal request payload: %v", err)
	}

	// Send PUT request to update the doctor
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", baseURL, doctorID), bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to create PUT request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to make PUT request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		t.Fatalf("Expected status code 200, got %d. Response body: %s", resp.StatusCode, string(bodyBytes))
	}

	// Verify the updated doctor details
	var updatedResp map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&updatedResp); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

}

// Helper function to fetch a random doctor ID
func fetchRandomDoctor(t *testing.T, serviceURL string) string {
	t.Helper()
	var doctors []map[string]interface{}
	resp, err := http.Get(serviceURL)
	if err != nil {
		t.Fatalf("Failed to fetch doctors: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Failed to fetch doctors, got status code %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&doctors); err != nil {
		t.Fatalf("Failed to decode doctor list: %v", err)
	}

	if len(doctors) == 0 {
		t.Fatalf("No doctors found")
	}

	doctorID, ok := doctors[0]["id"].(string)
	if !ok || doctorID == "" {
		t.Fatalf("Invalid doctor ID retrieved")
	}

	return doctorID
}
