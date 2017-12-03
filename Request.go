package main

import "strings"

type Request struct {
	eventname     string
	arwObject     ARWObject
	specialParams SpecialEventParam
}

func newRequest(eventName string, arwObj ARWObject) *Request {
	var newRequest *Request
	newRequest = new(Request)

	newRequest.eventname = eventName
	newRequest.arwObject = arwObj
	return newRequest
}

func (req *Request) Compress() []byte {
	var requestData string
	requestData = req.eventname + "^^"

	requestData += string(req.arwObject.Compress()) + "^^"
	requestData += req.specialParams.Compress()

	return []byte(requestData)
}

func ExtractRequest(bytes []byte) *Request {
	requestData := string(bytes)

	var newReq *Request
	newReq = new(Request)

	requestParams := strings.Split(requestData, "^^")
	if len(requestParams) == 3 {
		newReq.eventname = requestParams[0]

		var obj ARWObject
		obj.Extract([]byte(requestParams[1]))
		var specialObj SpecialEventParam
		specialObj.Extract(requestParams[2])

		newReq.arwObject = obj
		newReq.specialParams = specialObj
	}

	return newReq
}
