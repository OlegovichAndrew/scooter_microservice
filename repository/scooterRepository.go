package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"scooter_micro/proto"
)

//ScooterRepository the interface which implemented by functions which connect to the database.
type ScooterRepository interface {
	GetAllScooters(ctx context.Context, request *proto.Request) (*proto.ScooterList, error)
	GetAllScootersByStationID(context context.Context, id *proto.StationID) (*proto.ScooterList, error)
	GetScooterById(context context.Context, id *proto.ScooterID) (*proto.Scooter, error)
	GetScooterStatus(context context.Context, id *proto.ScooterID) (*proto.ScooterStatus, error)
	SendCurrentStatus(context context.Context, status *proto.SendStatus) (*proto.Response, error)
	CreateScooterStatusInRent(context context.Context, id *proto.ScooterID) (*proto.ScooterStatusInRent, error)
	GetStationById(ctx context.Context,id *proto.StationID) (*proto.Station, error)
}

type ScooterRepo struct {
	db *sql.DB
}

func NewScooterRepo(db *sql.DB) *ScooterRepo {
	return &ScooterRepo{db: db}
}

func (scr *ScooterRepo) GetAllScooters(ctx context.Context, request *proto.Request) (*proto.ScooterList, error) {
	scooterList := &proto.ScooterList{}

	querySQL := `SELECT s.id, sm.max_weight, sm.model_name, ss.battery_remain, ss.can_be_rent
					FROM scooters as s 
					JOIN scooter_models as sm 
					ON s.model_id=sm.id 
					JOIN scooter_statuses as ss 
					ON s.id=ss.scooter_id 
					ORDER BY s.id`

	rows, err := scr.db.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		var scooter *proto.Scooter
		err := rows.Scan(&scooter.Id, &scooter.MaxWeight, &scooter.ScooterModel, &scooter.BatteryRemain,
			&scooter.CanBeRent)
		if err != nil {
			return nil, err
		}
		scooterList.Scooters = append(scooterList.Scooters, scooter)
	}
	return scooterList, nil
}

func (scr *ScooterRepo) GetAllScootersByStationID(ctx context.Context, id *proto.StationID) (*proto.ScooterList, error) {
	scooterList := &proto.ScooterList{}

	querySQL := `SELECT s.id, sm.max_weight, sm.model_name, ss.battery_remain, ss.can_be_rent
					FROM scooters as s 
					JOIN scooter_models as sm 
					ON s.model_id=sm.id 
					JOIN scooter_statuses as ss 
					ON s.id=ss.scooter_id 
					WHERE ss.station_id=$1
					ORDER BY s.id`

	rows, err := scr.db.QueryContext(ctx, querySQL, id.Id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		var scooter *proto.Scooter
		err := rows.Scan(&scooter.Id, &scooter.MaxWeight, &scooter.ScooterModel, &scooter.BatteryRemain,
			&scooter.CanBeRent)
		if err != nil {
			return nil, err
		}
		scooterList.Scooters = append(scooterList.Scooters, scooter)
	}
	return scooterList, nil
}

//GetScooterById returns exact scooter by its ID.
func (scr *ScooterRepo) GetScooterById(ctx context.Context, id *proto.ScooterID) (*proto.Scooter, error) {
	scooter := &proto.Scooter{}
	querySQL := `SELECT s.id, sm.max_weight, sm.model_name, ss.battery_remain, ss.can_be_rent
					FROM scooters as s 
					JOIN scooter_models as sm 
					ON s.model_id=sm.id 
					JOIN scooter_statuses as ss 
					ON s.id=ss.scooter_id 
					WHERE s.id=$1`

	row := scr.db.QueryRowContext(ctx, querySQL, id.Id)
	err := row.Scan(&scooter.Id, &scooter.MaxWeight, &scooter.ScooterModel, &scooter.BatteryRemain, &scooter.CanBeRent)
	if err != nil {
		return nil, err
	}

	return scooter, nil
}

//GetScooterStatus returns the ScooterStatus model of the chosen scooter by its ID.
func (scr *ScooterRepo) GetScooterStatus(ctx context.Context, id *proto.ScooterID) (*proto.ScooterStatus, error) {
	var scooterStatus = &proto.ScooterStatus{}
	scooter, err := scr.GetScooterById(ctx, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	scooterStatus.Scooter = scooter

	querySQL := `SELECT battery_remain, latitude, longitude 
					FROM scooter_statuses
					WHERE scooter_id=$1`

	row := scr.db.QueryRowContext(ctx, querySQL, int(id.Id))
	err = row.Scan(&scooterStatus.BatteryRemain,
		&scooterStatus.Location.Latitude, &scooterStatus.Location.Longitude)
	if err != nil {
		return nil, err
	}

	return scooterStatus, nil
}

//CreateScooterStatusInRent creates a new record in ScooterStatusesInRent by scooter's ID and returns the
//ScooterStatusInRent model.
func (scr *ScooterRepo) CreateScooterStatusInRent(ctx context.Context,
	id *proto.ScooterID) (*proto.ScooterStatusInRent, error) {
	var scooterStatusInRent = proto.ScooterStatusInRent{}
	scooterStatus, err := scr.GetScooterStatus(ctx, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	scooterStatusInRent.Location = scooterStatus.Location

	querySQL := `INSERT INTO scooter_statuses_in_rent(date_time, latitude, longitude) 
					VALUES(now(), $1, $2) RETURNING id, date_time`

	err = scr.db.QueryRowContext(ctx, querySQL, scooterStatus.Location.Latitude,
		scooterStatus.Location.Longitude).Scan(&scooterStatusInRent.ScooterID, &scooterStatusInRent.DateTime)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &scooterStatusInRent, nil
}

//SendCurrentStatus updates ScooterStatus with given parameters.
func (scr *ScooterRepo) SendCurrentStatus(ctx context.Context, status *proto.SendStatus) (*proto.Response, error) {
	var canBeRent bool
	if status.BatteryRemain > 10 {
		canBeRent = true
	}

	querySQL := `UPDATE scooter_statuses 
					SET latitude=$1, longitude=$2, battery_remain=$3, can_be_rent=$4, station_id=$5
					WHERE scooter_id=$6`

	rows, err := scr.db.QueryContext(ctx, querySQL, status.Latitude, status.Longitude,
		status.BatteryRemain,
		canBeRent,
		status.StationID, status.ScooterID)
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	return &proto.Response{}, err
}

func (scr *ScooterRepo) GetStationById(ctx context.Context,id *proto.StationID) (*proto.Station, error) {
	station := &proto.Station{}

	querySQL := `SELECT * FROM scooter_stations WHERE id = $1;`
	row := scr.db.QueryRowContext(ctx, querySQL, id.Id)
	err := row.Scan(&station.Id, &station.Name, &station.IsActive, &station.Location.Latitude, &station.Location.Longitude)

	return station, err
}
