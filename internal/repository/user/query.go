package user

const (
	queryInsertNewUser = `
		INSERT INTO users 
		(
			username,
			password
		) VALUES (?, ?) RETURNING id, username
	`

	queryFindUserByUsername = `
			SELECT 
				id, 
				username,
				password
			FROM users WHERE username = ?
	`
)
