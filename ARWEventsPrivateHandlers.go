package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func P_ConnectionSuccess(arwServer *ARWServer, conn net.Conn, request *Request) {
	currentTime := time.Now().Format(time.StampMilli)
	oclock := strings.Split(string(currentTime), " ")[2]

	arwServer.sessionManager.StartSession(&conn)

	var connectionRequest *Request
	connectionRequest = new(Request)

	connectionRequest.eventname = Connection_Success
	connectionRequest.specialParams.PutString("server_time", oclock)
	connectionRequest.specialParams.PutString("error", "")

	arwServer.SendRequest(connectionRequest, conn)

	if arwServer.events.Connection.Handler == nil {
		return
	}
	arwServer.events.Connection.Handler(request.arwObject)
}

func P_Disconnection(arwServer *ARWServer, conn net.Conn, request *Request) {
	arwServer.sessionManager.CloseSession(arwServer, conn)

	if arwServer.events.Disconnection.Handler == nil {
		return
	}

	arwServer.events.Disconnection.Handler(request.arwObject)
}

func P_Login(arwServer *ARWServer, conn net.Conn, request *Request) {

	userName, _ := request.specialParams.GetString("user_name")
	user, err := arwServer.userManager.CreateUser(arwServer, conn, userName)

	if err != nil {
		fmt.Println("User Create Error :", err)
		return
	}

	if arwServer.events.Login.Handler != nil {
		request.arwObject.PutInt("user_id", user.id)
		arwServer.events.Login.Handler(request.arwObject)
	}

	var responseRequest *Request
	responseRequest = new(Request)

	responseRequest.eventname = Login
	responseRequest.specialParams.PutString("user_properties", user.CompressUserProperties(user))

	arwServer.SendRequest(responseRequest, conn)
}

func P_ExtensionResponse(arwServer *ARWServer, conn net.Conn, request *Request) {
	cmd, _ := request.specialParams.GetString("cmd")
	isRoomReq, _ := request.specialParams.GetString("isRoom")

	if isRoomReq == "False" {
		for _, extension := range arwServer.extensionHandlers {
			if cmd == extension.cmd {
				user, err := arwServer.userManager.FindUserWithSession(conn)
				if err == nil {
					extension.handler(arwServer, user, request.arwObject)
				}
			}
		}
	} else {
		roomId, err := request.specialParams.GetInt("roomId")
		if err != nil {
			return
		}

		room := arwServer.roomManager.FindRoomWithRoomId(roomId)
		for _, extension := range room.extensionRequests {
			if extension.cmd == cmd {
				user, userErr := arwServer.userManager.FindUserWithSession(conn)
				if userErr != nil {
					return
				}

				extension.handler(arwServer, user, request.arwObject)
			}
		}
	}
}
