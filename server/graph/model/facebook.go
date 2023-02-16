package model

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"

	fb "github.com/huandu/facebook"
)

type FacebookData struct {
	ID    string         `json:"id"`
	Posts *FacebookPosts `json:"posts"`
}

type FacebookPosts struct {
	Data   []*FacebookPost `json:"data"`
	Paging *FacebookPaging `json:"paging"`
}

type FacebookPost struct {
	ID          string            `json:"id"`
	FullPicture string            `json:"full_picture"`
	Comments    *FacebookComments `json:"comments"`
}

type FacebookComments struct {
	Data   []*FacebookComment `json:"data"`
	Paging *FacebookPaging    `json:"paging"`
}

type FacebookComment struct {
	ID          string              `json:"id"`
	CreatedTime string              `json:"created_time"`
	Message     string              `json:"message"`
	From        *FacebookUser       `json:"from"`
	Attachment  *FacebookAttachment `json:"attachment"`
	Comments    *FacebookComments   `json:"comments"`
}

type FacebookAttachment struct {
	Media struct {
		Image struct {
			Height int    `json:"height"`
			Src    string `json:"src"`
			Width  int    `json:"width"`
		} `json:"image"`
	} `json:"media"`
	Target struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"target"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type FacebookUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FacebookPaging struct {
	Cursors *FacebookCursors `json:"cursors"`
}

type FacebookCursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

func Facebook() FacebookData {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessToken := os.Getenv("REACT_APP_FACEBOOK_ACCESS_TOKEN")

	res, err := fb.Get("/emetetora?fields=posts{from,full_picture,message,comments{from,attachment,message,comments{from,message,comments{from,message}}}}", fb.Params{
		"fields":       "posts",
		"access_token": accessToken,
	})
	// res, err := fb.Get("/emetetora?fields=posts{full_picture,comments{message,from,comments}}", fb.Params{
	// 	"fields":       "posts",
	// 	"access_token": accessToken,
	// })

	if err != nil {
		fmt.Sprintf("There was an err in fb", err.Error())
	}

	jsonResult, _ := json.Marshal(res)
	fmt.Println(string(jsonResult))

	var output FacebookData
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &output,
		TagName:  "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)
	decoder.Decode(res)

	return output
}

func GettingFacebookData() (*FacebookData, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	accessToken := os.Getenv("REACT_APP_FACEBOOK_ACCESS_TOKEN")

	res, err := fb.Get("/emetetora?fields=posts{full_picture,comments{comments}}", fb.Params{
		"fields":       "posts,comments",
		"access_token": accessToken,
	})

	if err != nil {
		return nil, fmt.Errorf("error calling fb.Get: %w", err)
	}

	var output FacebookData
	if err := mapstructure.Decode(res, &output); err != nil {
		return nil, fmt.Errorf("error decoding fb.Get response: %w", err)
	}

	return &output, nil
}
