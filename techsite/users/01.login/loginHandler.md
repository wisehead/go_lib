#1.Login

```
loginHandler
--logic.NewLoginLogic(r.Context(), ctx)
--LoginLogic.Login
----validate := validator.New()
----err := validate.Struct(req)
----user, err := l.svcCtx.UserModel.FindOneByEmail(req.Email)
----phpass := util.NewPhpass(8)
----phpass.Verify([]byte(req.Password), []byte(user.Password))
----user_map := map[string]interface{}{
		"id":              user.Id,
		"last_login_time": time.Now().Unix(),
	}
----l.svcCtx.UserModel.UpdatePartial(user_map)
----return &types.LoginResp{
		UserId:        user.Id,
		Email:         user.Email,
		UserRole:      user.UserRole,
		Channel:       user.Channel,
		IsActivated:   user.IsActivated,
		IsBlocked:     user.IsBlocked,
		IsDeleted:     user.IsDeleted,
		IsBot:         user.IsBot,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Avatar:        user.Avatar,
		Username:      user.Username,
		AboutYou:      user.AboutYou,
		Location:      user.Location,
		Website:       user.Website,
		FullName:      user.FullName,
		Twitter:       user.Twitter,
		Facebook:      user.Facebook,
		Linkedin:      user.Linkedin,
		Interest:      user.Interest,
		LastLoginTime: user.LastLoginTime,
		CreatedAt:     user.CreateTime.Unix(),
	}	
--r.Context().Value(common.ContextKeySession).(*sessions.Session)
--session.Options.SameSite = http.SameSiteNoneMode
--session.Options.Secure = true
--session.Values["userid"] = resp.UserId
--session.Values["email"] = resp.Email
--Session.save
----s.store.Save(r, w, s)
```