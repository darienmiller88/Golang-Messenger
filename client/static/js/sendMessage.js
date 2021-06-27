const API_URL = "http://localhost:3333/socket.io/"
const fullClassName = `${groupChatClass} ${activeColor}`
let counter = 0
// const socket = io.connect(API_URL)

// socket.on("chat-message", (data) => {
//     console.log(data)
//     appendMessage("other-message", data)
// })

// socket.on("disconnect", (reason) => {
//     socket.emit("disconnect", reason)
// });

formInputMessage.addEventListener("submit", e => {
    e.preventDefault()
    const message = messageInput.value

    appendMessage(counter % 2 === 0 ? yourMessage : otherUserMessage, message)
    changeChatText(`darienmiller88: ${message}`)
    bubbleUpChat()
    counter++
   // socket.emit("send-chat-message", message)
   formInputMessage.reset()
})

function appendMessage(className, message){
    const messageDiv     = document.createElement('div')
    const dateDiv        = document.createElement('div')
    const usernameDiv    = document.createElement('div')
    const dateAndIconDiv = document.createElement('div')

    //Add the following class names to allow the appropriate styling to each div
    messageDiv.className     = className
    usernameDiv.className    = "username"
    dateDiv.className        = "date" 
    dateAndIconDiv.className = dateAndIconClassName
    
    //Append the following information to the above divs 
    messageDiv.append(message)
    usernameDiv.append("darienmiller88")
    dateDiv.append(new Date().toLocaleString())
    dateAndIconDiv.append(dateDiv)

    //If the message is coming from you, and not the user you are messaging, add an ellipse to the message 
    //allowing you to delete it.
    if(className === yourMessage){
        const ellipseIcon     = document.createElement('i')

        ellipseIcon.className = ellipseIconName
        ellipseIcon.setAttribute("data-toggle", "modal")
        ellipseIcon.setAttribute("data-target", "#removeMessageModal")
        dateAndIconDiv.append(ellipseIcon)
    }

    //Finally, add the username and date divs to the message div, and add the message div to the chat inner div
    messageDiv.append(usernameDiv, dateAndIconDiv)
    chatInner.append(messageDiv)

    //After the message has been added to the chat, scroll to the bottom automatically.
    chat.scrollTo(0, chat.scrollHeight)
}

function changeChatText(userMessage){
    document.querySelectorAll(`.${groupChatClass}`).forEach(chatDiv => {
        if(chatDiv.className === fullClassName){
            chatDiv.querySelector(".most-recent-message").textContent = userMessage
        }
    })
}

//The goal of this function is to "bubble up" a group chat up to the top anytime a message is sent when one 
//is "active".
function bubbleUpChat(){
    //If the first element in the group chats collection also happens to be already be the active one, simply
    //end the function call as nothing has to be "bubbled up".
    if(groupChats.querySelectorAll(`.${groupChatClass}`)[0].className === fullClassName){
        groupChats.scrollTo(0, 0)
        return
    }

    //In order to accomplish this, first turn the collection of divs from the "group-chats" class into an array
    let arrayOfGroupChats = [].slice.call(groupChats.querySelectorAll(`.${groupChatClass}`))

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