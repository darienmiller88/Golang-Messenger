const GET_MESSAGES_API_URL   = `http://localhost:${location.port}/api/messages/public-messages`
const REMOVE_MESSAGE_API_URL = `http://localhost:${location.port}/api/messages/delete-message`
const CHECK_EXPIRED_API_URL  = `http://localhost:${location.port}/api/users/session-expired`
const SOCKET_API_URL         = `http://localhost:${location.port}`
const form                   = document.querySelector("form")
const inputBox               = document.getElementById("message-input")
const removeMessageButton    = document.getElementById("remove-comment-button")
const username               = document.querySelector(".username").textContent
const ellipseIconName        = "bi bi-three-dots-vertical"
const socket                 = io.connect(SOCKET_API_URL)
let messageToBeRemoved

getPublicMessages()

socket.on("from_server", data => {
    console.log(data)
    appendMessage(data, otherUserMessage, false)
})

form.addEventListener("submit", async e => {
   
})

inputBox.addEventListener("keydown", async e => {
    if (e.key !== "Enter" || inputBox.value.trim().length === 0)
        return

    const response =  await fetch(CHECK_EXPIRED_API_URL, {
        method: 'POST'
    })
    const result = await response.json()

    if(result["is_session_Expired"]){
        window.location.href = "/"
        return
    }
    const message = inputBox.value
    const messageData = {
        message_content: message,
        message_date: new Date().toLocaleString(),
        user_name: username,
        chatname: "public",
    }

    form.style.display = 'none'

    socket.emit("from_client", messageData)
    appendMessage(messageData, yourMessage, true)

    setInterval(() => {
        form.style.display = ''
    }, 1000)

    form.reset()
})

chatInner.addEventListener("click", e => {
    if(e.target.className === ellipseIconName){
        messageToBeRemoved = e.target.parentElement.parentElement
    }
})

removeMessageButton.addEventListener("click", async () => {
    const message_content = messageToBeRemoved.querySelector(".message").textContent
    const message_date    = messageToBeRemoved.querySelector(".date-and-comment-removal").textContent
    const messageToRemove = {
        message_content,
        message_date,
        username
    }

    messageToBeRemoved.remove()

    const response = await fetch(REMOVE_MESSAGE_API_URL, {
        method: "DELETE",
        body: JSON.stringify(messageToRemove),
        headers: {
            "Content-type": "application/json"
        }
    })

    const result = await response.json()

    console.log(result);

    $('#removeMessageModal').modal('hide')
})

async function getPublicMessages(){
    const response = await fetch(GET_MESSAGES_API_URL)
    const result   = await response.json()

    chatInner.innerHTML = ''
    result.forEach(messageBody => {
        const className = (messageBody.user_name === username) ? yourMessage : otherUserMessage
       
        appendMessage(messageBody, className, true)
    })
}