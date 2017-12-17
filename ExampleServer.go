package main

import (
	"fmt"
)

var arwServer *ARWServer
var db *DataBaseManager

const (
	GetUserData          = "GetUserData"
	SendMessage          = "SendMessage"
	FindConversation     = "FindConversation"
	FindedConversation   = "FindedConversation"
	CannotFindActiveUser = "CannotFindActiveUser"
)

func main() {
	arwServer = new(ARWServer)
	db = new(DataBaseManager)
	db.dbPath = "ServerFiles/Database/"
	db.InitAllDb()

	arwServer.AddExtensionHandler(GetUserData, GetUserDataHandler)
	arwServer.AddExtensionHandler(SendMessage, SendMessageHandler)
	arwServer.AddExtensionHandler(FindConversation, FindConversationHandler)
	arwServer.Initialize()

	arwServer.ProcessEvents()
}

func GetUserDataHandler(server *ARWServer, user *ARWUser, arwObj ARWObject) {
	player_id, _ := arwObj.GetString("player_id")
	player_nickname, _ := arwObj.GetString("player_nickname")
	language, _ := arwObj.GetString("language")

	if !db.UserIsExist(player_id) {
		userData, err := db.RegisterNewUser(player_id, player_nickname, language, user)

		var obj ARWObject
		obj.PutString("player_data", userData)
		if err != nil {
			obj.PutString("error", err.Error())
		} else {
			obj.PutString("error", "")
		}

		x, _ := obj.GetString("player_data")
		fmt.Println(x)
		server.SendExtensionRequest(GetUserData, user, obj)
	} else {
		userData, player, err := db.GetUserData(player_id)
		player.arwUser = user

		var obj ARWObject
		obj.PutString("player_data", userData)
		if err != nil {
			obj.PutString("error", err.Error())
		} else {
			obj.PutString("error", "")
		}

		server.SendExtensionRequest(GetUserData, user, obj)
	}
}

func FindConversationHandler(server *ARWServer, user *ARWUser, arwObj ARWObject) {
	activeUsers := db.GetActiveUsers()

	var findedUser *Player
	owner := db.FindUserByARWUser(user)
	for _, player := range activeUsers {
		if player.id == owner.id {
			break
		}
		conversationIsExist := '0'
		for _, talk := range player.talks {
			if talk.receiverPlayer == owner.id {
				conversationIsExist = '1'
			}
		}
		if conversationIsExist == '0' {
			findedUser = player
		}
	}

	if findedUser == nil {
		var aObj ARWObject
		arwServer.SendExtensionRequest(CannotFindActiveUser, user, aObj)
		fmt.Println("Can not find Active User")
		return
	}
	fmt.Println("Conversation Find : ", owner.nickname, findedUser.nickname)
	var ownerTalk *Talk
	ownerTalk = new(Talk)
	ownerTalk.CreateNewTalk(owner, findedUser)
	owner.AddTalk(ownerTalk)

	var findedUserTalk *Talk
	findedUserTalk = new(Talk)
	findedUserTalk.CreateNewTalk(findedUser, owner)
	findedUser.AddTalk(findedUserTalk)

	var ownerObj ARWObject
	ownerObj.PutString("talk_data", ownerTalk.GetTalkData())

	arwServer.SendExtensionRequest(FindedConversation, user, ownerObj)

	if findedUser.arwUser != nil {
		var obj ARWObject
		obj.PutString("talk_data", findedUserTalk.GetTalkData())

		arwServer.SendExtensionRequest(FindedConversation, findedUser.arwUser, obj)
	}
}

func SendMessageHandler(server *ARWServer, user *ARWUser, arwObj ARWObject) {

}
