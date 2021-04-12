const iconClassName         = "bi bi-x-circle-fill"
const ellipseIconName       = "bi bi-three-dots-vertical"
const activeColor           = "clicked-color"
const groupChatNameClass    = "group-chat-name"
const lastMessageClassName  = "most-recent-message"
const newUserClass          = "new-user"
const selectedUserClass     = "selected-user"
const selectedStyling       = "selected-styling"
const dropdownClass         = "drop-down"
const groupChatClass        = "group-chat"
const chatTitleTextClass    = "title-text"
const chatMember            = "chat-member"
const removeMemberBtnName   = "btn btn-danger remove"
const chatMemberListClass   = "add-border"
const dateAndIconClassName  = "date-and-comment-removal"
const yourMessage           = "your-message"
const otherUserMessage      = "other-user-message"
const groupChatsContainer   = document.querySelector(".group-chats")
const chatTitleDiv          = document.querySelector(".chat-title")

const noGroupChatNameError  = document.querySelector(".error-message")

//Reference to the users in any given group chat
const chatMemberList        = document.querySelector(".chat-member-list")

//Buttons inside the three main modals. I love how all of them have the same length LMFAO.
const createNewChatButton   = document.getElementById("create-new-chat")
const createNewUserButton   = document.getElementById("create-new-user")
const addNewMembersButton   = document.getElementById("add-new-members")
const removeMessageButton   = document.getElementById("remove-comment-button")

//Forms
const formGroupChat         = document.querySelector(".form-group-chat")
const formMessageUser       = document.querySelector(".form-message-user")
const formAddUsers          = document.querySelector(".form-add-user")
const formInputMessage      = document.querySelector(".input-message-form")

//Input bars
const newUsersInput         = document.getElementById("add-new-members-input")
const newChatInput          = document.getElementById("add-new-chats")
const messageUserInput      = document.getElementById("message-new-user")
const groupChatNameInput    = document.getElementById("input-groupt-chat-name")
const messageInput          = document.querySelector(".form-control")

const dropDownNewGroup      = document.getElementById("drop-down-new-group-chat")
const dropDownNewUsers      = document.getElementById("drop-down-add-new-user")
const dropDownMessageUser   = document.getElementById("drop-down-message-user")

const selectedUsersNewGroup = document.getElementById("selected-users-new-group-chat")
const selectedUsersNewUsers = document.getElementById("selected-users-add-new-user")
const selectedUserMessaged  = document.querySelector(".user-messaged")

const chat                  = document.querySelector(".chat")
const chatInner             = document.querySelector(".chat-inner")
const groupChats            = document.querySelector(".group-chats")
