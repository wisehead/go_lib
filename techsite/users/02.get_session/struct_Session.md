#1.struct Session

```go
// Session stores the values and optional configuration for a session.
type Session struct {
	// The ID of the session, generated by stores. It should not be used for
	// user data.
	ID string
	// Values contains the user-data for the session.
	Values  map[interface{}]interface{}
	Options *Options
	IsNew   bool
	store   Store
	name    string
}
```