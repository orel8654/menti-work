package main

import (
	"fmt"
	"log"
	"menti/pkg/config"
	delivery_api "menti/src/api/internal/delivery"
	repo_api "menti/src/api/internal/repo"
	service_api "menti/src/api/internal/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	if err := run(":8010"); err != nil {
		log.Fatal(err)
	}
}

func run(host string) error {
	conf, err := config.Config("../../../configs/database.yaml")
	if err != nil {
		return err
	}
	s := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", conf.User, conf.Password, conf.DbName, conf.Host, conf.Port)
	db, err := sqlx.Connect("postgres", s)
	if err != nil {
		return err
	}
	repoBasic := repo_api.NewRepo(*db)
	service := service_api.NewService(repoBasic)
	handler := delivery_api.NewHandlers(service)
	return handler.Listen(host)
}
