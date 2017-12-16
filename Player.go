package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Player struct {
	id          string
	nickname    string
	language    string
	createdData string
	talks       []*Talk
	arwUser     *ARWUser
}

func (player *Player) Init(userData []byte) error {

	var userMap map[string]interface{}

	if err := json.Unmarshal(userData, &userMap); err != nil {
		return err
	}

	player.id = fmt.Sprintf("%v", userMap["player_id"])
	player.nickname = fmt.Sprintf("%v", userMap["player_nickname"])
	player.language = fmt.Sprintf("%v", userMap["language"])
	player.createdData = fmt.Sprintf("%v", userMap["created_date"])

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
