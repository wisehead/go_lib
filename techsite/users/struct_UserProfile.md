#1.struct UserProfile

```go
type UserProfile struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserType  int       `json:"user_type"`
	Company   string    `json:"company"`
	Industry  StringMap `json:"industry" gorm:"type:json"`
}
```