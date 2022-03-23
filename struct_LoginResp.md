#1.struct LoginResp

```go
type LoginResp struct {
	UserId        int64  `json:"userid"`
	Email         string `json:"email"`
	UserRole      int64  `json:"-"`
	Channel       string `json:"-"`
	IsActivated   int64  `json:"is_activated"`
	IsBlocked     int64  `json:"-"`
	IsDeleted     int64  `json:"-"`
	IsBot         int64  `json:"-"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Avatar        string `json:"avatar"`
	Username      string `json:"username"`
	AboutYou      string `json:"about_you"`
	Location      string `json:"location"`
	Website       string `json:"website"`
	FullName      string `json:"full_name"`
	Twitter       string `json:"twitter"`
	Facebook      string `json:"facebook"`
	Linkedin      string `json:"linkedin"`
	Interest      string `json:"interest"`
	LastLoginTime int64  `json:"last_login_time"`
	FollowerCnt   int64  `json:"follower_cnt"`
	FolloweeCnt   int64  `json:"followee_cnt"`
	CreatedAt     int64  `json:"-"`
}

```