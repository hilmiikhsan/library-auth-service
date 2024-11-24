package helpers

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func SeedAdminAccount(db *sqlx.DB) error {
	var existingAdminID uuid.UUID
	err := db.Get(&existingAdminID, "SELECT id FROM users WHERE username = $1", "admin")
	if err == nil {
		Logger.Error("Admin user already exists.")
		return nil // Admin already exists, nothing to do
	}

	hashedPassword, err := HashPassword("admin123") // Default password is 'admin'
	if err != nil {
		Logger.Error("Error hashing the password:", err)
		return err
	}

	query := `
		INSERT INTO users (username, password, full_name, role)
		VALUES ($1, $2, $3, $4)
	`

	values := []interface{}{
		"admin",         // Username
		hashedPassword,  // Hashed password
		"Administrator", // Full name
		"Admin",         // Role
	}

	_, err = db.Exec(query, values...)
	if err != nil {
		Logger.Println("Error creating admin user:", err)
		return err
	}

	Logger.Println("Admin user created successfully.")
	return nil
}
