package HackerNewsBot

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

const (
	Best = "https://hacker-news.firebaseio.com/v0/beststories.json?print=pretty"
	New  = "https://hacker-news.firebaseio.com/v0/newstories.json?print=pretty"
	Top  = "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"
	Ask  = "https://hacker-news.firebaseio.com/v0/askstories.json?print=pretty"
	Show = "https://hacker-news.firebaseio.com/v0/showstories.json?print=pretty"
	Job  = "https://hacker-news.firebaseio.com/v0/jobstories.json?print=pretty"
)

type BotClient struct{}

func (client *BotClient) Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Print("Response status:", resp.Status)
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(bytes), nil
}

func (client *BotClient) GetItem(id int) (*Item, error) {
	var url string = "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(id) + ".json?print=pretty"
	log.Print(url)
	content, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
		return new(Item), err
	}
	var item *Item
	if err = json.Unmarshal([]byte(content), &item); err != nil {
		panic(err)
	}
	return item, nil
}

func (client *BotClient) GetUser(id string) (*User, error) {
	var url string = "https://hacker-news.firebaseio.com/v0/user/" + id + ".json?print=pretty"
	log.Print(url)
	content, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
		return new(User), err
	}
	var user *User
	if err = json.Unmarshal([]byte(content), &user); err != nil {
		panic(err)
	}
	return user, nil
}

func (client *BotClient) GetStors(storyType string) (*Stories, error) {
	content, err := client.Get(storyType)
	if err != nil {
		log.Fatal(err)
		return new(Stories), err
	}
	var stories *Stories
	if err = json.Unmarshal([]byte(content), &stories); err != nil {
		panic(err)
	}
	return stories, nil
}
