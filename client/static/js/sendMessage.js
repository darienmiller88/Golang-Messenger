const API_URL = "http://localhost:3333/socket.io/"
const inputMessage = document.querySelector(".input-message")
const chat = document.querySelector(".chat")
const chatInner = document.querySelector(".chat-inner")
const userInput = document.querySelector(".form-control")
const groupChats = document.querySelector(".group-chats")
const className = "group-chat-name"
const cc = "clicked-color" 
const groupChat = "group-chat"
const fullClassName = `${groupChat} ${cc}`
let counter = 0
// const socket = io.connect(API_URL)

// socket.on("chat-message", (data) => {
//     console.log(data)
//     appendMessage("other-message", data)
// })

// socket.on("disconnect", (reason) => {
//     socket.emit("disconnect", reason)
// });

inputMessage.addEventListener("submit", e => {
    e.preventDefault()
    const message = userInput.value

    appendMessage(counter % 2 === 0 ? "message" : "other-message", message)

    changeChatText(`darienmiller88: ${message}`)
    bubbleUpChat()
    counter++
   // socket.emit("send-chat-message", message)
    inputMessage.reset()
})

function appendMessage(className, message){
    const messageDiv = document.createElement('div')
    const dateDiv = document.createElement('div')
    const usernameDiv = document.createElement('div')

    //Append the following information to the above divs 
    messageDiv.append(message)
    usernameDiv.append("darienmiller88")
    dateDiv.append(new Date().toLocaleString())

    //Add the following class names to allow the appropriate styling to each div
    messageDiv.className = className
    usernameDiv.className = "username"
    dateDiv.className = "date"

    //Finally, add the username and date divs to the message div, and add the message div to the chat inner div
    messageDiv.append(usernameDiv)
    messageDiv.append(dateDiv)
    chatInner.append(messageDiv)

    //After the message has been added to the chat, scroll to the bottom automatically.
    chat.scrollTo(0, chat.scrollHeight)
}

function changeChatText(userMessage){
    document.querySelectorAll(`.${groupChat}`).forEach(chatDiv => {
        if(chatDiv.className === fullClassName){
            chatDiv.querySelector(".most-recent-message").innerHTML = userMessage
        }
    })
}

//The goal of this function is to "bubble up" a group chat up to the top anytime a message is sent when one 
//is "active".
function bubbleUpChat(){
    //If the first element in the group chats collection also happens to be already be the active one, simply
    //end the function call as nothing has to be "bubbled up".
    if(groupChats.querySelectorAll(`.${groupChat}`)[0].className === fullClassName){
        groupChats.scrollTo(0, 0)
        return
    }

    //In order to accomplish this, first turn the collection of divs from the "group-chats" class into an array
    let arrayOfGroupChats = [].slice.call(groupChats.querySelectorAll(`.${groupChat}`))

    //Afterwards, find the index of the group chat that is currently active.
    const indexOfClickedDiv = arrayOfGroupChats.findIndex(elem => {
        return elem.className === fullClassName
    })

    //Using the above index, remove it from the array, add it to front!
    let removed = arrayOfGroupChats.splice(indexOfClickedDiv, 1)

    //Finally, add the element that was removed to the from the array to  the front, and append this copy to
    //the actual list of group chat divs to reflect the changes.
    arrayOfGroupChats.unshift(removed[0])
    arrayOfGroupChats.forEach(elem => {
        groupChats.append(elem)
    })

    //After bubbling up a group chat, sidebar should scroll back up too!
    groupChats.scrollTo(0, 0)
}