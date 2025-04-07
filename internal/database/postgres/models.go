package postgres

type CreateUserParams struct {
	Email    *string
	Password *string
}

type User struct {
	Id       string
	Email    string
	Password string
}

type Task struct {
	Id          string
	Title       string
	Description string
}

type CreateTaskParams struct {
	Description *string
	Title       *string
}
