package dto

type CreateAccountRequest struct {
	Username string `json:"username" binding:"required"`
}

type BalanceResponse struct {
	Balance float64 `json:"balance"`
}
