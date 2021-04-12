$('#newChatModal').on('show.bs.modal', e => {
    createNewChatButton.disabled = true
    createNewChatButton.classList.add("disabled")
})

$('#addChatMembersModal').on('show.bs.modal', e => {
    addNewMembersButton.disabled = true
    addNewMembersButton.classList.add("disabled")
})

$('#newUserModal').on('show.bs.modal', e => {
    createNewUserButton.disabled = true
    createNewUserButton.classList.add("disabled")
})

$('#viewAndRemoveModal').on('show.bs.modal', e => {
    chatMemberList.classList.add(chatMemberListClass)
})

//Modals use jQuery, so this function will specifically for that. When the modal is closed, it should clear
//all of the selected users the user clicked on.
$('#newChatModal').on('hidden.bs.modal', e => {
    removeSelectedUsers(newChatInput, dropDownNewGroup, selectedUsersNewGroup)
    groupChatNameInput.value = ""
})

$('#addChatMembersModal').on('hidden.bs.modal', e => {
    removeSelectedUsers(newUsersInput, dropDownNewUsers, selectedUsersNewUsers)
})

$('#newUserModal').on('hidden.bs.modal', e => {
    removeSelectedUsers(messageUserInput, dropDownMessageUser, selectedUserMessaged)
})

function removeSelectedUsers(inputBarValue, dropdown, selectedUsers){
    inputBarValue.value = ""
    filterUsers(inputBarValue, dropdown)
    selectedUsers.querySelectorAll(`.${selectedUserClass}`).forEach(user => {
        const username = user.querySelector(".username").innerText

        usernames[username] = false
        user.remove()
    })
}
