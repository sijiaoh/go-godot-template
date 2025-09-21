package serializers

import "github.com/sijiaoh/go-godot-template/api_server/models"

type UserSerializer struct {
	ID   int     `json:"id"`
	Name *string `json:"name,omitempty"`
}

func NewUserSerializer(user *models.User) *UserSerializer {
	return &UserSerializer{
		ID:   user.ID,
		Name: user.Name,
	}
}
