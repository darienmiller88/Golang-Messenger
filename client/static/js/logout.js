const logoutButton = document.querySelector(".log-out")
const SIGNOUT_API_URL = `http://localhost:${location.port}/api/users/signout`

logoutButton.addEventListener("click", async e => {
    const response = await fetch(SIGNOUT_API_URL, {
        method: "POST",
        headers: {
            "Content-type": "application/json"
        }
    })

    const result = await response.json()
    window.location.href = "/signin"
})