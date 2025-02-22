package user

// User представляет сущность пользователя.
type User struct {
	id       int
	Username string
	email    string
}

// NewUser создает нового пользователя.
func NewUser(id int, username, email string) *User {
	return &User{
		id:       id,
		Username: username,
		email:    email,
	}
}

// GetID возвращает ID пользователя.
func (u *User) GetID() int {
	return u.id
}

// GetEmail возвращает email пользователя.
func (u *User) GetEmail() string {
	return u.email
}

// SetEmail обновляет email пользователя.
func (u *User) SetEmail(email string) {
	u.email = email
}
