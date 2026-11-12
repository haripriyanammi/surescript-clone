package main 

import "net/http"


func main(){
   mux:=http.NewServeMux()
   mux.HandleFunc("GET /healthz",healthz)
   http.ListenAndServe(":8080",mux)
}

func healthz(w http.ResponseWriter, r *http.Request ){
  	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))

} 
   