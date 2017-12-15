package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
)

type ARWServer struct {
	serverSettings    ServerSettings
	events            ARWEvents
	listener          net.Listener
	sessionManager    SessionManager
	requestManager    RequestManager
	userManager       UserManager
	roomManager       RoomManager
	logManager        LogManager
	extensionHandlers []*ExtensionRequest
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

func (arwServer *ARWServer) AddExtensionHandler(cmd string, handler ExtensionHandler) error {
	for ii := 0; ii < len(arwServer.extensionHandlers); ii++ {
		currentExtension := arwServer.extensionHandlers[ii]
		if currentExtension.cmd == cmd {
			return errors.New("Extension Command already exist")
		}
	}

	var newExtension *ExtensionRequest
	newExtension = new(ExtensionRequest)

	newExtension.cmd = cmd
	newExtension.handler = handler
	arwServer.extensionHandlers = append(arwServer.extensionHandlers, newExtension)

	return nil
}

func (arwServer *ARWServer) SendExtensionRequest(cmd string, user *ARWUser, arwObj ARWObject) {
	var request *Request
	request = new(Request)

	request.eventname = Extension_Response
	request.arwObject = arwObj
	request.specialParams.PutString("cmd", cmd)
	request.specialParams.PutString("isRoomReq", "false")

	arwServer.SendRequest(request, user.session.conn)
}

func (arwServer *ARWServer) SendRequest(request *Request, conn net.Conn) {
	var reqData string
	reqData = "|"
	reqData += string(request.Compress())
	reqData += "|"
	conn.Write([]byte(reqData))
}
