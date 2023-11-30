package entities

type Meta struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Label     string `json:"label"`
}

type Summary struct {
	Meta
	Credit float64 `json:"credit"`
	Debit  float64 `json:"debit"`
}

type TransactionSummary struct {
	Meta
	Data []TransactionSummaryRecord `json:"data"`
}

type TransactionSummaryRecord struct {
	TransactionID int64    `json:"transaction_id"`
	Account       Account  `json:"account"`
	Category      Category `json:"category"`
	Amount        float64  `json:"amount"`
	Charges       float64  `json:"charges"`
	Timestamp     string   `json:"timestamp"`
}

type AccountSummary struct {
	Meta
	Data []AccountSummaryRecord `json:"data"`
}

type AccountSummaryRecord struct {
	AccountID        int64   `json:"account_id"`
	AccountName      string  `json:"account_name"`
	Balance          float64 `json:"balance"`
	Credit           float64 `json:"credit"`
	Debit            float64 `json:"debit"`
	CreditPercentage float64 `json:"credit_percentage"`
	DebitPercentage  float64 `json:"debit_percentage"`
}

type CategorySummary struct {
	Meta
	Data []CategorySummaryRecord `json:"data"`
}

type CategorySummaryRecord struct {
	CategoryID       int64   `json:"category_id"`
	CategoryName     string  `json:"category_name"`
	Credit           float64 `json:"credit"`
	Debit            float64 `json:"debit"`
	CreditPercentage float64 `json:"credit_percentage"`
	DebitPercentage  float64 `json:"debit_percentage"`
}
