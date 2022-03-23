#1.getSessionHandler

```go
getSessionHandler
--logic.NewGetSessionLogic(r.Context(), ctx)
--GetSessionLogic.GetSession
----session, ok := l.ctx.Value(common.ContextKeySession).(*sessions.Session)
----uid, find := session.Values["userid"]
----resp.UserId = uid.(int64)
----if resp.UserId <= 0
------return resp, nil
----get_req := &userclient.GetUserReq{Id: resp.UserId,}
----l.svcCtx.UserRpc.GetUser(l.ctx, get_req)
------client := user.NewUserClient(m.cli.Conn())
------client.GetUser(ctx, in, opts...)
--------c.cc.Invoke(ctx, "/user.User/GetUser", in, out, opts...)
----------UserServer.GetUser
------------GetUserLogic.GetUser
--------------if in.Id > 0
----------------l.svcCtx.UserModel.FindOne(in.Id)
------------------defaultUserModel.FindOne
--------------------select %s from %s where `id` = ? limit 1
--------------else if len(in.Email) > 0
----------------one, err = l.svcCtx.UserModel.FindOneByEmail(in.Email)
--------------else if len(in.Username) > 0
----------------one, err = l.svcCtx.UserModel.FindOneByUsername(in.Username)
--------------up, err := l.svcCtx.UserProfileModel.Find(uint(in.Id))
----------------UserProfileModel.Find
------------------ u.db.Find(&userProfile, "id = ?", id)
```