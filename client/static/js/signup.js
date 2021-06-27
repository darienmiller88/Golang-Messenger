const signUpForm = document.querySelector("form")
const API_URL = `http://localhost:${location.port}/api/users/signup`
const bordering = "add-border"
const errorMesagePassword = document.querySelector(".error-message-password")
const weakUsernameDiv = document.querySelector(".error-message-username")
const weakPasswordDiv = document.querySelector(".weak-password")
const userNameInputBox = document.getElementById("username-inputbox")
const passwordInputBox = document.getElementById("password-inputbox")
let errorMessageAddedUsername = false, errorMessageAddedPassword = false

signUpForm.addEventListener("submit", async e => {
    e.preventDefault()
    removeUsernameErrors()
    removePasswordErrors()

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
        window.location.href = "/"
    }else if(result["username_taken"]){
        appendUsernameCorrection(result["username_taken"], userNameInputBox)
        errorMessageAddedUsername = true
    }else{
        if (result["weak_username_err"]) {
            appendUsernameCorrection(result["weak_username_err"], userNameInputBox)
            errorMessageAddedUsername = true
        } 
        if(result["password_errors"]){
            appendPasswordCorrections(result["password_errors"])
        }
    }

    //signUpForm.reset()
})

userNameInputBox.addEventListener("keydown", () => {
    if(errorMessageAddedUsername){
        userNameInputBox.classList.remove(bordering)
        errorMessageAddedUsername = false
    }
})

//Function to append the password restrictions for the signup sent from the server in case the users password
//is too weak.
const appendPasswordCorrections = (passwordErrors) =>{
    //Iterate through each password error, and add them to div, coloring them red or green as needed.
    passwordErrors.forEach(elem => {
        let li = document.createElement("li")

        //If the password was too weak, and it checked off one of the conditions, color it red to let the user know
        //their password needs to include the additional checks.
        li.style = (elem["is_password_weak"]) ? "color: red" : "color: green"
        li.append(elem["password_error"])
        weakPasswordDiv.append(li)
    })
}

const appendUsernameCorrection = (message, inputBox) => {
    const h6 = document.createElement("h6")

    h6.append(message)
    weakUsernameDiv.append(h6)
    inputBox.classList.add("add-border")
}

const removePasswordErrors = () => {  
    weakPasswordDiv.querySelectorAll("li").forEach(elem => {
        weakPasswordDiv.removeChild(elem)
    })
}

const removeUsernameErrors = () => {
    weakUsernameDiv.querySelectorAll("h6").forEach(elem => {
        weakUsernameDiv.removeChild(elem)
    })
}