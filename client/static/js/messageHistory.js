const API_URL  = `http://localhost:${location.port}/api/messages/get-message-history`
const username = document.querySelector(".username").textContent

getMessages()

async function getMessages() {
    const user = {
        user_name: username,
        chat_name: "public"
    }

    const response = await fetch(API_URL, {
        method: "POST",
        headers: {
            "Conten-type": "application/json",
        },
        body: JSON.stringify(user)
    })
    const result = await response.json()

    result.forEach(message => {
        console.log(message);
        appendMessage(message, "your-message", false)
    })
}