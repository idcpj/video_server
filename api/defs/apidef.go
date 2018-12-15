package defs

type UserCredential struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

type Comment struct {
	Id         string
	VideoId    string
	AuthorName string
	Content    string
}
