package routing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"scooter_micro/proto"
	"scooter_micro/service"
	"strconv"
)

var (
	scooterIDKey = "scooterId"
	stationIDKey = "stationId"
)
var chosenScooterID, chosenStationID int

type combineForTemplate struct {
	*proto.ScooterList
	*proto.StationList
}

type Routing interface {
	getAllScooters(w http.ResponseWriter, r *http.Request)
	getScooterById(w http.ResponseWriter, r *http.Request)
	startScooterTrip(w http.ResponseWriter, r *http.Request)
	showTripPage(w http.ResponseWriter, r *http.Request)
	ChooseScooter(w http.ResponseWriter, r *http.Request)
	ChooseStation(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	scooterService *service.ScooterService
}

func newHandler(service *service.ScooterService) *handler {
	return &handler{
		scooterService: service,
	}
}

func NewRouter(service *service.ScooterService) *mux.Router {
	router := mux.NewRouter()
	handler := newHandler(service)
	router.HandleFunc(`/scooters`, handler.getAllScooters).Methods("GET")
	router.HandleFunc(`/scooter/{`+scooterIDKey+`}`, handler.getScooterById).Methods("GET")
	router.HandleFunc(`/start-trip/{`+stationIDKey+`}`, handler.showTripPage).Methods("GET")
	router.HandleFunc(`/run`, handler.startScooterTrip).Methods("GET")
	router.HandleFunc(`/choose-station`, handler.ChooseStation).Methods("POST")
	router.HandleFunc(`/choose-scooter`, handler.ChooseScooter).Methods("POST")
	return router
}

func (h *handler) getAllScooters(w http.ResponseWriter, r *http.Request) {
	scooters, err := h.scooterService.GetAllScooters(context.Background(), &proto.Request{})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	json.NewEncoder(w).Encode(scooters)

}

func (h *handler) getScooterById(w http.ResponseWriter, r *http.Request) {
	scooterID, err := strconv.Atoi(mux.Vars(r)[scooterIDKey])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scooter, err := h.scooterService.GetScooterById(context.Background(), &proto.ScooterID{Id: uint64(scooterID)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	json.NewEncoder(w).Encode(scooter)
}

func (h *handler) startScooterTrip(w http.ResponseWriter, r *http.Request) {
	//userFromRequest := GetUserFromContext(r)

	statusStart, err := h.scooterService.CreateScooterStatusInRent(context.Background(),
		&proto.ScooterID{Id: uint64(chosenScooterID)})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(statusStart)

	err = h.scooterService.InitAndRun(context.Background(), &proto.ScooterID{Id: uint64(chosenScooterID)},
		&proto.StationID{Id: uint64(chosenStationID)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//statusEnd, err := scooterService.CreateScooterStatusInRent(chosenScooterID)

	//distance := statusEnd.Location.Distance(statusStart.Location)

	//_, err = orderService.CreateOrder(*userFromRequest, chosenScooterID, statusStart.ID, statusEnd.ID, distance)
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func (h *handler) showTripPage(w http.ResponseWriter, r *http.Request) {
	stationID, err := strconv.Atoi(mux.Vars(r)[stationIDKey])

	scooterList, err := h.scooterService.GetAllScootersByStationID(context.Background(),
		&proto.StationID{Id: uint64(stationID)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stationList, err := h.scooterService.GetAllStations(context.Background(), &proto.Request{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("./templates/scooter-run.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmpl.Execute(w, combineForTemplate{scooterList, stationList})
	if err != nil {
		fmt.Println(err)
	}
}

func (h *handler) ChooseScooter(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	chosenScooterID, err = strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	fmt.Println(chosenScooterID)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) ChooseStation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	chosenStationID, err = strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	fmt.Println(chosenStationID)
	w.WriteHeader(http.StatusOK)
}
