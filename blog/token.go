package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var Jwtkey =[]byte("my first jwt")
type Myclaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//生成token
func SetToken(username string)(string,error){
	expireTime:=time.Now().Add(10*time.Hour)//生成10个小时后失效
	setClamis:=Myclaims{
		Username:       username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "myblog",
		},
	}
	requestclaim:=jwt.NewWithClaims(jwt.SigningMethodHS256,setClamis)//加密
	token,err:=requestclaim.SignedString(Jwtkey)
	if err!=nil{
		return "",err
	}
	return token, err
}
//验证token
func CheckToken(token string)(*Myclaims,error){
	setToken,_:=jwt.ParseWithClaims(token,&Myclaims{}, func(token *jwt.Token) (interface{}, error) {
		return Jwtkey,nil
	})
	if key,_:=setToken.Claims.(*Myclaims);setToken.Valid{
		return key,errors.New("token验证失败")
	}
	return nil,nil
}

//jet中间件
func Jwttoken()gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenHeader:=c.Request.Header.Get("Authorization")
		if tokenHeader=="" {
			fmt.Println("token不存在")
			c.Abort()
			return
		}
		checkToken:=strings.SplitN(tokenHeader," ",2)
		if len(checkToken)!=2&&checkToken[0]!="Bearer"{
			fmt.Println("格式错误")
			c.Abort()
			return
		}
		key,err:=CheckToken(checkToken[1])
		if err!=nil{
			fmt.Println(err)
			c.Abort()
			return
		}
		if time.Now().Unix()>key.ExpiresAt{
			fmt.Println("token过期")
			c.Abort()
			return
		}
		fmt.Println("验证通过")
		c.Set("username",key.Username)
		c.Next()
	}
}