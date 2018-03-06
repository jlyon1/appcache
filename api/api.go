package api

import (
	"encoding/json"
	"fmt"
	"github.com/jlyon1/appcache/database"
	"net/http"
	"time"
)

type API struct {
	DB *database.Redis
}

type CacheRequest struct {
	Address string `json: Address`
	TTL     int    `json: TTL`
}

func log(data string) {
	fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05-0700") + " " + data)
}

func (api *API) Ask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/html")

	var cache CacheRequest
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&cache)

	if err != nil {
		log("Requested " + cache.Address)
		log("Failed to evaluate address string")
		http.Error(w, "[appcache] Failed to evaluate requested source", 500)
		return
	}
	log("Requested " + cache.Address)

	val := api.DB.Find("cr" + cache.Address)

	if val != "" {
		log("Found in cache " + cache.Address)
		w.Write([]byte(val))
		return
	} else {
		log("Not Found in cache " + cache.Address)
		http.Error(w, "Not found in cache", 500)
	}

}

func (api *API) Main(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Setting this up later"))
}
