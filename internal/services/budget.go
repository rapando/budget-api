package services

import (
	"github.com/rapando/budget-api/internal/entities"
	"github.com/rapando/budget-api/internal/repositories"
	"github.com/rapando/budget-api/pkg/log"
)

type IBudgetService interface {
	ListAccounts() (accounts []entities.Account, err error)
	ListCategories() (categories []entities.Category, err error)
	AddTransaction(transaction entities.Transaction) (err error)
}

type BudgetService struct {
	*repositories.BudgetRepository
	*repositories.SummaryRepository
}

func (s *BudgetService) ListAccounts() (accounts []entities.Account, err error) {
	log.Infof("listing accounts")
	accounts, err = s.BudgetRepository.ListAccounts()
	if err != nil {
		log.Warnf("failed to fetch accounts because %v", err)
		return accounts, err
	}
	log.Infof("got %d accounts", len(accounts))
	log.Debugf("accounts: %+v", accounts)
	return accounts, err
}

func (s *BudgetService) ListCategories() (categories []entities.Category, err error) {
	log.Infof("listing categories")
	categories, err = s.BudgetRepository.ListCategories()
	if err != nil {
		log.Warnf("failed to fetch categories because %v", err)
		return categories, err
	}
	log.Infof("got %d categories", len(categories))
	log.Debugf("categories: %+v", categories)
	return categories, err
}

func (s *BudgetService) AddTransaction(transaction entities.Transaction) (err error) {
	log.Infof("add transaction")
	err = s.BudgetRepository.AddTransaction(transaction)
	if err != nil {
		log.Warnf("failed to add transaction because %v", err)
		return err
	}
	log.Infof("saved transaction successfully")
	return nil
}

func (s *BudgetService) GetSummary(period string) (summary entities.Summary, err error) {
	log.Infof("getting summary for [%s]", period)
	summary, err = s.SummaryRepository.GetSummary(period)
	if err != nil {
		log.Warnf("failed to read summary because %v", err)
		return summary, err
	}
	log.Infof("found summary")
	log.Debugf("summary: %+v", summary)
	return summary, nil
}

func (s *BudgetService) Transactions(period string) (summary entities.TransactionSummary, err error) {
	log.Infof("fetch transactions for [%s]", period)
	summary, err = s.SummaryRepository.Transactions(period)
	if err != nil {
		log.Warnf("failed to load transactions summary because %v", err)
		return summary, err
	}
	log.Infof("found transaction summary")
	log.Debugf("transaction summary: %+v", summary)
	return summary, err
}

func (s *BudgetService) AccountSummary(period string) (summary entities.AccountSummary, err error) {
	log.Infof("fetch account summary for [%s]", period)
	summary, err = s.SummaryRepository.AccountSummary(period)
	if err != nil {
		log.Warnf("failed to load account summary because %v", err)
		return summary, err
	}
	log.Infof("found account summary")
	log.Debugf("account summary: %+v", summary)
	return summary, err
}
func (s *BudgetService) CategorySummary(period string) (summary entities.CategorySummary, err error) {
	log.Infof("fetch category summary for [%s]", period)
	summary, err = s.SummaryRepository.CategorySummary(period)
	if err != nil {
		log.Warnf("failed to load category summary because %v", err)
		return summary, err
	}
	log.Infof("found category summary")
	log.Debugf("category summary: %+v", summary)
	return summary, err
}
