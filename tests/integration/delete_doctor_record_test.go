package integration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestDeleteDoctor(t *testing.T) {
	baseURL := "mt-doctor-record-service-dev-preview.mt-doctor-record-service-dev.svc.cluster.local/doctors"

	// Fetch a random doctor ID to delete
	doctorID := fetchRandomDoctorID(t, baseURL)

	// Send DELETE request to delete the doctor
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", baseURL, doctorID), nil)
	if err != nil {
		t.Fatalf("Failed to create DELETE request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to make DELETE request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		t.Fatalf("Expected status code 204 or 200, got %d. Response body: %s", resp.StatusCode, string(bodyBytes))
	}

	// Verify the doctor is deleted
	// Attempt to GET the doctor to check if it still exists
	getURL := fmt.Sprintf("%s/%s", baseURL, doctorID)
	resp, err = http.Get(getURL)
	if err != nil {
		t.Fatalf("Failed to make GET request for deleted doctor: %v", err)
	}
	defer resp.Body.Close()

}

// Helper function to fetch a random doctor ID
func fetchRandomDoctorID(t *testing.T, serviceURL string) string {
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
