package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"scooter_micro/proto"
	"scooter_micro/repository"
	"scooter_micro/service"
)

func main() {
	connectionString := "host=localhost port=5444 user=scooteradmin password=Megascooter! dbname=scooterdb sslmode" +
		"=disable"
	//connectionDB := fmt.Sprint("host=localhost port=5444 user=scooteradmin password=Megascooter! dbname=scooterdb
	//sslmode
	//=disable",
	//config.PG_HOST,
	//config.PG_PORT,
	//config.POSTGRES_USER,
	//config.POSTGRES_PASSWORD,
	//config.POSTGRES_DB)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panicf("%s: failed to open db connection - %v", "scooter_micro", err)
	}
	defer db.Close()

	scooterRepo := repository.NewScooterRepo(db)
	scooterService := service.NewScooterService(scooterRepo)
	id := &proto.ScooterID{Id: 3}

	scooter, err := scooterService.GetScooterById(context.Background(), id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(scooter)

	conn, err := grpc.DialContext(context.Background(), ":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_ = proto.NewScooterServiceClient(conn)
}
