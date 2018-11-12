package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"log"
	"math/rand"
	"net/http"
)

func zerorone(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	fmt.Fprint(w, rand.Intn(2))
}

func main() {
	router := httprouter.New()
	router.POST("/create", zerorone)
	//router.GET("/response", zerorone)

	log.Print("Listening...8080");
	handler := cors.Default().Handler(router)
	http.ListenAndServe(":8080", handler)
}
