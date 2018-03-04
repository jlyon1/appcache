package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jlyon1/appcache/api"
	"github.com/jlyon1/appcache/database"
	"net/http"
)

func connectDB(db *database.Redis) {
  while(!db.Connect()){
    fmt.Printf("Trying to connect"))
  }
  fmt.Printf("Connected")

}

func main() {
	cachestore := &database.Redis{}
	cachestore.IP = "localhost"
	cachestore.Port = "6379"
	cachestore.DB = 0
	cachestore.Password = ""
	connectDB(cachestore)

	r := mux.NewRouter()
	api := api.API{
		DB: cachestore,
	}
	r.HandleFunc("/ask", api.Ask).Methods("POST")
	r.HandleFunc("/", api.Main).Methods("GET")
	http.ListenAndServe("0.0.0.0:8080", r)
}
