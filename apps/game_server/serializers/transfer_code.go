package serializers

import "github.com/sijiaoh/go-godot-template/game_server/models"

type TransferCodeSerializer struct {
	Code string `json:"code"`
}

func NewTransferCodeSerializer(ts *models.TransferCode) *TransferCodeSerializer {
	return &TransferCodeSerializer{
		Code: ts.Code,
	}
}
