package main

import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var hm HandlerManager = HandlerManager{}

func main() {

	//It is for heroku
	port := os.Getenv("PORT")
	if port == "" {
		getLogger().Info("Unable to fetch port from environmental variable PORT now will run default on 8080")
		port = ":8080"
	} else {
		port = ":" + port
	}

	rtr := mux.NewRouter()
	srv := http.Server{Addr: port}

	rtr.HandleFunc("/hash", hm.EncodePasswordHandler)
	rtr.HandleFunc("/hash/{id:[0-9]+}", hm.GetHashHandler)
	rtr.HandleFunc("/stats", hm.StatsHandler)
	rtr.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		getLogger().Info("Shutting Down Server")
		srv.Shutdown(context.Background())
	})
	http.Handle("/", rtr)

	err := srv.ListenAndServe()

	if err != nil {
		getLogger().Error("Failed to start Server:" + err.Error())
	} else {
		getLogger().Info("Server started at port 8080")
	}

}
