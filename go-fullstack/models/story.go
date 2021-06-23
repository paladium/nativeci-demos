package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Story stores a store from hackernews
type Story struct {
	By    string `json:"by"`
	ID    int    `json:"id"`
	Score int    `json:"score"`
	Time  int64  `json:"time"`
	Title string `json:"title"`
	Type  string `json:"type"`
	URL   string `json:"url"`
}

// LoadStory a single story
func LoadStory(id int) (*Story, error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	story := new(Story)
	err = json.Unmarshal(body, story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

// LoadTopStories top stories
func LoadTopStories() ([]Story, error) {
	url := "https://hacker-news.firebaseio.com/v0/newstories.json"
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	storyIDs := []int{}
	err = json.Unmarshal(body, &storyIDs)
	if err != nil {
		return nil, err
	}
	//Only take 10 latest
	storyIDs = storyIDs[0:10]
	stories := []Story{}
	for _, storyID := range storyIDs {
		story, err := LoadStory(storyID)
		if err != nil {
			return nil, err
		}
		stories = append(stories, *story)
	}
	return stories, nil
}
