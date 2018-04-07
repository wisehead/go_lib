package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// User is a struct
type User struct {
	Name      string
	IsAdmin   bool
	Followers uint
}

// Month is a struct
type Month struct {
	MonthNumber int
	YearNumber  int
}

// MarshalJSON is self-implemented
func (m Month) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d/%d", m.MonthNumber, m.YearNumber)), nil
}

// UnmarshalJSON is self-implemented
func (m *Month) UnmarshalJSON(value []byte) error {
	parts := strings.Split(string(value), "/")
	//var err error
	//m.MonthNumber = strconv.ParseInt(parts[0], 10, 32)
	//m.YearNumber = strconv.ParseInt(parts[1], 10, 32)
	x1, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		fmt.Printf("xxx error: %v\n", err.Error())
	}
	fmt.Printf("x1:%v\n", x1)
	x2, err := strconv.ParseInt(parts[1], 10, 32)
	fmt.Printf("x2:%v\n", x2)
	//m.MonthNumber = int32(x1)
	//åm.YearNumber = x2
	return nil
}

func main() {
	//1.Marshal sample
	user := User{
		Name:      "cizixs",
		IsAdmin:   true,
		Followers: 36,
	}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("xxx error: %v\n", err.Error())
	}
	fmt.Printf("xxx %s\n", data)

	//2.Unmarshal sample
	data = []byte(`{"Name":"gopher","IsAdmin":false,"Followers":8900}`)
	var newUser = new(User)
	err = json.Unmarshal(data, &newUser)
	if err != nil {
		//fmt.Errorf("Can not decode data: %v\n", err)
		fmt.Printf("xxx error: %v\n", err.Error())
	}
	fmt.Printf("%v\n", newUser)

	//3.解析动态内容: interface{}
	data = []byte(`{"Name":"cizixs","IsAdmin":true,"Followers":36}`)

	var f interface{}
	json.Unmarshal(data, &f)
	name := f.(map[string]interface{})["Name"].(string)
	fmt.Printf("f.name:%v\n", name)

	//4. another way for:interface{}
	var f2 map[string]interface{}
	json.Unmarshal(data, &f2)
	// 省去了上面 f 的 type assertion 步骤
	name = f2["Name"].(string)
	fmt.Printf("f2.name:%v\n", name)
	followers := int(f2["Followers"].(float64))
	fmt.Printf("f2.followers:%v\n", followers)

	//5.延迟解析：json.RawMessage
	type BasicAuth struct {
		Email    string
		Password string
	}

	type TokenAuth struct {
		Token string
	}

	type User struct {
		Name      string
		IsAdmin   bool
		Followers uint
		Auth      json.RawMessage
	}
	data0 := []byte(`{"Name":"cizixs","IsAdmin":true,"Followers":36, "Auth":{"Email":"chenhui@baidu.com","Password":"xxxxxx"}}`)
	var user0 = new(User)
	err = json.Unmarshal(data0, &user0)
	fmt.Printf("%v\n", user0)

	data = user0.Auth
	fmt.Printf("user0.Auth:%v\n", data)
	//data = []byte(`{"Email":"chenhui@baidu.com","Password":"xxxxxx"}`)
	var basicAuth = new(BasicAuth)
	var tokenAuth = new(TokenAuth)
	err = json.Unmarshal(data, &basicAuth)
	if basicAuth.Email != "" {
		// 这是用户名/密码认证方式，在这里继续做一些处理
		fmt.Printf("basicAuth.Email:%v\n", basicAuth.Email)
	} else {
		json.Unmarshal(data, &tokenAuth)
		if tokenAuth.Token != "" {
			// 这是 token 认证方法
			fmt.Printf("tokenAuth.Token:%v\n", tokenAuth.Token)
		}
	}
	fmt.Printf("xxx %s\n", data)

	//6.自定义解析方法
	type Month struct {
		MonthNumber int
		YearNumber  int
	}

}
