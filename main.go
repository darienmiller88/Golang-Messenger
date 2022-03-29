package main

import (
	"Golang-Messenger/api/controllers"
	"Golang-Messenger/api/database"

	"fmt"
	"log"
	"net/http"
	"os"
	//"path/filepath"


	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type m map[string]interface{}

func main() {
	router := chi.NewRouter()
	dir, _ := os.Getwd()

	fmt.Println("dir:", dir)
	godotenv.Load()

	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(cors.AllowAll().Handler)

	viewsController := controllers.ViewsController{}
	socketController := controllers.SocketController{}
	apiController := controllers.APIControllers{}
	db := database.DB{}

	db.InitDB(false)
	socketController.Init(db.DB)
	viewsController.Init()
	apiController.Init(db.DB)

	router.Mount("/", viewsController.Router)
	router.Mount("/api", apiController.Router)
	router.Mount("/socket.io/", socketController.Router)

	go socketController.SocketServer.Serve()
	defer socketController.SocketServer.Close()

	fmt.Println("running on port", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}