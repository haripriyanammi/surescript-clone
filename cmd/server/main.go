package main 

import "net/http"
import "log"


func main(){
   mux:=http.NewServeMux()
   mux.HandleFunc("GET /healthz",healthz)
   err:=http.ListenAndServe(":8080",mux)
   log.Fatal(err)
}

func healthz(w http.ResponseWriter, r *http.Request ){
  	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))

} 
   