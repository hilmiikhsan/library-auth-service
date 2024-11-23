package user_session

const (
	queryFindUserSessionByToken = `
		SELECT
			id,
			user_id,
			token,
			refresh_token,
			token_expired,
			refresh_token_expired
		FROM user_sessions 
		WHERE token = ?
	`

	queryInsertNewUserSession = `
		INSERT INTO user_sessions
		(
			user_id,
			token,
			refresh_token,
			token_expired,
			refresh_token_expired
		) VALUES (?, ?, ?, ?, ?)
	`

	queryFindUserSessionByRefreshToken = `
		SELECT
			id,
			user_id,
			token,
			refresh_token,
			token_expired,
			refresh_token_expired
		FROM user_sessions 
		WHERE refresh_token = ?
	`

	queryDeleteUserSession = `
		DELETE FROM user_sessions WHERE token = ?
	`

	queryUpdateTokenByRefreshToken = `
		UPDATE user_sessions
		SET
			token = ?
		WHERE refresh_token = ?
	`
)
