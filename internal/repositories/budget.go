package repositories

import (
	"database/sql"
	"github.com/rapando/budget-api/internal/entities"
)

type IBudgetRepository interface {
	ListAccounts() (accounts []entities.Account, err error)
	ListCategories() (categories []entities.Category, err error)
	AddTransaction(transaction entities.Transaction) (err error)
}

type BudgetRepository struct {
	DB *sql.DB
}

func (r *BudgetRepository) ListAccounts() (accounts []entities.Account, err error) {
	accounts = []entities.Account{}
	var query = `
		SELECT account_id, name, balance, created, modified 
		FROM account`
	var modified sql.NullString
	var account entities.Account
	var rows *sql.Rows

	rows, err = r.DB.Query(query)
	if err != nil {
		return accounts, err
	}

	for rows.Next() {
		err = rows.Scan(
			&account.AccountID,
			&account.Name,
			&account.Balance,
			&account.Created,
			&modified,
		)
		if err != nil {
			return accounts, err
		}
		account.Modified = modified.String
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r *BudgetRepository) ListCategories() (categories []entities.Category, err error) {
	var query = `
		SELECT category_id, name, created, modified
		FROM category`
	var modified sql.NullString
	var category entities.Category
	var rows *sql.Rows

	rows, err = r.DB.Query(query)
	if err != nil {
		return categories, err
	}
	for rows.Next() {
		err = rows.Scan(
			&category.CategoryID,
			&category.Name,
			&category.Name,
			&modified,
		)
		if err != nil {
			return categories, err
		}
		category.Modified = modified.String
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *BudgetRepository) AddTransaction(transaction entities.Transaction) (err error) {
	var tx *sql.Tx
	var stmt *sql.Stmt

	var trxQuery = `
	INSERT INTO transaction (account_id, category_id, transaction_type, description, amount, charges)
	VALUES (?,?,?,?,?,?)`
	var accountQuery = `UPDATE account SET balance=balance+? WHERE account_id=? LIMIT 1`
	if transaction.TransactionType == "Debit" {
		accountQuery = `UPDATE account SET balance=balance-? WHERE account_id=? LIMIT 1`
	}

	tx, err = r.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare(trxQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		transaction.Account.AccountID,
		transaction.Category.CategoryID,
		transaction.TransactionType,
		transaction.Description,
		transaction.Amount,
		transaction.Charges,
	)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	stmt, err = tx.Prepare(accountQuery)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = stmt.Exec(transaction.Charges+transaction.Amount, transaction.Account.AccountID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}
