package main

import "fmt"

type YoutubeChannel struct {
	subscriberList map[string]Subscriber
	name string
}

func newYoutubeChannel(name string) *YoutubeChannel {
	return &YoutubeChannel{
		name: name,
	}
}

func (channel *YoutubeChannel) updateNewVideo() {
	fmt.Printf("New video is available")
	channel.notifyAll()
}

func (channel *YoutubeChannel) notifyAll() {
	for _, subscriber := range channel.subscriberList {
		subscriber.update(channel.name)
	}
}

func (channel *YoutubeChannel) subscribe(s Subscriber) {
	channel.subscriberList[s.getID()] = s
}

func (channel *YoutubeChannel) unsubscribe(s Subscriber) {
	channel.subscriberList = removeSubscriber(channel.subscriberList, s)
}

func removeSubscriber(subscriberList map[string]Subscriber, subscriberToRemove Subscriber) map[string]Subscriber {
	delete(subscriberList, subscriberToRemove.getID())
	return subscriberList
}