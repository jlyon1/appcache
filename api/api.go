package api

import (
  "encoding/json"
  "net/http"
  // "fmt"
  "github.com/jlyon1/appcache/database"
)

type API struct{
  DB *database.Redis
}

type CacheRequest struct{
  Address string `json: Address`
  TTL     int    `json: TTL`
}

func (api *API) Ask(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "text/html")

  var cache CacheRequest
  dec := json.NewDecoder(r.Body)
  err:=dec.Decode(&cache)

  if(err != nil){
    http.Error(w,"[appcache] Failed to resolve requested source",500)
    return
  }

  val := api.DB.Find("cr" + cache.Address)
  if val  != ""{
    w.Write([]byte(val))
  }else{
    
  }

}

func (api *API)Main(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("coming soon"));
}
