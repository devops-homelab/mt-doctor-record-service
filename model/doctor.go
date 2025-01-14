package model

type Doctor struct {
	ID          string `json:"id"`           // Unique identifier (UUID)
	FirstName   string `json:"first_name"`   // Doctor's first name
	LastName    string `json:"last_name"`    // Doctor's last name
	Specialty   string `json:"specialty"`    // Medical specialty (e.g., Cardiologist, Dermatologist)
	PhoneNumber string `json:"phone_number"` // Contact number
	Email       string `json:"email"`        // Email address
}
