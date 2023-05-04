package user_api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
	"go-gin-example/internal/service/user_service"
	"go-gin-example/pkg/app"
	"go-gin-example/pkg/config"
	"time"
)

func Login(c *gin.Context) {
	a := app.Gin{C: c}
	var req LoginReq
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		a.Error(1100, err.Error())
		return
	}

	user, err := user_service.Login(req.Username, req.Password)
	if err != nil {
		a.Error(1101, "用户名或密码错误")
		return
	}

	token, err := generateToken(user.ID)
	if err != nil {
		a.Error(1102, "生成Token出错")
		return
	}

	a.Ok(&LoginRes{
		Token:  token,
		Expire: config.Config.Jwt.Expire,
	})
}

func Logout(c *gin.Context) {

}

func Profile(c *gin.Context) {
	a := app.Gin{C: c}
	userId, err := getUser(c)
	if err != nil {
		a.Response(200, 401, nil)
		return
	}
	user, err := user_service.FindById(userId)
	if err != nil {
		a.Error(1105, "用户不存在")
		return
	}
	a.Ok(user)
}

func generateToken(userId int64) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// 设置token的过期时间
	expireTime := time.Now().Add(time.Second * time.Duration(config.Config.Jwt.Expire)).Unix()
	claims["userId"] = userId
	claims["exp"] = expireTime

	// 使用密钥签署token
	tokenString, err := token.SignedString([]byte(config.Config.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getUser(c *gin.Context) (int64, error) {
	tokenString := c.GetHeader("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("my_secret_key"), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("无效的Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("无法获取claims")
	}

	userId, ok := claims["userId"].(int64)
	if !ok {
		return 0, errors.New("无法获取userId")
	}

	return userId, nil
}
