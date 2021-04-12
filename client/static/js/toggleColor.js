const activeState = (e) => {    
    //First, reference the chat div that was clicked
    const chatDiv = e.target.closest(`.${groupChatClass}`)

    //Afterwards, pull out the chat name, and the most recent message from that particular div into two
    //seperate strings to be sent to the server for validation. 
    const groupChatName = chatDiv.querySelector(`.${groupChatNameClass}`).innerText
    //const lastMessage   = chatDiv.getElementsByTagName("div")[1].innerText

    //When the user clicks on a chat, change the name of the chat title to reflect the name of the group chat
    chatTitleDiv.querySelector(`.${chatTitleTextClass}`).innerText = groupChatName

    //Next, remove the "clicked-color" class from each div in the chat side bar to return it to its base color
    document.querySelectorAll(`.${groupChatClass}`).forEach(chatDiv =>{
        if(chatDiv.className !== groupChatClass){
            chatDiv.classList.remove(activeColor)
        }
    })
    
    //Finally, target the div that was clicked, and add the "clicked-color" class to allow that div's color
    //to be changed from its base version to the color as described by the "clicked-color" styling.
    if(chatDiv.className === groupChatClass){
        chatDiv.classList.add(activeColor)
    }
}

groupChatsContainer.addEventListener('click', e => {  
   activeState(e)
})