 const socket = io.connect("http://localhost:3333/socket.io")

const form = document.querySelector("form")
const input = document.getElementById("chat")

socket.on("new_message", data => {
    console.log(data)
})

socket.on("disconnect", () => {
    socket.emit("disconnect", "user disconnected")
})

form.addEventListener("submit", e => {
    e.preventDefault()
    const message = input.value

    console.log(message)
    socket.emit("message", message)
    form.reset()
})
