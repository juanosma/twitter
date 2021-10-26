package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/juanosma/twitter/middlew"
	"github.com/juanosma/twitter/routers"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el handler y pongo a escuchar*/
func Manejadores() {
	router := mux.NewRouter()
	PORT := os.Getenv("PORT")

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
