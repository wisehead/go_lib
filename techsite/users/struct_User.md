#1.struct User

```
	User struct {
		Id            int64     `db:"id"`
		Email         string    `db:"email"`
		FullName      string    `db:"full_name"`
		Password      string    `db:"password"`
		UserRole      int64     `db:"user_role"`
		Channel       string    `db:"channel"`
		IsActivated   int64     `db:"is_activated"`
		IsBlocked     int64     `db:"is_blocked"`
		IsDeleted     int64     `db:"is_deleted"`
		IsBot         int64     `db:"is_bot"`
		FirstName     string    `db:"first_name"`
		LastName      string    `db:"last_name"`
		Avatar        string    `db:"avatar"`
		Username      string    `db:"username"`
		AboutYou      string    `db:"about_you"`
		Location      string    `db:"location"`
		Website       string    `db:"website"`
		Twitter       string    `db:"twitter"`
		Facebook      string    `db:"facebook"`
		Linkedin      string    `db:"linkedin"`
		Interest      string    `db:"interest"`
		LastLoginTime int64     `db:"last_login_time"`
		CreateTime    time.Time `db:"create_time"`
		UpdateTime    time.Time `db:"update_time"`
	}
```