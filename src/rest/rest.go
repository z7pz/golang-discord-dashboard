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
	router.HandleFunc("/discord/redirect", DiscordRedirect)
	router.HandleFunc("/dashboard/guilds", GetGuilds).Methods("GET")
	// 	router.HandleFunc("/dashboard/:id/prefix", GetGuildPrefix).Methods("GET")
	// 	router.HandleFunc("/dashboard/:id/prefix", ChangeGuildPrefix).Methods("POST")
	log.Println("Server is running on port " + r.Port)
	http.ListenAndServe(r.Port, router)
}
