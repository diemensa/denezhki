package dto

type BalanceRequest struct {
	Balance float64 `json:"balance" binding:"required"`
}
type BalanceResponse struct {
	Balance float64 `json:"balance"`
}

type CreateAccountRequest struct {
	Alias string `json:"alias"`
}
