const signInForm = document.querySelector("form")
const signInRoute = "signin"
const closebuttonClassName = "bi bi-x"
const API_URL = `http://localhost:7000/api/users/${signInRoute}`
const passwordDiv = document.querySelector(".password-div")
const errorMessageBodyDiv = document.querySelector(".error-message-body")
const errorMessageDiv = document.querySelector(".error-message")
const createNewAccountButton = document.getElementById("create-account-button")

errorMessageBodyDiv.style.display = 'none'

signInForm.addEventListener("submit", async e => {
    e.preventDefault()
    const formData =new FormData(signInForm)
    const username = formData.get("userName")
    const password = formData.get("password")
    const userData = {
        username, 
        password,
    }

    console.log(userData)
    const response = await fetch(API_URL, {
        method: "POST",
        body: JSON.stringify(userData),
        headers: {
            "Content-type": "application/json"
        }
    })

    const result = await response.json()
    
    if(result["error_message"]){
        errorMessageBodyDiv.style.display = ''
        errorMessageDiv.innerHTML = result["error_message"]
        return
    }
    
    console.log(result)
    signInForm.reset()
    window.location.href = "/"
})

// errorMessageBodyDiv.addEventListener("click", e => {
//     if(e.target.className === closebuttonClassName){
//         console.log("x button hit!");
//         e.target.parentElement.remove()
//     }
// })

createNewAccountButton.addEventListener("click", e =>{
    window.location.href = "/signup"
})

const createElement = (message) => {
    const errorBodyDiv = document.createElement("div")
    const errorMessage = document.createElement("div")
    const buttonIcon   = document.createElement("i")

    errorMessage.className = "error-message"
    errorBodyDiv.className = "error-message-body"
    buttonIcon.className   = closebuttonClassName

    errorMessage.append(message)
    errorBodyDiv.append(errorMessage)
    errorBodyDiv.append(buttonIcon)
    passwordDiv.append(errorBodyDiv)
}