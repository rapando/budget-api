package repositories

import (
	"database/sql"
	"github.com/rapando/budget-api/internal/entities"
	"github.com/rapando/budget-api/pkg/fn"
)

type ISummaryRepository interface {
	GetSummary(period string) (summary entities.Summary, err error)
	Transactions(period string) (summary entities.TransactionSummary, err error)
	AccountSummary(period string) (summary entities.AccountSummary, err error)
	CategorySummary(period string) (summary entities.CategorySummary, err error)
}

type SummaryRepository struct {
	DB *sql.DB
}

func (r *SummaryRepository) GetSummary(period string) (summary entities.Summary, err error) {
	var transactionType string
	var amount, charges float64
	var rows *sql.Rows

	summary.StartTime, summary.EndTime, summary.Label = fn.GetPeriodData(period)
	var query = `
		SELECT transaction_type, amount, charges
		FROM transaction
		WHERE created >= ?
			AND created <= ?`

	rows, err = r.DB.Query(query, summary.StartTime, summary.EndTime)
	if err != nil {
		return summary, err
	}
	for rows.Next() {
		err = rows.Scan(
			&transactionType,
			&amount,
			&charges,
		)
		if err != nil {
			return summary, err
		}
		if transactionType == "Credit" {
			summary.Credit += amount + charges
		} else {
			summary.Debit += amount + charges
		}
	}
	return summary, nil
}

func (r *SummaryRepository) Transactions(period string) (summary entities.TransactionSummary, err error) {
	var transaction entities.TransactionSummaryRecord
	var rows *sql.Rows

	summary.Data = []entities.TransactionSummaryRecord{}
	summary.StartTime, summary.EndTime, summary.Label = fn.GetPeriodData(period)
	var query = `
		SELECT t.transaction_id, t.account_id, a.name as account_name, 
			t.category_id, c.name as category_name, t.amount, t.charges, t.created
		FROM transaction t
		INNER JOIN category c ON c.category_id = t.category_id
		INNER JOIN account a ON a.account_id = t.account_id
		WHERE t.created >= ? 
			AND t.created <= ?`
	rows, err = r.DB.Query(query, summary.StartTime, summary.EndTime)
	if err != nil {
		return summary, err
	}

	for rows.Next() {
		err = rows.Scan(
			&transaction.TransactionID,
			&transaction.Account.AccountID,
			&transaction.Account.Name,
			&transaction.Category.CategoryID,
			&transaction.Category.Name,
			&transaction.Amount,
			&transaction.Charges,
			&transaction.Timestamp,
		)
		if err != nil {
			return summary, err
		}
		summary.Data = append(summary.Data, transaction)
	}
	return summary, nil
}
func (r *SummaryRepository) AccountSummary(period string) (summary entities.AccountSummary, err error) {
	var rows *sql.Rows
	var record entities.AccountSummaryRecord
	var totalCredit, totalDebit float64

	summary.Data = []entities.AccountSummaryRecord{}

	summary.StartTime, summary.EndTime, summary.Label = fn.GetPeriodData(period)
	var query = `
		SELECT t.account_id, a.name as account_name, a.balance,
		       SUM(
		       	CASE WHEN t.transaction_type = 'Credit' 
		       	    THEN amount + charges 
		       	    ELSE 0 
		       	END
		       ) as credit,
		    SUM(
		       	CASE WHEN t.transaction_type = 'Debit' 
		       	    THEN amount + charges 
		       	    ELSE 0 
		       	END
		       ) as debit
		FROM transaction t 
		INNER JOIN account a ON a.account_id = t.account_id
		WHERE t.created >= ?
			AND t.created <= ?
		GROUP BY (t.account_id)`
	rows, err = r.DB.Query(query, summary.StartTime, summary.EndTime)
	if err != nil {
		return summary, err
	}
	for rows.Next() {
		err = rows.Scan(
			&record.AccountID,
			&record.AccountName,
			&record.Balance,
			&record.Credit,
			&record.Debit,
		)
		if err != nil {
			return summary, err
		}
		totalCredit += record.Credit
		totalDebit += record.Debit
		summary.Data = append(summary.Data, record)
	}

	for i := 0; i < len(summary.Data); i++ {
		var rec = summary.Data[i]
		if rec.Credit > 0 {
			rec.CreditPercentage = (rec.Credit / totalCredit) * 100.0
		}
		if rec.Debit > 0 {
			rec.DebitPercentage = (rec.Debit / totalDebit) * 100.0
		}
		summary.Data[i] = rec
	}
	return summary, err
}
func (r *SummaryRepository) CategorySummary(period string) (summary entities.CategorySummary, err error) {
	var rows *sql.Rows
	var record entities.CategorySummaryRecord
	var totalCredit, totalDebit float64

	summary.Data = []entities.CategorySummaryRecord{}
	summary.StartTime, summary.EndTime, summary.Label = fn.GetPeriodData(period)

	var query = `
		SELECT t.category_id, c.name as category_name,
			SUM(
		       	CASE WHEN t.transaction_type = 'Credit' 
		       	    THEN amount + charges 
		       	    ELSE 0 
		       	END
		       ) as credit,
		    SUM(
		       	CASE WHEN t.transaction_type = 'Debit' 
		       	    THEN amount + charges 
		       	    ELSE 0 
		       	END
		       ) as debit
		FROM transaction t 
		INNER JOIN category c ON c.category_id = t.category_id
		WHERE t.created >= ?
			AND t.created <= ?
		GROUP BY (t.category_id)`

	rows, err = r.DB.Query(query, summary.StartTime, summary.EndTime)
	if err != nil {
		return summary, err
	}
	for rows.Next() {
		err = rows.Scan(
			&record.CategoryID,
			&record.CategoryName,
			&record.Credit,
			&record.Debit,
		)
		if err != nil {
			return summary, err
		}
		totalCredit += record.Credit
		totalDebit += record.Debit
		summary.Data = append(summary.Data, record)
	}

	for i := 0; i < len(summary.Data); i++ {
		var rec = summary.Data[i]
		if rec.Credit > 0 {
			rec.CreditPercentage = (rec.Credit / totalCredit) * 100.0
		}
		if rec.Debit > 0 {
			rec.DebitPercentage = (rec.Debit / totalDebit) * 100.0
		}
		summary.Data[i] = rec
	}
	return summary, err
}
