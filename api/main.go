package main

import (
	//	"encoding/json"
	//"chat_app/fileserver"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	

	// "github.com/googollee/go-socket.io"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/nkovacs/go-socket.io"
	"gopkg.in/unrolled/render.v1"
	//"github.com/go-chi/render"
)

type m map[string]interface{}

func main(){
	godotenv.Load()
	router := chi.NewRouter()
	socketServer, err := socketio.NewServer(nil)

	if err != nil{
		log.Fatal(err)
	}
	
	//Get absolute path to the folder, and append the client folder to it
	dir, _ := os.Getwd()
	root := http.Dir(filepath.Join(dir, "../client"))

	//Render package allows ease of rendering HTML files
	r := render.New(render.Options{
		Extensions: []string{".tmpl", ".html"},
		Directory: "../client/templates",
	})

	//middlewares stack
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	router.Get("/socket.io/",  socketServer.ServeHTTP)
	router.Post("/socket.io/", socketServer.ServeHTTP)
	router.Get("/", func(res http.ResponseWriter, req *http.Request) {	
		r.HTML(res, http.StatusOK, "test", nil)
	})

	

	//Later on in the program, I will wrap the handler function in a authentication middleware for the login
	//routes 


	socketServer.On("connection", func (so socketio.Socket)  {
		fmt.Println("client connected:", so.Id())
		so.Join("chat")
		fmt.Println(socketServer.Count())
		so.On("message", func(message string){
			fmt.Println("client", so.Id(), ":", message)
			so.BroadcastTo("chat", "new_message", so.Id() + ":" + message)
		})

		so.On("disconnect", func(reason string) {
			fmt.Println("reason:", reason)
		})
	})

	//Serve static files
	router.Handle("/*", http.FileServer(root))
	fmt.Println("running on port", os.Getenv("PORT"), "and dir:", root)
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}

func fileServer(router *chi.Mux, root string) {
	fs := http.FileServer(http.Dir(root))

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
}

func renderTemplate(response http.ResponseWriter, fileName string, data interface{}){
	myTemplate := template.Must(template.ParseFiles(fileName))
	if err := myTemplate.Execute(response, data); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}