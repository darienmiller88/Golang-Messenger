const API_URL = "http://localhost:3333/socket.io"
const form = document.querySelector("form")
const chatInput = document.getElementById("chat")
const socket = io.connect(API_URL)

socket.on("chat-message", (data) => {
    console.log(data)
    appendMessage("other-message", data)
})

socket.on("disconnect", (reason) => {
    socket.emit("disconnect", reason)
});

form.addEventListener("submit", e => {
    e.preventDefault()
    console.log("message: ", chatInput.value)
    const message = chatInput.value

    appendMessage("message", message)
    socket.emit("send-chat-message", message)
    form.reset()
})

function appendMessage(className, message){
    const div = document.createElement('div')

    div.append(message)
    div.className = className
    chatbox.append(div)
    chatbox.scrollTo(0, chatbox.scrollHeight)
}