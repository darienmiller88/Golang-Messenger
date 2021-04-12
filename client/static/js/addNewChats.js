//Saving a new group chat to the database
createNewChatButton.addEventListener("click", () => {
    appendNewGroupChat(groupChatNameInput.value)
    groupChatNameInput.value = ""
    newChatInput.value       = ""
    $('#newChatModal').modal('hide')
})

createNewUserButton.addEventListener("click", () => {
    const username = selectedUserMessaged.querySelector(`.${selectedUserClass}`).querySelector(".username").innerHTML
    console.log("you selected: " + username)
    appendNewGroupChat(username)
    messageUserInput.value = ""
    $('#newUserModal').modal('hide')
})

addNewMembersButton.addEventListener("click", e => {
    //This array will be sent to the server!
    const users = selectedUsersNewUsers.querySelectorAll(`.${selectedUserClass}`)

    appendUsersToMemberList(users)
    newUsersInput.value = ""
    $('#addChatMembersModal').modal('hide')
})

//newUserNames - array of all users to be added
function appendUsersToMemberList(newUserNames){
    let arrayOfUsers = [].slice.call(chatMemberList.querySelectorAll(`.${chatMember}`))

    newUserNames.forEach(elem => {
        chatMemberList.append(createNewUser(elem.querySelector(".username").innerHTML))
    })

    chatMemberList.classList.add(chatMemberListClass)
    // arrayOfUsers.unshift(createNewUser(newUserName))
    // arrayOfUsers.forEach(elem => {
    //     chatMemberList.append(elem)
    // }) 
}

function appendNewGroupChat(chatname){
    //first turn the collection of divs from the "group-chats" class into an array
    let arrayOfGroupChats = [].slice.call(groupChats.querySelectorAll(`.${groupChatClass}`))
    
    //Afterwards, remove the active color class from the first element that has it.
    arrayOfGroupChats.forEach(elem => {
        if(elem.classList.contains(activeColor)){
            elem.classList.remove(activeColor)
        }
    })

    //Add the new group chat div to the front of the array, and add the active color class to it
    arrayOfGroupChats.unshift(createNewGroupChat(chatname, ""))
    arrayOfGroupChats[0].classList.add(activeColor)

    //Finally, append the changes to the actual "group-chats" class.
    arrayOfGroupChats.forEach(elem => {
        groupChats.append(elem)
    }) 
}

function createNewGroupChat(groupChatName, mostRecentComment){
    let groupChatDiv         = document.createElement("div")
    let groupChatNameDiv     = document.createElement("div")
    let mostRecentMessageDiv = document.createElement("div")

    groupChatDiv.className         = groupChatClass
    groupChatNameDiv.className     = groupChatNameClass
    mostRecentMessageDiv.className = lastMessageClassName

    groupChatNameDiv.append(groupChatName)
    mostRecentMessageDiv.append(mostRecentComment)
    groupChatDiv.append(groupChatNameDiv)
    groupChatDiv.append(mostRecentMessageDiv)

    return groupChatDiv
}

function createNewUser(newUserName){
    let chatMemberDiv = document.createElement("div")
    let username      = document.createElement("h5")
    let removeButton  = document.createElement("button")

    removeButton.className  = "btn btn-danger remove"
    removeButton.type       = "button"
    chatMemberDiv.className = "chat-member"

    username.append(newUserName)
    removeButton.append("Remove")
    chatMemberDiv.append(username)
    chatMemberDiv.append(removeButton)

    return chatMemberDiv
}