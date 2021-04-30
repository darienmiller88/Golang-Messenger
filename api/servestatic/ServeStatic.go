package servestatic

import (
	"net/http"
	"os"
	"path/filepath"
)

type ServeStatic struct{
	fs http.Handler
}

//This static file handler will use the absolute path to the directory it is in 
//(c:\Users\Computername\Desktop\etc...). This parameter will then take the path from this directory to the 
//directry of the static folder containing css, img and js files.
func (s *ServeStatic) Init(relativePathToStaticFolder string){
	workDir, _ := os.Getwd()
	root := http.Dir(filepath.Join(workDir, relativePathToStaticFolder))
	s.fs = http.FileServer(http.Dir(root))
}

//This function will be used as a "HandlerFunc" in routers. Please route to the "/*" path using a get request
func (s *ServeStatic) StaticFileHandler(res http.ResponseWriter, req *http.Request){
	s.fs.ServeHTTP(res, req)
}