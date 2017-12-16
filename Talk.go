package main

import (
	"fmt"
	"strconv"
	"strings"
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

func (talk *Talk) CreateNewTalk(owner *Player, playerTwo *Player) {
	talk.id = int64(len(owner.talks))
	talk.playerOneId = owner.id
	talk.playerTwoId = playerTwo.id
	talk.receiverName = playerTwo.nickname

	owner.talks = append(owner.talks, talk)
}

func (talk *Talk) GetTalkData() string {
	talkData := "{"
	talkData += "\"talk_id\":" + fmt.Sprintf("%v", talk.id) + ","
	talkData += "\"receiver_name\":\"" + talk.receiverName + "\","
	talkData += "\"receiver_id\":\"" + talk.playerTwoId + "\","
	talkData += "\"sender_id\":\"" + talk.playerOneId + "\","
	talkData += "\"talk_messages\":["

	for _, msj := range talk.messages {
		talkData += msj.GetMessageData() + ","
	}

	talkData = strings.TrimRight(talkData, ",")
	talkData += "]}"
	return talkData
}
