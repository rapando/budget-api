package entities

import "fmt"

type AddTransactionRequest struct {
	AccountID       int64   `json:"account_id"`
	CategoryID      int64   `json:"category_id"`
	TransactionType string  `json:"transaction_type"`
	Description     string  `json:"description"`
	Amount          float64 `json:"amount"`
	Charges         float64 `json:"charges"`
}

func (r *AddTransactionRequest) Validate() error {
	if r.AccountID <= 0 {
		return fmt.Errorf("account_id is required")
	}
	if r.CategoryID <= 0 {
		return fmt.Errorf("category_id is required")
	}
	if !(r.TransactionType == "Credit" || r.TransactionType == "Debit") {
		return fmt.Errorf("invalid transaction_type")
	}
	if r.Description == "" {
		return fmt.Errorf("description is required")
	}
	if r.Amount <= 0.0 {
		return fmt.Errorf("amount is required")
	}
	if r.Charges < 0 {
		return fmt.Errorf("charges is required")
	}
	return nil
}

func (r *AddTransactionRequest) ToTransaction() Transaction {
	return Transaction{
		Account:         Account{AccountID: r.AccountID},
		Category:        Category{CategoryID: r.CategoryID},
		TransactionType: r.TransactionType,
		Description:     r.Description,
		Amount:          r.Amount,
		Charges:         r.Charges,
	}
}
