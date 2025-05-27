package dto

import "github.com/diemensa/denezhki/internal/repository/postgres/model"

type BalanceRequest struct {
	Balance float64 `json:"balance"`
}
type BalanceResponse struct {
	Balance float64 `json:"balance"`
}
type CreateAccountRequest struct {
	Alias string `json:"alias"`
}

type AccountResponse struct {
	ID      string  `json:"id"`
	UserID  string  `json:"user_id"`
	Alias   string  `json:"alias"`
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

func NewAccountResponse(acc *model.Account) AccountResponse {
	return AccountResponse{
		ID:      acc.ID.String(),
		UserID:  acc.UserID.String(),
		Alias:   acc.Alias,
		Owner:   acc.Owner,
		Balance: acc.Balance,
	}
}
