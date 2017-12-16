package main

import (
	"io/ioutil"
	"os"
	"time"
)

type DataBaseManager struct {
	dbPath string
	users  []*Player
}

func (db *DataBaseManager) InitAllDb() {
	allUsers := db.GetAllFiles()

	for _, userFile := range allUsers {
		userData, err := ioutil.ReadFile(db.dbPath + userFile.Name())
		if err == nil {
			var user *Player
			user = new(Player)

			user.Init(userData)
			db.users = append(db.users, user)
		}
	}
}

func (db *DataBaseManager) GetAllFiles() []os.FileInfo {
	files, err := ioutil.ReadDir(db.dbPath)

	if err != nil {
		return nil
	}

	return files
}

func (db *DataBaseManager) UserIsExist(userId string) bool {
	usersFiles := db.GetAllFiles()

	if usersFiles == nil {
		return false
	}

	for _, file := range usersFiles {
		if file.Name() == userId {
			return true
		}
	}

	return false
}

func (db *DataBaseManager) RegisterNewUser(userId string, nickname string, language string, arwUser *ARWUser) (string, error) {
	userData := "{"
	userData += "\"player_id\":\"" + userId + "\","
	userData += "\"player_nickname\":\"" + nickname + "\","
	userData += "\"language\":\"" + language + "\","
	userData += "\"created_date\":\"" + time.Now().Format(time.Stamp) + "\","
	userData += "\"player_talks\":[]}"

	err := ioutil.WriteFile(db.dbPath+userId, []byte(userData), 0644)

	var user *Player
	user = new(Player)
	user.Init([]byte(userData))
	user.arwUser = arwUser

	db.users = append(db.users, user)
	return userData, err
}

func (db *DataBaseManager) GetUserData(userId string) (string, *Player, error) {
	byteArray, err := ioutil.ReadFile(db.dbPath + userId)

	if err != nil {
		return "", nil, err
	}

	for _, user := range db.users {
		if user.id == userId {
			return string(byteArray), user, nil
		}
	}

	return "", nil, err
}

func (db *DataBaseManager) GetActiveUsers() []*Player {

	activeUsers := make([]*Player, 0)
	for _, player := range db.users {
		if player.arwUser != nil {
			activeUsers = append(activeUsers, player)
		}
	}
	return activeUsers
}
