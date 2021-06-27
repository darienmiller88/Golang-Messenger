package main

import (
	//	"encoding/json"
	//"chat_app/fileserver"
	"chat_app/api/controllers"
	"chat_app/api/database"
	"chat_app/api/models"
	"encoding/json"
	"io/ioutil"

	//"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	// "github.com/googollee/go-socket.io"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	//"gopkg.in/unrolled/render.v1"
)

type m map[string]interface{}

func main() {
	dir, _ := os.Getwd()
	router := chi.NewRouter()

	godotenv.Load(filepath.Join(dir, "../.env"))

	router.Use(middleware.Logger)
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

	//getPublicMessages(db)
	fmt.Println("running on port", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func getPublicMessages(d database.DB) {
	var users []models.User

	d.DB.Find(&users)

	data, _ := json.MarshalIndent(users, "", " ")

	ioutil.WriteFile("users.json", data, 0644)
}
