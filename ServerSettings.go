package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type ServerSettings struct {
	tcpPort                    string
	udpPort                    string
	defaultRoomName            string
	defaultRoomTag             string
	maximumUserCount           int64
	maximumRoomCount           int64
	defaultRoomCapacity        int64
	minimumLenghtOfuserName    int64
	maximumLenghtOfuserName    int64
	minimumLenghtOfroomName    int64
	maximumLenghtOfroomName    int64
	maxRoomExtensionRequests   int64
	maxServerExtensionRequests int64
}

func (setttings *ServerSettings) InitializeServerSettings(path string) {
	jsonBytes, jsonError := ioutil.ReadFile(path)
	if jsonError != nil {
		panic(jsonError)
	}

	var allJsonData map[string]interface{}
	if unMarshalError := json.Unmarshal(jsonBytes, &allJsonData); unMarshalError != nil {
		panic(unMarshalError)
	}

	setttings.tcpPort = fmt.Sprintf("%v", allJsonData["tcpPort"])
	setttings.udpPort = fmt.Sprint("%v", allJsonData["udpPort"])
	setttings.defaultRoomName = fmt.Sprintf("%v", allJsonData["defaultRoomName"])
	setttings.defaultRoomTag = fmt.Sprintf("%v", allJsonData["defaultRoomTag"])

	maximumUserCountString := fmt.Sprintf("%v", allJsonData["maximumUserCount"])
	maximumUserCount, parseErr := strconv.ParseInt(maximumUserCountString, 10, 64)
	if parseErr != nil {
		panic(parseErr)
	}
	setttings.maximumUserCount = maximumUserCount

	maximumRoomCountString := fmt.Sprintf("%v", allJsonData["maximumRoomCount"])
	maximumRoomCount, parseErr := strconv.ParseInt(maximumRoomCountString, 10, 64)
	if parseErr != nil {
		panic(parseErr)
	}
	setttings.maximumRoomCount = maximumRoomCount

	defaultRoomCapacityString := fmt.Sprintf("%v", allJsonData["defaultRoomCapacity"])
	defaultRoomCapacity, parseErr := strconv.ParseInt(defaultRoomCapacityString, 10, 64)
	if parseErr != nil {
		panic(parseErr)
	}
	setttings.defaultRoomCapacity = defaultRoomCapacity

	minimumLenghtOfroomNameString := fmt.Sprintf("%v", allJsonData["minimumLenghtOfroomName"])
	minimumLenghtOfroomName, parseErr := strconv.ParseInt(minimumLenghtOfroomNameString, 10, 64)
	if parseErr != nil {
		panic(parseErr)
	}
	setttings.minimumLenghtOfroomName = minimumLenghtOfroomName

	maximumLenghtOfroomNameString := fmt.Sprintf("%v", allJsonData["maximumLenghtOfroomName"])
	maximumLenghtOfroomName, parseErr := strconv.ParseInt(maximumLenghtOfroomNameString, 10, 64)
	if parseErr != nil {
		panic(parseErr)
	}
	setttings.maximumLenghtOfroomName = maximumLenghtOfroomName

	minimumLenghtOfuserNameString := fmt.Sprintf("%v", allJsonData["minimumLenghtOfuserName"])
	minimumLenghtOfuserName, parseErr := strconv.ParseInt(minimumLenghtOfuserNameString, 10, 64)
	if parseErr != nil {
		panic(parseErr)
	}
	setttings.minimumLenghtOfuserName = minimumLenghtOfuserName

	maximumLenghtOfuserNameString := fmt.Sprintf("%v", allJsonData["maximumLenghtOfuserName"])
	maximumLenghtOfuserName, parseErr := strconv.ParseInt(maximumLenghtOfuserNameString, 10, 64)
	if parseErr != nil {
		panic(parseErr)
	}
	setttings.maximumLenghtOfuserName = maximumLenghtOfuserName

	maxRoomExtensionRequestsString := fmt.Sprintf("%v", allJsonData["maxRoomExtensionRequests"])
	maxRoomExtensionRequests, parseErr := strconv.ParseInt(maxRoomExtensionRequestsString, 10, 64)
	if parseErr != nil {
		panic(parseErr)
	}
	setttings.maxRoomExtensionRequests = maxRoomExtensionRequests

	maxServerExtensionRequestsString := fmt.Sprintf("%v", allJsonData["maxRoomExtensionRequests"])
	maxServerExtensionRequests, parseErr := strconv.ParseInt(maxServerExtensionRequestsString, 10, 64)
	if parseErr != nil {
		panic(parseErr)
	}
	setttings.maxServerExtensionRequests = maxServerExtensionRequests
}
