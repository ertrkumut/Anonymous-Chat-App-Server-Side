package main

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	id       string
	nickname string
	language string
	talks    []*Talk
}

func (player *Player) Init(userData []byte) error {

	var userMap map[string]interface{}

	if err := json.Unmarshal(userData, &userMap); err != nil {
		return err
	}

	player.id = fmt.Sprintf("%v", userMap["player_id"])
	player.nickname = fmt.Sprintf("%v", userMap["player_nickname"])
	player.language = fmt.Sprintf("%v", userMap["language"])

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
