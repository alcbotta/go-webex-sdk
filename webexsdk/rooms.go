package webexsdk

import (
	"net/url"
	"strconv"
	"time"
)

type Rooms struct {
	Items []Room `json:"items"`
}

type Room struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Type         string    `json:"type"`
	IsLocked     bool      `json:"isLocked"`
	LastActivity time.Time `json:"lastActivity"`
	CreatorID    string    `json:"creatorId"`
	Created      time.Time `json:"created"`
	OwnerID      string    `json:"ownerId"`
}

type IRooms interface {
	GetRooms(teamID, roomType, sortBy string, max uint) ([]Room, error)
	GetRoom(roomID string) (Room, error)
}

/*
teamId
string
List rooms associated with a team, by ID.

type
string
List rooms by type.

Possible values: direct, group
sortBy
string
Sort results.

Possible values: id, lastactivity, created
max
number
Limit the maximum number of rooms in the response.

Default: 100
*/
func (room *Room) GetRooms(teamID, roomType, sortBy string, max uint) ([]Room, error) {
	rooms := &Rooms{}
	relativeURL := "/rooms/"
	u, err := url.Parse(relativeURL)
	if err != nil {
		return nil, err
	}

	queryString := u.Query()
	queryString.Set("teamID", teamID)
	queryString.Set("roomType", roomType)
	queryString.Set("sortBy", sortBy)
	queryString.Set("max", strconv.FormatUint(uint64(max), 10))
	u.RawQuery = queryString.Encode()

	err = Get(webexTokenID, baseURL.ResolveReference(u).String(), rooms)
	if err != nil {
		return nil, err
	}
	return rooms.Items, nil
}

func (room *Room) GetRoom(roomID string) (*Room, error) {
	tmpRoom := &Room{}
	relativeURL := "/rooms/"
	u, err := url.Parse(relativeURL + roomID)
	if err != nil {
		return nil, err
	}
	err = Get(webexTokenID, baseURL.ResolveReference(u).String(), tmpRoom)
	if err != nil {
		return nil, err
	}
	return tmpRoom, nil
}
