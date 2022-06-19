package tokens

import (
	"github.com/NetworkPy/TestTaskJuniorBackDev/internal/users"
)

type Session struct {
	RefreshToken string
	User         users.User
}
