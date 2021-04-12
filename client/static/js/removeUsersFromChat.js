chatMemberList.addEventListener("click", e => {
    if(e.target.className === removeMemberBtnName){
        e.target.parentNode.remove()
        
        if(chatMemberList.children.length === 0){
            chatMemberList.classList.remove(chatMemberListClass)
        }
    }
})