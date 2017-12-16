package main

import (
	"fmt"
	"strconv"
)

type Talk struct {
	id           int64
	playerOneId  string
	playerTwoId  string
	receiverName string
	messages     []*Message
}

func (talk *Talk) Init(talkData map[string]interface{}) {

	idString := fmt.Sprintf("%v", talkData["talk_id"])
	talk.id, _ = strconv.ParseInt(idString, 10, 64)

	talk.playerOneId = fmt.Sprintf("%v", talkData["receiver_id"])
	talk.playerTwoId = fmt.Sprintf("%v", talkData["sender_id"])
	talk.receiverName = fmt.Sprintf("%v", talkData["receiver_name"])

	messages := talkData["talk_messages"].([]interface{})
	for _, val := range messages {
		msjData := val.(map[string]interface{})

		var msj *Message
		msj = new(Message)
		msj.Init(msjData)
		talk.messages = append(talk.messages, msj)
	}
}
