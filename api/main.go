package main

import (
	//	"encoding/json"
	//"chat_app/fileserver"
	"chat_app/api/controllers"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	// "github.com/googollee/go-socket.io"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	viewsController := controllers.ViewsController{}
	socketController := controllers.SocketController{}
	
	viewsController.Init()
	socketController.Init()
	
	router.Mount("/", viewsController.Router)
	router.Mount("/socket.io/", socketController.Router)
	//Later on in the program, I will wrap the handler function in a authentication middleware for the login
	//routes 

	fmt.Println("running on port", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}