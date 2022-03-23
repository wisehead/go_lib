#1.interface User

```go
	User interface {
		CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error)
		GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error)
		GetUsers(ctx context.Context, in *GetUsersReq, opts ...grpc.CallOption) (*GetUsersResp, error)
		GetBots(ctx context.Context, in *GetBotsReq, opts ...grpc.CallOption) (*GetBotsResp, error)
		UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
		DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserResp, error)
	}
```

#2.GetUser

```
GetUser
--c.cc.Invoke(ctx, "/user.User/GetUser", in, out, opts...)
--
```