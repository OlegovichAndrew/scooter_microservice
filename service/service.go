package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"scooter_micro/proto"
	"scooter_micro/repository"
	"time"
)

const (
	step          = 0.0001
	dischargeStep = 0.1
	interval      = 450
)

//ScooterService is a service which responsible for gRPC scooter.
type ScooterService struct {
	Repo *repository.ScooterRepo
	*proto.UnimplementedScooterServiceServer
}

func (gss *ScooterService) Receive(server proto.ScooterService_ReceiveServer) error {
	return nil
}

func (gss *ScooterService) Register(request *proto.ClientRequest, server proto.ScooterService_RegisterServer) error {
	return nil
}

//ScooterClient is a struct with parameters which will be translated by the gRPC connection.
type ScooterClient struct {
	ID            uint64
	coordinate    proto.Location
	batteryRemain float64
	stream        proto.ScooterService_ReceiveClient
}

//NewScooterService creates a new GrpcScooterService.
func NewScooterService(repoScooter *repository.ScooterRepo) *ScooterService {
	return &ScooterService{
		Repo: repoScooter,
	}
}

//NewGrpcScooterClient creates a new GrpcScooterClient with given parameters.
func NewGrpcScooterClient(id uint64, coordinate *proto.Location, battery float64,
	stream proto.ScooterService_ReceiveClient) *ScooterClient {
	return &ScooterClient{
		ID:            id,
		coordinate:    *coordinate,
		batteryRemain: battery,
		stream:        stream,
	}
}

//InitAndRun the main function of scooter's trip. It analyzes the scooter parameters from database by its ID.
//If they satisfy the conditions, function creates connection to the gRPC server, creates gRPC client,
//calls 'run' function which moves the scooter to the destination point.
//After finished moves it sends the current scooter status to the database.
func (gss *ScooterService) InitAndRun(ctx context.Context, id *proto.ScooterID, chosenStationID int) error {
	stationID := &proto.StationID{Id: uint64(chosenStationID)}
	scooter, err := gss.GetScooterById(ctx, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	scooterStatus, err := gss.GetScooterStatus(ctx, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if scooter.CanBeRent {
		var coordinate *proto.Location
		station, err := gss.GetStationByID(ctx, stationID)
		if err != nil {
			return err
		}
		coordinate.Latitude = station.Location.Latitude
		coordinate.Longitude = station.Location.Longitude

		conn, err := grpc.DialContext(ctx, ":8000", grpc.WithInsecure())

		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		sClient := proto.NewScooterServiceClient(conn)
		stream, err := sClient.Receive(ctx)
		if err != nil {
			log.Fatal(err)
		}

		client := NewGrpcScooterClient(uint64(id.Id),
			scooterStatus.Location, scooter.BatteryRemain, stream)
		err = client.run(coordinate)
		if err != nil {
			fmt.Println(err)
		}

		sendStatus := &proto.SendStatus{
			ScooterID: client.ID, StationID: uint64(chosenStationID),
			Latitude: client.coordinate.Latitude, Longitude: client.coordinate.Longitude, BatteryRemain: client.batteryRemain}

		_, err = gss.SendCurrentStatus(ctx, sendStatus)
		if err != nil {
			fmt.Println(err)
		}

		if client.batteryRemain <= 0 {
			err = fmt.Errorf("scooter battery discharged. Trip is over")
			return err
		}
		return nil
	}

	err = fmt.Errorf("you can't use this scooter. Choose another one")
	fmt.Println(err.Error())
	return err
}

//grpcScooterMessage sends the message be gRPC stream in a format which defined in the *proto file.
func (s *ScooterClient) grpcScooterMessage() {
	intPol := time.Duration(interval) * time.Millisecond

	fmt.Println("executing run in client")
	msg := &proto.ClientMessage{
		Id:        s.ID,
		Latitude:  s.coordinate.Latitude,
		Longitude: s.coordinate.Longitude,
	}
	err := s.stream.Send(msg)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(intPol)
}

//run is responsible for scooter's movements from his current position to the destination point.
//Run also is responsible for scooter's discharge. Every step battery charge decrease by the constant discharge value.
func (s *ScooterClient) run(station *proto.Location) error {

	switch {
	case s.coordinate.Latitude <= station.Latitude && s.coordinate.Longitude <= station.Longitude:
		for ; s.coordinate.Latitude <= station.Latitude && s.coordinate.Longitude <= station.Longitude && s.
			batteryRemain > 0; s.
			coordinate.Latitude,
			s.coordinate.Longitude, s.batteryRemain = s.coordinate.Latitude+step, s.coordinate.Longitude+step,
			s.batteryRemain-dischargeStep {
			s.grpcScooterMessage()
		}
		fallthrough
	case s.coordinate.Latitude >= station.Latitude && s.coordinate.Longitude <= station.Longitude:
		for ; s.coordinate.Latitude >= station.Latitude && s.coordinate.Longitude <= station.Longitude && s.
			batteryRemain > 0; s.coordinate.
			Latitude,
			s.coordinate.Longitude, s.batteryRemain = s.coordinate.Latitude-step, s.coordinate.Longitude+step,
			s.batteryRemain-dischargeStep {
			s.grpcScooterMessage()
		}
		fallthrough
	case s.coordinate.Latitude >= station.Latitude && s.coordinate.Longitude >= station.Longitude:
		for ; s.coordinate.Latitude >= station.Latitude && s.coordinate.Longitude >= station.Longitude && s.
			batteryRemain > 0; s.coordinate.
			Latitude,
			s.coordinate.Longitude, s.batteryRemain = s.coordinate.Latitude-step, s.coordinate.Longitude-step,
			s.batteryRemain-dischargeStep {
			s.grpcScooterMessage()
		}
		fallthrough
	case s.coordinate.Latitude <= station.Latitude && s.coordinate.Longitude >= station.Longitude:
		for ; s.coordinate.Latitude <= station.Latitude && s.coordinate.Longitude >= station.Longitude && s.
			batteryRemain > 0; s.coordinate.
			Latitude,
			s.coordinate.Longitude, s.batteryRemain = s.coordinate.Latitude+step, s.coordinate.Longitude-step,
			s.batteryRemain-dischargeStep {
			s.grpcScooterMessage()
		}
		fallthrough
	case s.coordinate.Latitude <= station.Latitude:
		for ; s.coordinate.Latitude <= station.Latitude && s.
			batteryRemain > 0; s.coordinate.Latitude, s.batteryRemain = s.coordinate.Latitude+step, s.batteryRemain-dischargeStep {
			s.grpcScooterMessage()
		}
		fallthrough
	case s.coordinate.Latitude >= station.Latitude:
		for ; s.coordinate.Latitude >= station.Latitude && s.
			batteryRemain > 0; s.coordinate.Latitude, s.batteryRemain = s.coordinate.Latitude-step, s.batteryRemain-dischargeStep {
			s.grpcScooterMessage()
		}
		fallthrough
	case s.coordinate.Longitude >= station.Longitude:
		for ; s.coordinate.Longitude >= station.Longitude && s.
			batteryRemain > 0; s.coordinate.Longitude, s.batteryRemain = s.coordinate.Longitude-step,
			s.batteryRemain-dischargeStep {
			s.grpcScooterMessage()
		}
		fallthrough
	case s.coordinate.Longitude <= station.Longitude:
		for ; s.coordinate.Longitude <= station.Longitude && s.
			batteryRemain > 0; s.coordinate.Longitude, s.batteryRemain = s.coordinate.Longitude+step,
			s.batteryRemain-dischargeStep {
			s.grpcScooterMessage()
		}
	default:
		return fmt.Errorf("error happened")
	}
	return nil
}

//GetAllScooters gives the access to the ScooterRepo.GetAllScooters function.
func (gss *ScooterService) GetAllScooters(ctx context.Context, request *proto.Request) (*proto.ScooterList, error) {
	return gss.Repo.GetAllScooters(ctx, request)
}

func (gss *ScooterService) GetAllScootersByStationID(ctx context.Context, id *proto.StationID) (*proto.ScooterList,
	error) {
	return gss.Repo.GetAllScootersByStationID(ctx, id)
}

//GetScooterById gives the access to the ScooterRepo.GetScooterById function.
func (gss *ScooterService) GetScooterById(ctx context.Context, id *proto.ScooterID) (*proto.Scooter, error) {
	return gss.Repo.GetScooterById(ctx, id)
}

//GetScooterStatus gives the access to the ScooterRepo.GetScooterStatus function.
func (gss *ScooterService) GetScooterStatus(ctx context.Context, status *proto.ScooterID) (*proto.ScooterStatus, error) {
	return gss.Repo.GetScooterStatus(ctx, status)
}

//SendCurrentStatus gives the access to the ScooterRepo.SendCurrentStatus function.
func (gss *ScooterService) SendCurrentStatus(ctx context.Context, status *proto.SendStatus) (*proto.Response, error) {
	return gss.Repo.SendCurrentStatus(ctx, status)
}

//CreateScooterStatusInRent gives the access to the ScooterRepo.CreateScooterStatusInRent function.
func (gss *ScooterService) CreateScooterStatusInRent(ctx context.Context,
	id *proto.ScooterID) (*proto.ScooterStatusInRent,
	error) {
	return gss.Repo.CreateScooterStatusInRent(ctx, id)
}
