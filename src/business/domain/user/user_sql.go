package user

const (
	readUser = `SELECT
		id,
		username,
		full_name,
		job,
		age
	FROM user
	WHERE id = ? AND status = 1
	`

	readUserList = `SELECT
	id,
	username,
	full_name,
	job,
	age
FROM user
WHERE status = 1
`

	userCount = `SELECT COUNT(*) FROM user WHERE status = 1`
)
