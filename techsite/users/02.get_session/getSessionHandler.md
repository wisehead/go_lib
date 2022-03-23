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
----
```