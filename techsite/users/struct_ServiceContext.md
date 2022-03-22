#1.ServiceContext

```go
type ServiceContext struct {
	Config               config.Config
	UserModel            model.UserModel
	UserProfileModel     model.UserProfileModel
	RecaptchaMiddleware  rest.Middleware
	SessionMiddleware    rest.Middleware
	CheckLoginMiddleware rest.Middleware
	UserRpc              userclient.User
	U2Rpc                u2client.U2
	Redis                *redis.Redis
}
```