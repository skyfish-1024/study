package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var Jwtkey =[]byte("myfirstjwt")
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
	key,_:=setToken.Claims.(*Myclaims);
	if !setToken.Valid{
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
			c.JSON(200,gin.H{"验证结果：":"验证失败！"})
			c.Abort()
			return
		}
		cToken:=strings.SplitN(tokenHeader," ",2)
		if len(cToken)!=1 && cToken[0]!="Bearer"{
			fmt.Println("ctoken:",cToken)
			fmt.Println("len:",len(cToken),"ctoken:",cToken[0])
			fmt.Println("格式错误")
			c.JSON(200,gin.H{"验证结果：":"验证失败！"})
			c.Abort()
			return
		}
		key,err:=CheckToken(cToken[0])
		if err!=nil{
			fmt.Println(err)
			c.JSON(200,gin.H{"验证结果：":"验证失败！"})
			c.Abort()
			return
		}
		if time.Now().Unix()>key.ExpiresAt{
			fmt.Println("token过期")
			c.JSON(200,gin.H{"验证结果：":"验证失败！"})
			c.Abort()
			return
		}
		fmt.Println("验证通过")
		c.Set("username",key.Username)
		c.Next()
	}
}