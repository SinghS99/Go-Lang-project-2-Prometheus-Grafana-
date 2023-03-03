package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var REQUEST_INPROGRESS = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "go_app_requests_inprogress",
	Help: "No.of app request in progress",
})

func jain() {
	//Starting app point
	rRtartMyApp()
}

func rRtartMyApp() {

	router := mux.NewRouter()
	router.HandleFunc("/birthday/{name}", func(rw http.ResponseWriter, r *http.Request) {
		REQUEST_INPROGRESS.Inc()
		vars := mux.Vars(r)
		name := vars["name"]
		greetings := fmt.Sprintf("Happy BirthDay %s :", name)
		time.Sleep(5 * time.Second)
		rw.Write([]byte(greetings))
		REQUEST_INPROGRESS.Dec()

	}).Methods("GET")

	log.Println("Staring the application server")
	router.Path("/metrics").Handler(promhttp.Handler())
	http.ListenAndServe(":8080", router)
}
