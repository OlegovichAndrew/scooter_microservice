package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"scooter_micro/config"
	"scooter_micro/proto"
	"scooter_micro/repository"
	"scooter_micro/routing"
	"scooter_micro/routing/grpcserver"
	"scooter_micro/routing/httpserver"
	"scooter_micro/service"
)

func main() {
	//connStr := "host=localhost port=5444 user=scooteradmin password=Megascooter! dbname=scooterdb sslmode=disable"
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.PG_HOST,
		config.PG_PORT,
		config.POSTGRES_USER,
		config.POSTGRES_PASSWORD,
		config.POSTGRES_DB)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panicf("%s: failed to open db connection - %v", "scooter_micro", err)
	}
	defer db.Close()

	scooterRepo := repository.NewScooterRepo(db)
	scooterService := service.NewScooterService(scooterRepo)

	handler := routing.NewRouter(scooterService)
	httpServer := httpserver.New(handler, httpserver.Port("8080"))
	handler.HandleFunc("/scooter", httpServer.ScooterHandler)
	grpcServer := grpcserver.NewGrpcServer()
	proto.RegisterScooterServiceServer(grpcServer, httpServer)
	http.ListenAndServe(":8080", handler)
}
