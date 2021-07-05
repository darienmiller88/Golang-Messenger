const signInForm             = document.querySelector("form")
const closebuttonClassName   = "bi bi-x"
const API_URL                = `http://localhost:${location.port}/api/users/signin`
const SESSION_EXPIRED_URL    = `http://localhost:${location.port}/api/users/session-expired`
const errorMessageTarget     = document.querySelector(".error-message-target")
const sessionExpiredTarget   = document.querySelector(".session-expired-target")
const createNewAccountButton = document.getElementById("create-account-button")

// checkSessionExpired()

// if(document.cookie === ""){
//     console.log("cookie not set!");
// }else{
//     console.log("cookie set!" + document.cookie);
// }

signInForm.addEventListener("submit", async e => {
    e.preventDefault()
    const formData =new FormData(signInForm)
    const username = formData.get("userName")
    const password = formData.get("password")
    const userData = {
        username, 
        password,
    }

    const response = await fetch(API_URL, {
        method: "POST",
        body: JSON.stringify(userData),
        headers: {
            "Content-type": "application/json"
        }
    })

    const result = await response.json()
    
    if(result["error_message"]){
        errorMessageTarget.innerHTML = ''
        createErrorElement(result["error_message"])
        return
    }
    
    signInForm.reset()
    window.location.href = "/"
})

errorMessageTarget.addEventListener("click", e => {
    if(e.target.className === closebuttonClassName){
        e.target.parentElement.remove()
    }
})

createNewAccountButton.addEventListener("click", e =>{
    window.location.href = "/signup"
})

async function checkSessionExpired(){
    const response = await fetch(SESSION_EXPIRED_URL, {
        method: 'POST',
        headers: {
            "Content-type": "application/json"
        },
    })

    const result = await response.json()
    
    console.log(result)
    if(result["is_session_Expired"]){
        createSessionExpiredElement(result["session_expired_message"])
    }
}

function createSessionExpiredElement(message){
    const sessionExpiredDiv = document.createElement("div")

    sessionExpiredDiv.className = "session-expired"
    sessionExpiredDiv.append(message)
    sessionExpiredTarget.append(sessionExpiredDiv)
}

function createErrorElement(message) {
    const errorBodyDiv = document.createElement("div")
    const errorMessage = document.createElement("div")
    const buttonIcon   = document.createElement("i")

    errorMessage.className = "error-message"
    errorBodyDiv.className = "error-message-body"
    buttonIcon.className   = closebuttonClassName

    errorMessage.append(message)
    errorBodyDiv.append(errorMessage)
    errorBodyDiv.append(buttonIcon)
    errorMessageTarget.append(errorBodyDiv)
}