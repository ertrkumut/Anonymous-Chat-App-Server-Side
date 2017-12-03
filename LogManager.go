package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type LogManager struct {
	logFile string
	allLogs []*Log
}

func (logManager *LogManager) Init(logFilePath string) {
	logManager.logFile = logFilePath

	err := ioutil.WriteFile(logFilePath, nil, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func (logManager *LogManager) UpdateLogsFile() {
	logsData := "["
	for ii := 0; ii < len(logManager.allLogs); ii++ {
		logData, err := logManager.allLogs[ii].ConvertJson()
		if err == nil {
			logsData += string(logData) + ","
		}
	}

	logsData = strings.TrimRight(logsData, ",")
	logsData += "]"

	ioutil.WriteFile(logManager.logFile, []byte(logsData), 0644)
}

type Log struct {
	log_type    string
	color       string
	log_time    string
	log_message string
}

func (log *Log) ConvertJson() ([]byte, error) {
	logMap := map[string]string{"log_type": log.log_type, "color": log.color, "log_time": log.log_time, "log_message": log.log_message}

	logBytes, err := json.Marshal(logMap)
	return logBytes, err
}

func newLog(message string, logManager *LogManager) *Log {
	var newLog *Log
	newLog = new(Log)

	newLog.log_type = "User_Log"
	newLog.color = "is-info"
	newLog.log_message = message
	newLog.log_time = time.Now().Format(time.StampMilli)

	logManager.allLogs = append(logManager.allLogs, newLog)
	logManager.UpdateLogsFile()

	return newLog
}
