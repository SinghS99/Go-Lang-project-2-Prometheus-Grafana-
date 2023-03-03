package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var REQUEST_COUNT = promauto.NewCounter(prometheus.CounterOpts{
	Name: "go_app_requests_count",
	Help: "Total_App Http Request Count.",
})

func Main() {
	//Starting app point
	startMyApp()
}

func StartMyApp() {

	router := mux.NewRouter()
	router.HandleFunc("/birthday/{name}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		greetings := fmt.Sprintf("Happy BirthDay %s :", name)
		rw.Write([]byte(greetings))
		REQUEST_COUNT.Inc()

	}).Methods("GET")

	log.Println("Staring the application server")
	router.Path("/metrics").Handler(promhttp.Handler())
	http.ListenAndServe(":8080", router)
}
