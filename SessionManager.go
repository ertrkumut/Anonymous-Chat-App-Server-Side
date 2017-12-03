package main

import (
	"fmt"
	"net"
)

type SessionManager struct {
	allSessions []*Session
}

func (sessionManager *SessionManager) StartSession(conn *net.Conn) {
	var session *Session
	session = new(Session)

	session.Init(conn, sessionManager)
}

func (sessionManager *SessionManager) CloseSession(arwServer *ARWServer, conn net.Conn) {

	user, err := arwServer.userManager.FindUserWithSession(conn)

	if err == nil {
		user.DestroyUser(arwServer)
		fmt.Println("Session Close : ", conn)
		conn.Close()
	}
}
