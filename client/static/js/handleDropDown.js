dropDownNewGroup.addEventListener("mousedown", e => {
    handleDropDownClick(e, dropDownNewGroup, selectedUsersNewGroup)
})

dropDownNewUsers.addEventListener("mousedown", e => {
    handleDropDownClick(e, dropDownNewUsers, selectedUsersNewUsers)
    enableButton(addNewMembersButton, selectedUsersNewUsers)
})

dropDownMessageUser.addEventListener("mousedown", e => {
    handleDropDownClick(e, dropDownMessageUser, selectedUserMessaged)
    enableButton(createNewUserButton, selectedUserMessaged)
    
    //After clicking on a username from the dropdown, remove the styling and the usernames.
    dropDownMessageUser.classList.remove(selectedStyling)
    dropDownMessageUser.querySelectorAll(`.${newUserClass}`).forEach(elem => {
        elem.remove()
    })
})

//Adding a new group to the list of groupchats
selectedUsersNewGroup.addEventListener("click", e => {   
    removeSelectedUser(e, selectedUsersNewGroup, createNewChatButton)
})

//Adding ONE other user to privately message
selectedUserMessaged.addEventListener("click", e => {
    removeSelectedUser(e, selectedUserMessaged, createNewUserButton)
    disableButton(createNewUserButton, selectedUserMessaged)
})

//Adding new users to a group chat
selectedUsersNewUsers.addEventListener("click", e => {
    removeSelectedUser(e, selectedUsersNewUsers, addNewMembersButton)
    disableButton(addNewMembersButton, selectedUsersNewUsers) 
})

function disableButton(button, selectedUsers){
    if(selectedUsers.childNodes.length === 0){
        button.disabled = true
        button.classList.add("disabled")
    }
}

function enableButton(button, selectedUsers){
    if(selectedUsers.childNodes.length > 0){
        button.disabled = false
        button.classList.remove("disabled")
    }
}

function removeSelectedUser(e, selectedUsers, button){
    //Only remove the div if the user clicks on the x icon
    if(e.target.className === iconClassName){

        //When the user clicks to remove a username from the list, set its clicked flag back to false
        usernames[e.target.parentElement.innerText] = false
        e.target.parentElement.remove()
    }
}


function handleDropDownClick(e, dropdown, selectedUsers){
    //When the user clicks on a username from the dropdown, set its clicked flag to true to signal
    //its later removal from the list of users 
    usernames[e.target.innerText] = true                
    e.target.remove()

    //If after removing, there are no more usernames in the dropdown, removing the additional styling
    if(dropdown.querySelectorAll(`.${newUserClass}`).length === 0){
        dropdown.classList.remove(selectedStyling)
    }

    //Finally, add the username to the list of selected users.
    appendUserToSelectedUsers(e.target.innerText, selectedUsers)
}

function appendUserToSelectedUsers(userName, selectedUsers){
    selectedUsers.append(makeNewUser(userName))
}

function makeNewUser(username){
    let newUser     = document.createElement("div")
    let userNameDiv = document.createElement("div")
    let icon        = document.createElement("i")

    icon.className        = iconClassName
    userNameDiv.className = "username"
    newUser.className     = "selected-user"

    userNameDiv.append(username)
    newUser.append(userNameDiv)
    newUser.append(icon)

    return newUser
}