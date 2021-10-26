package main

import (
	"log"

	"github.com/juanosma/twitter/bd"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	log.Fatal("Con conexion a la BD")
}
