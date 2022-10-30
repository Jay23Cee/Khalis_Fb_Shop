package model

import (
	"fmt"

	"github.com/mitchellh/mapstructure"

	fb "github.com/huandu/facebook"
)

type FB_POSTdata struct {
	Attachments struct {
		Data []struct {
			Media struct {
				Image struct {
					Height int    `json:"height"`
					Src    string `json:"src"`
					Width  int    `json:"width"`
				} `json:"image"`
			} `json:"media"`
			Description string `json:"description"`
		} `json:"data"`
	} `json:"attachments"`
	Comments struct {
		Data []struct {
			Message    string `json:"message"`
			ID         string `json:"id"`
			Attachment struct {
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
			} `json:"attachment,omitempty"`
			Comments struct {
				Data []struct {
					CreatedTime string `json:"created_time"`
					From        struct {
						Name string `json:"name"`
						ID   string `json:"id"`
					} `json:"from"`
					Message string `json:"message"`
					ID      string `json:"id"`
				} `json:"data"`
				Paging struct {
					Cursors struct {
						Before string `json:"before"`
						After  string `json:"after"`
					} `json:"cursors"`
				} `json:"paging"`
			} `json:"comments,omitempty"`
		} `json:"data"`
		Paging struct {
			Cursors struct {
				Before string `json:"before"`
				After  string `json:"after"`
			} `json:"cursors"`
		} `json:"paging"`
	} `json:"comments"`
	ID string `json:"id"`
}

func Facebook() FB_POSTdata {

	res, err := fb.Get("/108798115240903_135911885872707?fields=attachments{media,description},comments{attachment, message, comments}", fb.Params{
		"fields": "attachments,comments",
		//"access_token": "EAAdbFKpjrhYBAHVRnyj0eLaIMQW0ymcEmLI5f9FF8iOFPtr0Kj04Dhb9Y91BZByV79ZAEIKZAA5Iy9fVtEXdCazkLn0lDp0sXVCirYt35KpfZCEhj0T7YQLdTZBIZB1eItT5zIsDHpJoZBinyTxZC3jZBxuBUZAVm2NsNZBPSVa81TZAQuSJBXNjeDcEGZArz76BilV3sZCyaWjWjaQ1Jd8js1PwcoL2YSNQpJxoL7kVtt8R3u13873PZBQmZA7YgEXVhrP73EYZD",
		"access_token": "EAAdbFKpjrhYBAOeNlkc5LxFeZA5dmJXTl7Q0peKDckTsV8HGQNlvoWLhOMqmQf1ZAGZB4cathK3D6GqKeukUENd9ur9dxiZB8ZAdkRAcov5nEZCVE7ZCONEYiTINLHjbxvtRBxQcF7TKOMLOHIhYZA55cPfxCOrIpr4rYXdCz0bZBZBAWFyxl54W3rgc9dlvB9RCYZD",
	})

	if err != nil {
		fmt.Sprintf("There was an err in fb", err.Error())
	}

	var output FB_POSTdata
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &output,
		TagName:  "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)
	decoder.Decode(res)

	fmt.Println("boolean: ", output)
	return output
}


