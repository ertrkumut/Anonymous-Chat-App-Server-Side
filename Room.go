package main

import (
	"errors"
	"strconv"
	"strings"
)

type RoomInitializeMethod func(arwServer *ARWServer, room *ARWRoom)

type ARWRoom struct {
	tag               string
	name              string
	id                int
	cappacity         int
	userList          []*ARWUser
	extensionRequests []*ExtensionRequest
	InitializeMethod  RoomInitializeMethod
}

func (room *ARWRoom) Init(arwServer *ARWServer) {
	room.userList = make([]*ARWUser, 0, room.cappacity)
	room.extensionRequests = make([]*ExtensionRequest, 0, arwServer.serverSettings.maxRoomExtensionRequests)

	if room.InitializeMethod != nil {
		room.InitializeMethod(arwServer, room)
	}
}

func (room *ARWRoom) AddUserInRoom(user *ARWUser, arwServer *ARWServer) error {

	if len(room.userList) >= room.cappacity {
		return errors.New("Room is full")
	}

	var userEnterRoomRequest *Request
	userEnterRoomRequest = new(Request)

	for ii := 0; ii < len(room.userList); ii++ {
		userEnterRoomRequest.eventname = User_Enter_Room

		userEnterRoomRequest.specialParams.PutString("user_properties", user.CompressUserProperties(room.userList[ii]))
		userEnterRoomRequest.specialParams.PutInt("room_id", room.id)
		arwServer.SendRequest(userEnterRoomRequest, room.userList[ii].session.conn)
	}

	room.userList = append(room.userList, user)
	user.lastRoom = room

	var roomJoinRequest *Request
	roomJoinRequest = new(Request)

	roomJoinRequest.eventname = Join_Room
	roomJoinRequest.specialParams.PutString("room_properties", room.CompressRoomSettings(user))
	arwServer.SendRequest(roomJoinRequest, user.session.conn)

	return nil
}

func (room *ARWRoom) RemoveUserInRoom(user *ARWUser, arwServer *ARWServer) {
	for ii := 0; ii < len(room.userList); ii++ {
		if room.userList[ii].id == user.id {
			room.userList = append(room.userList[:ii], room.userList[ii+1:]...)

			var userExitRoomRequest *Request
			userExitRoomRequest = new(Request)

			userExitRoomRequest.eventname = User_Exit_Room
			userExitRoomRequest.specialParams.PutInt("user_id", user.id)
			userExitRoomRequest.specialParams.PutInt("room_id", room.id)
			for ii := 0; ii < len(room.userList); ii++ {
				arwServer.SendRequest(userExitRoomRequest, room.userList[ii].session.conn)
			}
		}
	}
}

func (room *ARWRoom) AddExtensionHandler(cmd string, handler ExtensionHandler) error {
	for _, extension := range room.extensionRequests {
		if extension.cmd == cmd {
			return errors.New("Extension Already Exist")
		}
	}

	var newExtension *ExtensionRequest
	newExtension = new(ExtensionRequest)

	newExtension.cmd = cmd
	newExtension.handler = handler

	room.extensionRequests = append(room.extensionRequests, newExtension)
	return nil
}

func (room *ARWRoom) SendExtensionRequest(arwServer *ARWServer, cmd string, arwObj ARWObject, user *ARWUser) {
	var req *Request
	req = new(Request)

	req.eventname = Extension_Response
	req.arwObject = arwObj
	req.specialParams.PutString("cmd", cmd)
	req.specialParams.PutString("isRoom", "true")
	req.specialParams.PutInt("roomId", room.id)

	arwServer.SendRequest(req, user.session.conn)
}

func (room *ARWRoom) CompressRoomSettings(user *ARWUser) string {

	roomData := "{"
	roomData += "\"name\":\"" + room.name + "\","
	roomData += "\"tag\":\"" + room.tag + "\","
	roomData += "\"id\":" + strconv.Itoa(room.id) + ","
	roomData += "\"users\":["

	for ii := 0; ii < len(room.userList); ii++ {
		roomData += room.userList[ii].CompressUserProperties(user) + ","
	}

	roomData = strings.TrimRight(roomData, ",")
	roomData += "]}"
	return roomData
}

func (room *ARWRoom) IsFull() bool {
	if len(room.userList) >= room.cappacity {
		return true
	}
	return false
}
