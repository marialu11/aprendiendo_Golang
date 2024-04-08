package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marialu11/practica-golang/db"
	"github.com/marialu11/practica-golang/routes"
)

func main() {

	config := db.LoadDBConfig()
	db.DBConnection(config)
	//db.DB.AutoMigrate(models.Task{})
	//db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
