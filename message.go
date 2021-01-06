package main

type Message struct {
	Msgtype    string     `json:"msgtype"`
	Text       Text       `json:"text,omitempty"`
	Markdown   Markdown   `json:"markdown,omitempty"`
	Link       Link       `json:"link,omitempty"`
	ActionCard ActionCard `json:"actionCard,omitempty"`
	FeedCard   FeedCard   `json:"feedCard,omitempty"`
	At         At         `json:"at,omitempty"`
}

type At struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

type Text struct {
	Content string `json:"content"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type Link struct {
	Text       string `json:"text,omitempty"`
	Title      string `json:"title"`
	PicURL     string `json:"picUrl,omitempty"`
	MessageURL string `json:"messageUrl"`
}

type ActionCard struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	BtnOrientation string `json:"btnOrientation,omitempty"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
	Btns           []Btn  `json:"btns,omitempty"`
}

type Btn struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type FeedCard struct {
	Links []Link `json:"links"`
}

type Result struct {
	Rrrcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
