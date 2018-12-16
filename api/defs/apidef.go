package defs

//requests
type UserCredential struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

//response
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

// data model
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

type SimpleSession struct {
	UserName string
	TTL      int64
}
