package servestatic

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"gopkg.in/unrolled/render.v1"
)

type ServeStatic struct{
	fs http.Handler
	root http.Dir
	render *render.Render
}

//This static file handler will use the absolute path to the directory it is in 
//(c:\Users\Computername\Desktop\etc...). This parameter will then take the path from this directory to the 
//directry of the static folder containing css, img and js files.
func (s *ServeStatic) Init(relativePathToStaticFolder string, router *chi.Mux){
	workDir, _ := os.Getwd()
	s.root = http.Dir(filepath.Join(workDir, relativePathToStaticFolder))
	s.fs = http.FileServer(http.Dir(s.root))
	s.render = render.New(render.Options{
		Extensions: []string{".tmpl", ".html"},
		Directory:  "./client/templates",
	})

	//router.Handle("/static/*", http.StripPrefix("/static", s.fs))
}

//This function will be used as a "HandlerFunc" in routers. Please route to the "/*" path using a get request
func (s *ServeStatic) StaticFileHandler(res http.ResponseWriter, req *http.Request){
	if _, err := os.Stat(fmt.Sprintf("%s", s.root) + req.RequestURI); os.IsNotExist(err) {
		s.render.HTML(res, 404, "404Error", nil)
	} else{
		s.fs.ServeHTTP(res, req)
	}
}