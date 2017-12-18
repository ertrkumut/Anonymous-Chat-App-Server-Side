package main

import (
	"fmt"
	"strconv"
)

type Message struct {
	id       int64
	body     string
	sendDate string
	senderId string
	talkId   int64
}

func (message *Message) Init(messageData map[string]interface{}) {

	idString := fmt.Sprintf("%v", messageData["message_id"])
	message.id, _ = strconv.ParseInt(idString, 10, 64)

	message.body = fmt.Sprintf("%v", messageData["body"])
	message.sendDate = fmt.Sprintf("%v", messageData["send_date"])
	message.senderId = fmt.Sprintf("%v", messageData["sender_id"])

	talkIdString := fmt.Sprintf("%v", messageData["talk_id"])
	message.talkId, _ = strconv.ParseInt(talkIdString, 10, 64)
}

func (message *Message) NewMessage(senderId string, body string, sendDate string) {
	message.senderId = senderId
	message.body = body
	message.sendDate = sendDate
}

func (message *Message) GetMessageData() string {
	messageData := "{"
	messageData += "\"message_id\":" + fmt.Sprintf("%v", message.id) + ","
	messageData += "\"body\":\"" + message.body + "\","
	messageData += "\"send_data\":\"" + message.sendDate + "\","
	messageData += "\"sender_id\":\"" + message.senderId + "\","
	messageData += "\"talk_id\":" + fmt.Sprintf("%v", message.talkId)
	messageData += "}"
	return messageData
}
