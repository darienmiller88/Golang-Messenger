let messageToBeRemoved

chatInner.addEventListener("click", e => {
    if(e.target.className === ellipseIconName){
        messageToBeRemoved = e.target.parentElement.parentElement
    }
})

removeMessageButton.addEventListener("click", () => {
    messageToBeRemoved.remove()
    $('#removeMessageModal').modal('hide')
})