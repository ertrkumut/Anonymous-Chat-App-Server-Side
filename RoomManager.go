package main

type RoomManager struct {
	allRooms      []*ARWRoom
	roomsTagByTag map[string][]*ARWRoom
	roomIdCounter int
}

func (roomManager *RoomManager) CreateRoom(arwServer *ARWServer, roomSettings *RoomSettings) *ARWRoom {
	var newRoom *ARWRoom
	newRoom = new(ARWRoom)

	newRoom.InitializeMethod = roomSettings.InitializeMethod

	if roomSettings.name != "" {
		newRoom.name = roomSettings.name
	} else {
		newRoom.name = arwServer.serverSettings.defaultRoomName
	}

	if roomSettings.tag != "" {
		newRoom.tag = roomSettings.tag
	} else {
		newRoom.tag = arwServer.serverSettings.defaultRoomTag
	}

	if roomSettings.cappacity != 0 {
		newRoom.cappacity = roomSettings.cappacity
	} else {
		newRoom.cappacity = int(arwServer.serverSettings.defaultRoomCapacity)
	}

	newRoom.id = roomManager.roomIdCounter
	roomManager.roomIdCounter++

	roomManager.allRooms = append(roomManager.allRooms, newRoom)
	if roomManager.roomsTagByTag == nil {
		roomManager.roomsTagByTag = make(map[string][]*ARWRoom)
	}
	roomManager.roomsTagByTag[newRoom.tag] = append(roomManager.roomsTagByTag[newRoom.tag], newRoom)

	newRoom.Init(arwServer)

	return newRoom
}

func (roomManager *RoomManager) FindRoomWithRoomId(roomId int) *ARWRoom {

	for ii := 0; ii < len(roomManager.allRooms); ii++ {
		if roomManager.allRooms[ii].id == roomId {
			return roomManager.allRooms[ii]
		}
	}

	return nil
}

func (roomManager *RoomManager) SearchRoomWithTag(tag string) *ARWRoom {
	for ii := 0; ii < len(roomManager.roomsTagByTag[tag]); ii++ {
		if roomManager.roomsTagByTag[tag][ii].IsFull() == false {
			return roomManager.roomsTagByTag[tag][ii]
		}
	}

	return nil
}
