const signUpForm = document.querySelector("form")
const API_URL = "http://localhost:7000/api/users/signup"
const signInRoute = "/signin"
const bordering = "add-border"
const errorMessageUsername = document.querySelector(".error-message-username")
const errorMesagePassword = document.querySelector(".error-message-password")
const userNameInputBox = document.getElementById("username-inputbox")
const passwordInputBox = document.getElementById("password-inputbox")
let errorMessageAddedUsername = false, errorMessageAddedPassword = false

signUpForm.addEventListener("submit", async e => {
    e.preventDefault()
    const formData = new FormData(signUpForm)
    const username = formData.get("userName")
    const password = formData.get("password")
    const userData = {
        username, 
        password,
    }

    const reponse = await fetch(API_URL, {
        method: "POST",
        body: JSON.stringify(userData),
        headers: {
            "Content-type": "application/json",
        }
    })
    const result = await reponse.json()

    if (result["success"]){
        window.location.href = signInRoute
    }else if(result["username_taken"]){
        createElement(result["username_taken"], errorMessageUsername, userNameInputBox)
        errorMessageAddedUsername = true
    }else{
        if (result["no_username_err"] && errorMessageAddedUsername) {
            createElement(result["no_username_err"], errorMessageUsername, userNameInputBox)
            errorMessageAddedUsername = true
        } 
        if (result["no_password_err"]) {
            createElement(result["no_password_err"], errorMesagePassword, passwordInputBox)  
            errorMessageAddedPassword = true
        }
    }

    console.log(result)
    signUpForm.reset()
})

userNameInputBox.addEventListener("keydown", () => {
    if(errorMessageAddedUsername){
        userNameInputBox.classList.remove(bordering)
        errorMessageUsername.removeChild(errorMessageUsername.firstChild)
        errorMessageAddedUsername = false
    }
})

passwordInputBox.addEventListener("keydown", e => {
    if(errorMessageAddedPassword){
        passwordInputBox.classList.remove(bordering)
        errorMesagePassword.removeChild(errorMesagePassword.firstChild)
        errorMessageAddedPassword = false
    }
})

const createElement = (message, divToAppendTo, inputBox) => {
    const h6 = document.createElement("h6")

    h6.append(message)
    divToAppendTo.append(h6)
    inputBox.classList.add("add-border")
}