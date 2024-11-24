package user

const (
	queryInsertNewUser = `
		INSERT INTO users 
		(
			username,
			password,
			full_name
		) VALUES (?, ?, ?) RETURNING id, username, full_name, role
	`

	queryFindUserByUsername = `
			SELECT 
				id, 
				username,
				password,
				full_name,
				role
			FROM users WHERE username = ?
	`
)
