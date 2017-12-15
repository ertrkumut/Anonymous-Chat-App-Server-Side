package main

import "fmt"

var arwServer *ARWServer
var db *DataBaseManager

func main() {
	arwServer = new(ARWServer)
	db = new(DataBaseManager)
	db.dbPath = "ServerFiles/"

	arwServer.AddEventHandler(&(arwServer.events.Login), Login_Event_Handler)
	arwServer.AddExtensionHandler("GetUserData", GetUserDataHandler)
	arwServer.Initialize()

	arwServer.ProcessEvents()
}

func Login_Event_Handler(arwObj ARWObject) {

	// user, _ := arwObj.GetUser(arwServer)
}

func GetUserDataHandler(server *ARWServer, user *ARWUser, arwObj ARWObject) {
	player_id, _ := arwObj.GetString("player_id")
	player_nickname, _ := arwObj.GetString("player_nickname")
	language, _ := arwObj.GetString("language")

	if !db.UserIsExist(player_id) {
		userData, err := db.RegisterNewUser(player_id, player_nickname, language)

		var obj ARWObject
		obj.PutString("player_data", userData)
		if err != nil {
			obj.PutString("error", err.Error())
		} else {
			obj.PutString("error", "")
		}

		x, _ := obj.GetString("player_data")
		fmt.Println(x)
		server.SendExtensionRequest("GetUserData", user, obj)
	} else {
		userData, err := db.GetUserData(player_id)

		var obj ARWObject
		obj.PutString("player_data", userData)
		if err != nil {
			obj.PutString("error", err.Error())
		} else {
			obj.PutString("error", "")
		}

		server.SendExtensionRequest("GetUserData", user, obj)
	}
}
