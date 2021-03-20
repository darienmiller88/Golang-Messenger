 # Chat 
 
  ![](https://img.shields.io/badge/made%20by-DarienMiller-blue)
  ![](https://img.shields.io/badge/Golang-1.14-yellow)
  ![](https://img.shields.io/badge/HTML%2B%20CSS-48%25-red)
  ![](https://img.shields.io/badge/test-passing-green)
 <img width="960" alt="chat app" src="https://user-images.githubusercontent.com/32966645/111588302-bb79cd80-8799-11eb-85a0-550fd92a1a8a.PNG">

 ## Description

Full stack chat application built using Golang and Postgres in the backend, and Javascript, HTML and CSS in the frontend. Enjoy!
Features to be added:

- User log in and log out using Golang sessions
- Public and private messaging using socket.io on the frontend and backend
- Storage of user messages in message history table

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
