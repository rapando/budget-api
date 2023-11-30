package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rapando/budget-api/internal/handlers/budget"
	"github.com/rapando/budget-api/internal/handlers/summary"
	"github.com/rapando/budget-api/internal/repositories"
	"github.com/rapando/budget-api/internal/services"
	"github.com/rapando/budget-api/pkg/conn"
	logger "github.com/rapando/budget-api/pkg/log"
	"log"
	"net/http"
	"os"
)

func main() {
	var err error
	var db *sql.DB
	var router = chi.NewRouter()

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load dotenv because %v")
	}

	logger.InitLogger(os.Getenv("ENV"))

	db, err = conn.DBConn(os.Getenv("DB_URI"))
	if err != nil {
		log.Fatalf("failed to connect to db because %v", err)
	}

	var budgetService = &services.BudgetService{
		&repositories.BudgetRepository{
			db,
		},
		&repositories.SummaryRepository{
			db,
		},
	}

	budget.RegisterRoutes(budgetService, router)
	summary.RegisterRoutes(budgetService, router)

	var host = fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
	logger.Infof("running API on: %v", host)
	log.Fatalln(http.ListenAndServe(host, router))
}
