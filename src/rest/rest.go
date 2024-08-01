package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Rest struct {
	Port string
}

func (r *Rest) Init(db *gorm.DB) {
	router := mux.NewRouter()
	log.Println("Server is running on port " + r.Port)
	http.ListenAndServe(r.Port, router)
}
