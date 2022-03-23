#1.service User

```
service User {
    rpc CreateUser(CreateUserReq) returns(CreateUserResp);
    rpc GetUser(GetUserReq) returns(GetUserResp);
    rpc GetUsers(GetUsersReq) returns(GetUsersResp);
    rpc GetBots(GetBotsReq) returns(GetBotsResp);
    rpc UpdateUser(UpdateUserReq) returns(UpdateUserResp);
    rpc DeleteUser(DeleteUserReq) returns(DeleteUserResp);
}

```