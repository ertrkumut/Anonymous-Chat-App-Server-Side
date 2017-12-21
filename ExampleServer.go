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

	arwServer.AddEventHandler(&(arwServer.events.Disconnection), DisconnectionEventHandler)

	arwServer.AddExtensionHandler("Login", LoginHandler)
	arwServer.AddExtensionHandler("Register", RegisterHandler)
	arwServer.AddExtensionHandler("Relogin", ReloginHandler)
	arwServer.AddExtensionHandler(SendMessage, SendMessageHandler)
	arwServer.AddExtensionHandler(FindConversation, FindConversationHandler)
	arwServer.Initialize()

	arwServer.ProcessEvents()
}

func DisconnectionEventHandler(arwObj ARWObject) {

}

func LoginHandler(server *ARWServer, user *ARWUser, arwObj ARWObject) {
	player_id, _ := arwObj.GetString("player_id")
	player_password, _ := arwObj.GetString("player_password")

	if db.UserIsExist(player_id) {
		userData, player, err := db.GetUserData(player_id)

		if err != nil || player == nil {
			return
		}

		if player.password != player_password {
			var wrongObj ARWObject
			arwServer.SendExtensionRequest("WrongPassword", user, wrongObj)
			return
		}

		fmt.Println("Login Success ", player_id)
		player.arwUser = user
		var obj ARWObject
		obj.PutString("player_data", userData)
		if err != nil {
			obj.PutString("error", err.Error())
		} else {
			obj.PutString("error", "")
		}

		server.SendExtensionRequest(GetUserData, user, obj)
	} else {
		var obj ARWObject
		arwServer.SendExtensionRequest("WrongPassword", user, obj)
	}
}

func ReloginHandler(server *ARWServer, user *ARWUser, arwObj ARWObject) {
	playerId, _ := arwObj.GetString("player_id")

	playerData, _, err := db.GetUserData(playerId)
	if err != nil {
		fmt.Println("Relogin Err ", playerId, err)
		return
	}

	var obj ARWObject
	obj.PutString("player_data", playerData)
	obj.PutString("error", "")
	arwServer.SendExtensionRequest("GetUserData", user, obj)
}

func RegisterHandler(server *ARWServer, user *ARWUser, arwObj ARWObject) {
	player_id, _ := arwObj.GetString("player_id")
	player_password, _ := arwObj.GetString("player_password")
	player_language, _ := arwObj.GetString("language")
	player_nickname, _ := arwObj.GetString("player_nickname")

	if db.UserIsExist(player_id) {
		var obj ARWObject
		arwServer.SendExtensionRequest("RegisterError", user, obj)
		return
	}

	userData, err := db.RegisterNewUser(player_id, player_nickname, player_language, player_password, user)
	var obj ARWObject

	if err != nil {
		obj.PutString("error", err.Error())
		arwServer.SendExtensionRequest("RegisterError", user, obj)
	} else {
		obj.PutString("player_data", userData)
		obj.PutString("error", "")
	}

	server.SendExtensionRequest("GetUserData", user, obj)
}

func FindConversationHandler(server *ARWServer, user *ARWUser, arwObj ARWObject) {
	activeUsers := db.GetActiveUsers()

	var findedUser *Player
	owner := db.FindUserByARWUser(user)
	if owner == nil {
		fmt.Println("user owner nil")
		return
	}

	for _, player := range activeUsers {
		if player.id != owner.id {
			conversationIsExist := '0'
			for _, talk := range owner.talks {
				if talk.receiverPlayer == player.id {
					conversationIsExist = '1'
				}
			}
			if conversationIsExist == '0' {
				findedUser = player
			}
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
	sender_id, _ := arwObj.GetString("sender_id")
	message_body, _ := arwObj.GetString("body")
	send_date, _ := arwObj.GetString("send_date")
	talkId, _ := arwObj.GetInt("talk_id")

	_, owner, err := db.GetUserData(sender_id)
	if err == nil {
		ownerTalk := owner.GetTalk(int64(talkId))
		if ownerTalk != nil {
			var ownerMsj *Message
			ownerMsj = new(Message)
			ownerMsj.NewMessage(sender_id, message_body, send_date)

			_, receiverPlayer, err := db.GetUserData(ownerTalk.receiverPlayer)

			if err == nil {
				for _, receiverTalk := range receiverPlayer.talks {
					if receiverTalk.receiverPlayer == owner.id {
						var receiverMsg *Message
						receiverMsg = new(Message)
						receiverMsg.NewMessage(owner.id, message_body, send_date)

						if receiverPlayer.arwUser != nil {
							fmt.Println("Send Msg : ", sender_id, message_body, talkId)
							receiverTalk.AddMessage(receiverMsg)
							ownerTalk.AddMessage(ownerMsj)

							db.UpdateUser(owner)
							db.UpdateUser(receiverPlayer)

							var ownerARWObj ARWObject
							ownerARWObj.PutString("message_data", ownerMsj.GetMessageData())
							arwServer.SendExtensionRequest("SendMessage", user, ownerARWObj)

							var receiverArwObj ARWObject
							receiverArwObj.PutString("message_data", receiverMsg.GetMessageData())
							arwServer.SendExtensionRequest("SendMessage", receiverPlayer.arwUser, receiverArwObj)
							return
						}
					}
				}
			}
		}
	}
}
