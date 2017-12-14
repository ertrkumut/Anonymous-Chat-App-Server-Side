package main

import "net"
import "errors"

type UserManager struct {
	allUser       []*ARWUser
	userIdCounter int
}

func (userManager *UserManager) CreateUser(arwServer *ARWServer, conn net.Conn, userName string) (*ARWUser, error) {

	for ii := 0; ii < len(arwServer.sessionManager.allSessions); ii++ {
		if arwServer.sessionManager.allSessions[ii].conn == conn {
			newUser := CreateUser(userName, userManager.userIdCounter, *(arwServer.sessionManager.allSessions[ii]))
			userManager.userIdCounter++
			userManager.allUser = append(userManager.allUser, newUser)
			return newUser, nil
		}
	}

	return nil, errors.New("Session Not Found!")
}

func (userManager *UserManager) Removeuser(arwServer *ARWServer, user *ARWUser) {
	for ii := 0; ii < len(userManager.allUser); ii++ {
		if userManager.allUser[ii] == user {
			userManager.allUser = append(userManager.allUser[:ii], userManager.allUser[ii+1:]...)
		}
	}
}

func (userManager *UserManager) FindUserWithId(userId int) (*ARWUser, error) {

	for ii := 0; ii < len(userManager.allUser); ii++ {
		if userManager.allUser[ii].id == userId {
			return userManager.allUser[ii], nil
		}
	}

	return nil, errors.New("User Does not Exist")
}

func (userManager *UserManager) FindUserWithSession(conn net.Conn) (*ARWUser, error) {

	for ii := 0; ii < len(userManager.allUser); ii++ {
		if userManager.allUser[ii].session.conn == conn {
			return userManager.allUser[ii], nil
		}
	}

	return nil, errors.New("User Does not Exist")
}
