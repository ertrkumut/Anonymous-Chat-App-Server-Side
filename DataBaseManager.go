package main

import (
	"io/ioutil"
	"os"
	"time"
)

type DataBaseManager struct {
	dbPath string
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

func (db *DataBaseManager) RegisterNewUser(userId string, nickname string, language string) (string, error) {
	userData := "{"
	userData += "\"player_id\":\"" + userId + "\","
	userData += "\"player_nickname\":\"" + nickname + "\","
	userData += "\"language\":\"" + language + "\","
	userData += "\"created_date\":\"" + time.Now().Format(time.Stamp) + "\","
	userData += "\"player_talks\":[]}"

	err := ioutil.WriteFile(db.dbPath+userId, []byte(userData), 0644)

	return userData, err
}

func (db *DataBaseManager) GetUserData(userId string) (string, error) {
	byteArray, err := ioutil.ReadFile(db.dbPath + userId)

	return string(byteArray), err
}
