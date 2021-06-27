const chat = document.querySelector(".chat")
const chatInner = document.querySelector(".chat-inner")
const yourMessage      = "your-message"
const otherUserMessage = "other-user-message"

const appendMessage = (messageData, messageClassName, isMessageRemovable) => {
    const message = createMessage(messageData, messageClassName, isMessageRemovable)
    chatInner.append(message)

    chat.scrollTo(0, chat.scrollHeight)
}

const createMessage = (messageData, messageClassName, isMessageRemovable) => {
    const messageBodyDiv = document.createElement('div')
    const messageDiv     = document.createElement('div')
    const dateDiv        = document.createElement('div')
    const usernameDiv    = document.createElement('div')
    const dateAndIconDiv = document.createElement('div')
    

    messageBodyDiv.className = messageClassName
    usernameDiv.className    = "message-username"
    dateDiv.className        = "date" 
    dateAndIconDiv.className = "date-and-comment-removal"
    messageDiv.className     = "message"

    messageDiv.append(messageData.message_content)
    usernameDiv.append(messageData.user_name)
    dateDiv.append(messageData.message_date)
    dateAndIconDiv.append(dateDiv)

    if(messageBodyDiv.className === yourMessage && isMessageRemovable){
        const ellipseIcon     = document.createElement('i')

        ellipseIcon.className = ellipseIconName
        ellipseIcon.setAttribute("data-toggle", "modal")
        ellipseIcon.setAttribute("data-target", "#removeMessageModal")
        dateAndIconDiv.append(ellipseIcon)
    }

    messageBodyDiv.append(messageDiv, usernameDiv, dateAndIconDiv)

    return messageBodyDiv
}