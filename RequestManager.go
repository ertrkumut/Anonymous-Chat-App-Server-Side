package main

import "net"

type RequestManager struct {
	wronData string
	beginReq bool
}

func (reqManager *RequestManager) StartOrStopRequest(arwServer *ARWServer, conn net.Conn) {
	if reqManager.beginReq == false {
		reqManager.beginReq = true
		return
	}

	request := ExtractRequest([]byte(reqManager.wronData))
	reqManager.DoRequest(arwServer, request, conn)

	reqManager.wronData = ""
	reqManager.beginReq = false
}

func (reqManager *RequestManager) AddChar(reqChar string) {
	reqManager.wronData += reqChar
}

func (reqManager *RequestManager) DoRequest(arwServer *ARWServer, request *Request, conn net.Conn) {
	for ii := 0; ii < len(arwServer.events.allEvents); ii++ {
		currentEvent := arwServer.events.allEvents[ii]

		if currentEvent.eventName == request.eventname {
			currentEvent.Private_Handler(arwServer, conn, request)
			return
		}
	}
}
