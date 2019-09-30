package test

import (
	"fmt"
	"github.com/go-webex-sdk/webexsdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InitWebex(t *testing.T) {
	webex, err := webexsdk.InitWebex(TConfig.TestTokenID)
	assert.Nil(t, err)
	rooms, err := webex.Rooms.GetRooms("", "", "", 100)
	assert.Nil(t, err)
	assert.NotNil(t, rooms)
	for i := range rooms {
		tmpRoom := rooms[i]
		fmt.Println(tmpRoom.ID)
	}
}

func Test_GetRoom_ExpectSuccess(t *testing.T) {
	webex, err := webexsdk.InitWebex(TConfig.TestTokenID)
	assert.Nil(t, err)
	room, err := webex.Rooms.GetRoom(TConfig.TestRoomID)
	assert.Nil(t, err)
	assert.NotNil(t, room)
}
