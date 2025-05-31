package dto

import "github.com/google/uuid"

type TransferRequest struct {
	ToID   uuid.UUID `json:"to_id" binding:"required"`
	Amount float64   `json:"amount" binding:"required,gte=1"`
}

type TransferResponse struct {
	TransferID uuid.UUID `json:"transfer_id"`
	Message    string    `json:"message"`
}

func NewTransferResponse(id uuid.UUID, success bool) *TransferResponse {
	msg := "transfer failed"

	if success {
		msg = "transfer successful"
	}

	return &TransferResponse{
		TransferID: id,
		Message:    msg,
	}
}
