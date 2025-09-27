package serializers

import "github.com/sijiaoh/go-godot-template/api_server/models"

type ClientSessionSerializer struct {
	Token string `json:"token"`
}

func NewClientSessionSerializer(cs *models.ClientSession) *ClientSessionSerializer {
	return &ClientSessionSerializer{
		Token: cs.Token,
	}
}
