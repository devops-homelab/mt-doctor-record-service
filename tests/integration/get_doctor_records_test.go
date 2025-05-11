package integration

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestGetAllDoctors(t *testing.T) {
	baseURL := "http://doctor-record-service-dev-preview.doctor-record-service-dev.svc.cluster.local:8084/doctors"

	// Send GET request to retrieve all doctors
	resp, err := http.Get(baseURL)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		t.Fatalf("Expected status code 200, got %d. Response body: %s", resp.StatusCode, string(bodyBytes))
	}

	// Decode response body to get list of doctors
	var doctors []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&doctors); err != nil {
		t.Fatalf("Failed to decode doctor list: %v", err)
	}

	if len(doctors) == 0 {
		t.Fatalf("No doctors found")
	}

}
