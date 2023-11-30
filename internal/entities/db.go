package entities

type Account struct {
	AccountID int64   `json:"account_id"`
	Name      string  `json:"name"`
	Balance   float64 `json:"balance,omitempty"`
	Created   string  `json:"created,omitempty"`
	Modified  string  `json:"modified,omitempty"`
}

type Category struct {
	CategoryID int64  `json:"category_id"`
	Name       string `json:"name"`
	Created    string `json:"created,omitempty"`
	Modified   string `json:"modified,omitempty"`
}

type Transaction struct {
	TransactionID   int64    `json:"transaction_id"`
	Account         Account  `json:"account"`
	Category        Category `json:"category"`
	TransactionType string   `json:"transaction_type"`
	Description     string   `json:"description"`
	Amount          float64  `json:"amount"`
	Charges         float64  `json:"charges"`
	Created         string   `json:"created"`
	Modified        string   `json:"modified,omitempty"`
}
