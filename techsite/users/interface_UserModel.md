#1.interface UserModel

```go
	UserModel interface {
		Insert(data User) (sql.Result, error)
		Find(ids []int64) ([]*User, error)
		FindOne(id int64) (*User, error)
		FindOneByEmail(email string) (*User, error)
		FindOneByUsername(username string) (*User, error)
		FindAll(from, size int64) (int64, []*User, error)
		FindBots() ([]int64, error)
		Update(data User) error
		UpdatePartial(data map[string]interface{}) error
		Delete(id int64) error
	}
```