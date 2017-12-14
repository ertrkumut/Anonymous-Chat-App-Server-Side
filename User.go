package main

import (
	"strconv"
)

type ARWUser struct {
	name     string
	id       int
	session  Session
	lastRoom *ARWRoom
}

func CreateUser(name string, id int, session Session) *ARWUser {
	var newUser *ARWUser
	newUser = new(ARWUser)

	newUser.name = name
	newUser.id = id
	newUser.session = session
	return newUser
}

func (user *ARWUser) CompressUserProperties(u *ARWUser) string {
	userData := "{"
	userData += "\"user_name\":\"" + user.name + "\","
	userData += "\"user_id\":" + strconv.Itoa(user.id) + ","

	if user.id == u.id {
		userData += "\"user_isMe\": \"true\""
	} else {
		userData += "\"user_isMe\": \"false\""
	}
	userData += "}"
	return userData
}

func (user *ARWUser) DestroyUser(arwServer *ARWServer) {
	arwServer.userManager.Removeuser(arwServer, user)

	if user.lastRoom != nil {
		user.lastRoom.RemoveUserInRoom(user, arwServer)
	}
}
