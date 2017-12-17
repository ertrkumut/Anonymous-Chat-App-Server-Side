package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Player struct {
	id          string
	nickname    string
	password    string
	language    string
	createdData string
	talks       []*Talk
	arwUser     *ARWUser
	talkCounter int
}

func (player *Player) Init(userData []byte) error {
	player.talkCounter = 0

	var userMap map[string]interface{}

	if err := json.Unmarshal(userData, &userMap); err != nil {
		return err
	}

	player.id = fmt.Sprintf("%v", userMap["player_id"])
	player.nickname = fmt.Sprintf("%v", userMap["player_nickname"])
	player.language = fmt.Sprintf("%v", userMap["language"])
	player.createdData = fmt.Sprintf("%v", userMap["created_date"])
	player.password = fmt.Sprintf("%v", userMap["player_password"])

	if userMap["player_talks"] == nil {
		return nil
	}
	talks := userMap["player_talks"].([]interface{})
	for _, talk := range talks {
		talkData := talk.(map[string]interface{})

		var newTalk *Talk
		newTalk = new(Talk)
		newTalk.Init(talkData)
		player.talks = append(player.talks, newTalk)
	}

	fmt.Println(player.id, player.nickname, player.language, len(player.talks))
	return nil
}

func (player *Player) AddTalk(talk *Talk) {
	player.talks = append(player.talks, talk)

	db.UpdateUser(player)
}

func (player *Player) GetPlayerData() string {
	playerData := "{"
	playerData += "\"player_id\":\"" + player.id + "\","
	playerData += "\"player_nickname\":\"" + player.nickname + "\","
	playerData += "\"language\":\"" + player.language + "\","
	playerData += "\"created_date\":\"" + player.createdData + "\","
	playerData += "\"player_talks\":["

	for _, talk := range player.talks {
		playerData += talk.GetTalkData() + ","
	}
	playerData = strings.TrimRight(playerData, ",")
	playerData += "]}"

	return playerData
}

func (player *Player) GetAllData() string {
	userData := "{"
	userData += "\"player_id\":\"" + player.id + "\","
	userData += "\"player_nickname\":\"" + player.nickname + "\","
	userData += "\"player_password\":\"" + player.password + "\","
	userData += "\"language\":\"" + player.language + "\","
	userData += "\"created_date\":\"" + time.Now().Format(time.Stamp) + "\","
	userData += "\"player_talks\":["

	for _, talk := range player.talks {
		userData += talk.GetTalkData() + ","
	}
	userData = strings.TrimRight(userData, ",")
	userData += "]}"

	return userData
}
