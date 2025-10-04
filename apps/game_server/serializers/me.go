package serializers

import "github.com/sijiaoh/go-godot-template/game_server/models"

type MeSerializer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewMeSerializer(user *models.User) *MeSerializer {
	return &MeSerializer{
		ID:   user.ID,
		Name: user.Name,
	}
}
