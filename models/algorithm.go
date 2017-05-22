package models

type Algorithm struct {
	Id          int64  `json:"id" orm:"column(ID)"`
	Name        string `json:"name" orm:"column(NAME)"`
	Version     string `json:"version" orm:"column(VERSION)"`
	Description string `json:"description" orm:"column(DESCRIPTION)"`
	Author      string `json:"author" orm:"column(AUTHOR)"`
	Parameters  string `json:"parameters" orm:"column(PARAMETERS)"` // p1=1,p2=...pn=n
	Image       string `json:"image" orm:"column(IMAGE)"`
	Downloads   int64  `json:"downloads" orm:"column(DOWNLOADS)"`
	Stars       int64  `json:"stars" orm:"column(STARS)"`
}
