let usernames = {
    "darienmiller88": false, 
    "nisey"         : false,
    "stevieboy"     : false,
    "vinny"         : false,
    "dp96"          : false,
    "mangs"         : false,
    "munckin"       : false,
    "peter"         : false,
    "neverbutter"   : false,
    "neverlutter"   : false,
    "nevergutter"   : false,
    "neverfutter"   : false,
    "neverttutter"  : false,
    "neverctrutter" : false,
    "neverdecutter" : false,
    "Thebigfatdog"  : false
}

//Form submission for creating a group chat
formGroupChat.addEventListener("submit", e => {
    e.preventDefault()
    appendUserToSelectedUsers(newChatInput.value)
    formGroupChat.reset()
})

//Form submission for adding a new user to a group chat
formAddUsers.addEventListener("submit", e => {
    e.preventDefault()
    appendUserToSelectedUsers(newUsersInput.value)
    formAddUser.reset()
})

//Form submisssion for messaging a new user
formMessageUser.addEventListener("submit", e => {
    e.preventDefault()
    appendUserToSelectedUsers(messageUserInput.value)
    formMessageUser.reset()
})

//Disable the "create new group chat" button when there is zero characters in the input box for entering 
//a group chat name
groupChatNameInput.addEventListener("keyup", () => {
    if(groupChatNameInput.value.split(" ").join("").length === 0){
        createNewChatButton.disabled = true
        createNewChatButton.classList.add("disabled")
    }else{
        createNewChatButton.disabled = false
        createNewChatButton.classList.remove("disabled")
    }
})

//Input bar for creating a new group chat
newChatInput.addEventListener("keyup", e => {
    filterUsers(newChatInput.value, dropDownNewGroup)
})

//Input bar for adding new users to a group chat
newUsersInput.addEventListener("keyup", e => {
    filterUsers(newUsersInput.value, dropDownNewUsers)
})

//Input bar for messaging a new user
messageUserInput.addEventListener("keyup", e => {
    //Only filter if the "selected-class" div has zero children.
    if(selectedUserMessaged.querySelectorAll(`.${selectedUserClass}`).length === 0){
        filterUsers(messageUserInput.value, dropDownMessageUser)
    }
})

//This function will be responsible for appending a list of users that contain a string typed in by the user.
//For example, typing in "n" will return darie(N)miller88, vi(N)(N)Y, and (N)isey
function filterUsers(filter, dropdown){   
    //First, remove any divs in the "users" div before filtering
    dropdown.querySelectorAll(`.${newUserClass}`).forEach(element => {
        dropdown.removeChild(element)
    })

    //If the value in the search bar amounts to an zero length string, remove the following
    //class from the div and end the function call. 
    if(filter.length === 0){
        dropdown.classList.remove(selectedStyling)
        return
    }

    appendUsersToDropDown(filter, dropdown)
}

function appendUsersToDropDown(filter, dropdown){
    for (const username in usernames) {
        //If the string that was typed by the user was found in any of the strings in the array,
        //and that particular name was not clicked yet, append it to the users div
        if(username.toLowerCase().indexOf(filter) > -1 && !usernames[username]){
            let userDiv = document.createElement("div")

            userDiv.className = newUserClass
            userDiv.append(username)
            dropdown.append(userDiv)
        }
    }

    //If there are no elements found in the "database" that satisfy the filter, remove the additional
    //styling around the drop down. Otherwise, add it!
    if(dropdown.querySelectorAll(`.${newUserClass}`).length === 0){
        dropdown.classList.remove(selectedStyling)
    }else{
        dropdown.classList.add(selectedStyling)
    }
}