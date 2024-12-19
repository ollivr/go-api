package main

import (
	"fmt"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"goapi/internal/handlers"
	"net/http"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API")
	fmt.Println(`GO API`)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err)
	}

}
