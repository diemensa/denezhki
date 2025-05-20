package dto

import "github.com/google/uuid"

type TransferRequest struct {
	FromID uuid.UUID `json:"from_id" binding:"required"`
	ToID   uuid.UUID `json:"to_id" binding:"required"`
	Amount float64   `json:"amount" binding:"required,gt=0"`
}

type TransferResponse struct {
	TransactionID uuid.UUID `json:"transaction_id"`
	Message       string    `json:"message"`
}
