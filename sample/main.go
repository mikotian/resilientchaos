package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/gorilla/mux"
)

type Ticket struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Priority      string `json:"priority"`
}

func main() {

	go func() {
			http.Handle("/metrics", promhttp.Handler())
        		http.ListenAndServe(":2112", nil)
	}()

	

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tickets", jobsPostHandler).Methods("POST")
	fmt.Printf("Starting Server...")
	log.Fatal(http.ListenAndServe(":9090", router))

}

func jobsPostHandler(w http.ResponseWriter, r *http.Request) {

	//Retrieve body from http request
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	//Save data into Job struct
	var _ticket Ticket
	err = json.Unmarshal(b, &_ticket)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Print("Saving to kafka")

	//saveJobToKafka(_ticket)

	//Convert job struct into json
	jsonString, err := json.Marshal(_ticket)
	fmt.Println(jsonString)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("content-type", "application/json")
	w.Header().Set("content-length", (string)(len("Ticket Accepted")))
	fmt.Println("Wrote Header")
	w.Write([]byte("Ticket Accepted"))
	/*	
	//Set content-type http header
	w.Header().Set("content-type", "application/json")
	fmt.Println("Wrote Header")
	//Send back data as response
	w.Write(jsonString)*/

}