 ## Description

Web Application to extract data from the NASA API and display them on dynamically on an HTML file! Written using the "Gin" web framework. View it at https://nasa-apod-application.herokuapp.com/. Click [here](https://github.com/gin-gonic/gin) to read more on the web framework I used for this project. So far, it allows the user to type in a specific date, and the app will re route the user to a new page with the NASA API picture and information for that date given, or an error page if invalid information is given. Here's a demo!

 ## Installation

```
   cd api
   go build 
   go mod vender //if you desire to have a folder including the dependencies, otherwise ignore
   .\fresh //to restart server on change or
   go run main.go //to run the server normally
```

  ## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Feel free to leave suggestions as well, I'm always looking for ways to improve!

  ## License
[MIT](https://choosealicense.com/licenses/mit/)
