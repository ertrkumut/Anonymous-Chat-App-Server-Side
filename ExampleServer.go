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

	if len(activeUsers) == 0 {
		//There is no player
		// Send CannotFindConversation Request
		return
	}

	owner := db.FindUserByARWUser(user)
	findedUser := activeUsers[0]

	var arwObject ARWObject
	arwObject.PutString("player_nickname", findedUser.nickname)
	arwObject.PutString("player_id", findedUser.id)

	arwServer.SendExtensionRequest(FindedConversation, user, arwObject)

	if findedUser.arwUser != nil {
		var obj ARWObject
		obj.PutString("player_nickname", owner.nickname)
		obj.PutString("player_id", owner.id)

		arwServer.SendExtensionRequest(FindedConversation, findedUser.arwUser, obj)
	}
}

func SendMessageHandler(server *ARWServer, user *ARWUser, arwObj ARWObject) {

}
