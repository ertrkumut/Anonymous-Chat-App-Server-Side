package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

type ARWServer struct {
	serverSettings ServerSettings
	events         ARWEvents
	listener       net.Listener
	sessionManager SessionManager
	requestManager RequestManager
	userManager    UserManager
	roomManager    RoomManager
	logManager     LogManager
}

func (arwServer *ARWServer) Initialize() {
	arwServer.serverSettings.InitializeServerSettings("ServerFiles/ServerProperties.json")
	arwServer.events.Initialize()

	arwServer.logManager.Init("ServerFiles/Log.json")
	tempListener, listenerError := net.Listen("tcp", arwServer.serverSettings.tcpPort)
	if listenerError != nil {
		panic(listenerError)
	}
	arwServer.listener = tempListener
	fmt.Println("ArwServer Initialize Success \n\n")
}

func (arwServer *ARWServer) ProcessEvents() {
	for {
		conn, acceptErr := arwServer.listener.Accept()

		if acceptErr != nil {
			fmt.Println("Error Accepting :", acceptErr)
		}

		go arwServer.HandleRequests(conn)
	}
}

func (arwServer *ARWServer) HandleRequests(conn net.Conn) {
	defer conn.Close()
	for {
		requestBytes := make([]byte, 1024)

		_, err := conn.Read(requestBytes)

		if err != nil {
			if err == io.EOF {
				arwServer.sessionManager.CloseSession(arwServer, conn)
				return
			}
		}

		requestBytes = bytes.Trim(requestBytes, "\x00")
		arwServer.ParseRequestBytes(requestBytes, conn)
	}
}

func (arwServer *ARWServer) ParseRequestBytes(bytes []byte, conn net.Conn) {
	message := string(bytes)

	for ii := 0; ii < len(message); ii++ {
		if message[ii] == '|' {
			arwServer.requestManager.StartOrStopRequest(arwServer, conn)
		} else {
			arwServer.requestManager.AddChar(string(message[ii]))
		}
	}
}

func (arwServer *ARWServer) AddEventHandler(event *ARWEvent, eventHandler EventHandler) {
	event.Handler = eventHandler
}

func (arwServer *ARWServer) SendRequest(request *Request, conn net.Conn) {
	var reqData string
	reqData = "|"
	reqData += string(request.Compress())
	reqData += "|"
	conn.Write([]byte(reqData))
}
