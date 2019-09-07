package defs

//user
type UserCredential struct {
	Username string	`json:"username" bson:"username"`
	PassWord string	`json:"password" bson:"username"`
	Email 	string	`json:"email" bson:"email"`
}
//video

type VideoInfo struct {
	Vid string	`json:"vid"`
	Uid int		`json:"uid"`
	Name string	`json:"name"`
	Ctime string`json:"ctime"`
}

type Comment struct {
	Cid string	`json:"cid"`
	Vid string	`json:"vid"`
	Uid string  `json:"uid"`
	Content string `json:"content"`
	Time string	`json:"time"`

}
type Session struct {
	Name string	`json:"name"`
	TTL int64	`json:"ttl"`
}

type SignedUp struct {
	Success bool `json:"success"`
	Sid string `json:"sid"`
}
