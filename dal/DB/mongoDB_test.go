package DB

import (
	"chat/model"
	"context"
	"log"
	"testing"
)

func TestGetFriends(t *testing.T) {
	var friendList []model.Friend
	friendList = make([]model.Friend, 0)
	list, err := GetFriends(70786257168367616)
	if err != nil {
		log.Print(err)
	}
	for list.Next(context.Background()) {
		var friend model.Friend
		if err := list.Decode(&friend); err != nil {
			log.Print(err)
		}
		friendList = append(friendList, friend)
		log.Print(friend)
	}
}
