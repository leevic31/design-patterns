package main

import "fmt"

type YoutubeSubscriber struct {
	id string
}

func (s *YoutubeSubscriber) update(channelName string) {
	fmt.Printf("Sending notification to subscriber %s for channel %s", s.id, channelName)	
}

func (s *YoutubeSubscriber) getID() string {
	return s.id
}