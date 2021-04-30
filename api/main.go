package main

import (
	//	"encoding/json"
	//"chat_app/fileserver"
	"chat_app/api/controllers"
	"chat_app/api/database"
	// "chat_app/api/models"
	// "chat_app/api/session"
	// "encoding/gob"
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

func main(){
	dir, _ := os.Getwd()
	router := chi.NewRouter()

	godotenv.Load(filepath.Join(dir, "../.env"))

	//middlewares stack
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(cors.AllowAll().Handler)

	viewsController := controllers.ViewsController{}
	messageController := controllers.MessageController{}
	socketController := controllers.SocketController{}
	userController := controllers.UserController{}
	db := database.DB{}
		
	db.InitDB(false)
	socketController.Init()	
	viewsController.Init()
	userController.Init(db.DB)
	messageController.Init(db.DB)

	router.Mount("/", viewsController.Router)
	router.Mount("/socket.io/", socketController.Router)
	router.Mount("/api/users", userController.Router)

	fmt.Println("running on port", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}

