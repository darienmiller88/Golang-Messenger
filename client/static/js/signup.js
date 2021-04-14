const signUpForm = document.querySelector("form")
const API_URL = "http://localhost:3333/"

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

    console.log()
    signUpForm.reset()
})