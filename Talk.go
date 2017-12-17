package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Talk struct {
	id             int64
	ownerPlayer    string
	receiverPlayer string
	receiverName   string
	messages       []*Message
}

func (talk *Talk) Init(talkData map[string]interface{}) {

	idString := fmt.Sprintf("%v", talkData["talk_id"])
	talk.id, _ = strconv.ParseInt(idString, 10, 64)

	talk.ownerPlayer = fmt.Sprintf("%v", talkData["receiver_id"])
	talk.receiverPlayer = fmt.Sprintf("%v", talkData["sender_id"])
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

func (talk *Talk) CreateNewTalk(ownerPlayer *Player, playerTwo *Player) {
	talk.id = int64(len(ownerPlayer.talks))
	talk.ownerPlayer = ownerPlayer.id
	talk.receiverPlayer = playerTwo.id
	talk.receiverName = playerTwo.nickname

	ownerPlayer.talks = append(ownerPlayer.talks, talk)
}

func (talk *Talk) GetTalkData() string {
	talkData := "{"
	talkData += "\"talk_id\":" + fmt.Sprintf("%v", talk.id) + ","
	talkData += "\"receiver_name\":\"" + talk.receiverName + "\","
	talkData += "\"receiver_id\":\"" + talk.receiverPlayer + "\","
	talkData += "\"sender_id\":\"" + talk.ownerPlayer + "\","
	talkData += "\"talk_messages\":["

	for _, msj := range talk.messages {
		talkData += msj.GetMessageData() + ","
	}

	talkData = strings.TrimRight(talkData, ",")
	talkData += "]}"
	return talkData
}
