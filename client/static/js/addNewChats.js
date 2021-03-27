const iconClassName     = "bi bi-x-circle-fill"
const newUserClass      = "new-user"
const selectedUserClass = "selected-user"
const newClass          = "selected-styling"
const dropdownClass     = "drop-down"

const dropDownNewGroup      = document.getElementById("drop-down-new-group-chat")
const dropDownNewUser       = document.getElementById("drop-down-add-new-user")
const dropDownMessageUser   = document.getElementById("drop-down-message-user")
const selectedUsersNewGroup = document.getElementById("selected-users-new-group-chat")
const selectedUsersNewUser  = document.getElementById("selected-users-add-new-user")
const selectedUserMessaged  = document.getElementById("user-messaged")

const modalBody         = document.getElementById("exampleModal")
const newUsersInput     = document.getElementById("add-new-members")
const newChatInput      = document.getElementById("add-new-chats")
const messageUserInput  = document.getElementById("message-new-user")
const formGroupChat     = document.querySelector(".form-group-chat")
const formMessageUser   = document.querySelector(".form-message-user")
const formAddUser       = document.querySelector(".form-add-user")
const saveNewChat       = document.querySelector(".save")
const modalButtons      = document.querySelectorAll(".modal-btn")

let isCreatingGroupChat
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
formAddUser.addEventListener("submit", e => {
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

//Input bar for creating a new group chat
newChatInput.addEventListener("keyup", e => {
    filterUsers(newChatInput.value, dropDownNewGroup)
})

//Input bar for adding new users to a group chat
newUsersInput.addEventListener("keyup", e => {
    filterUsers(newUsersInput.value, dropDownNewUser)
})

//Input bar for messaging a new user
messageUserInput.addEventListener("keyup", e => {
    filterUsers(messageUserInput.value, dropDownMessageUser)
})

dropDownNewUser.addEventListener("click", e => {
    handleDropDownClick(e, dropDownNewUser, selectedUsersNewUser)
})

dropDownNewGroup.addEventListener("click", e => {
    handleDropDownClick(e, dropDownNewGroup, selectedUsersNewGroup)
})

dropDownMessageUser.addEventListener("click", e => {
    handleDropDownClick(e, dropDownMessageUser, selectedUserMessaged)
})

//Saving a new group chat to the database
saveNewChat.addEventListener("click", e => {

})

function handleDropDownClick(e, dropdown, selectedUsers){
    //When the user clicks on a username from the dropdown, set its clicked flag to true to signal
    //its later removal from the list of users 
    usernames[e.target.innerText] = true                
    e.target.remove()

    //If after removing, there are no more usernames in the dropdown, removing the additional styling
    if(dropdown.querySelectorAll(`.${newUserClass}`).length === 0){
        dropdown.classList.remove(newClass)
    }

    //Finally, add the username to the list of selected users.
    appendUserToSelectedUsers(e.target.innerText, selectedUsers)
}
// selectedUsersNG.addEventListener("click", e => {
    
// })

// selectedUsersNewUser.addEventListener("click", e => {

// })

function removeSelectedUser(){
    //Only remove the div if the user clicks on the x icon
    if(e.target.className === iconClassName){
        
        //When the user clicks to remove a username from the list, set its clicked flag back to false
        usernames[e.target.parentElement.innerText] = false
        e.target.parentElement.remove()
    }
}

//This function will be responsible for appending a list of users that contain a string typed in by the user.
//For example, typing in "n" will return darie(N)miller88, vi(N)(N)Y, and (N)isey
function filterUsers(filter, dropdown){
    console.log("function hit, filter: " + filter)
   
    //First, remove any divs in the "users" div before filtering
    dropdown.querySelectorAll(`.${newUserClass}`).forEach(element => {
        dropdown.removeChild(element)
    })

    //If the value in the search bar amounts to an zero length string, remove the following
    //class from the div and end the function call. 
    if(filter.length === 0){
        dropdown.classList.remove(newClass)
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
        dropdown.classList.remove(newClass)
    }else{
        dropdown.classList.add(newClass)
    }
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

//Modals use jQuery, so this function will specifically for that. When the modal is closed, it should clear
//all of the selected users the user clicked on.
$('#newChatModal').on('hidden.bs.modal', e => {
    newChatInput.value = ""
    filterUsers(newChatInput.value)
    selectedUsersNewGroup.querySelectorAll(`.${selectedUserClass}`).forEach(user => {
        user.remove()
    })
})

$('#addChatMembersModal').on('hidden.bs.modal', e => {
    newUsersInput.value = ""
    filterUsers(newUsersInput)
    selectedUsersNewUser.querySelectorAll(`.${selectedUserClass}`).forEach(user => {
        user.remove()
    })
})

$('#newUserModal').on('hidden.bs.modal', e => {
    
})
