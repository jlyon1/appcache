package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jlyon1/appcache/api"
	"github.com/jlyon1/appcache/database"
	"net/http"
	"io/ioutil"
	"time"
	"os"
)

type Config struct {
	Sites []string `json:strings`
}

var cfg Config

var cachestore database.Redis

func connectDB(db *database.Redis) {
	for db.Connect() == false {
		fmt.Printf("Trying to connect\n")
	}
	fmt.Printf("Connected\n")

}

func loadConfig() (Config, bool) {
	file, err := os.Open("config.json")
	if err != nil {
		return Config{}, false
	}
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		return Config{}, false
	}
	return configuration, true
}

func cacheItems() {
	for {
		fmt.Println("Refreshing Cache")
		for _, item := range cfg.Sites {
			fmt.Println("Refreshing " + item)
			resp, _ := http.Get(item)
			bdyString := ""
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				bdy, _ := ioutil.ReadAll(resp.Body)
				bdyString = string(bdy)
			}
			cachestore.SetString("cr"+item, bdyString)
			cachestore.Expire("cr"+item, time.Duration(3600))
		}
		<-time.After(time.Second*3600)

	}
}

func main() {
	cachestore.IP = "redis"
	cachestore.Port = "6379"
	cachestore.DB = 0
	cachestore.Password = ""
	connectDB(&cachestore)

	var success bool

	cfg, success = loadConfig()
	if !success {
		panic("Failed to load configuration")
	}
	go cacheItems()

	r := mux.NewRouter()
	api := api.API{
		DB: &cachestore,
	}
	r.HandleFunc("/ask", api.Ask).Methods("POST")
	r.HandleFunc("/", api.Main).Methods("GET")
	http.ListenAndServe("0.0.0.0:8080", r)
}
