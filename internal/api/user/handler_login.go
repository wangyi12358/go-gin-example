package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
	"go-gin-example/pkg/app"
	"go-gin-example/pkg/config"
	"go-gin-example/pkg/e"
	"time"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	Expire int    `json:"expire"`
}

func (h *handler) Login(c *gin.Context) {
	a := app.Gin{C: c}
	req := new(LoginRequest)
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		a.Error(e.PARAMETER_ERROR, err.Error())
		return
	}

	user, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		a.ErrorCode(e.LOGIN_FAILED)
		return
	}
	if user == nil {
		a.ErrorCode(e.ACCOUNT_ERROR)
		return
	}

	token, err := generateToken(user.ID)
	if err != nil {
		a.ErrorCode(e.GENERATE_TOKEN_ERROR)
		return
	}

	a.Ok(&LoginResponse{
		Token:  token,
		Expire: config.Config.Jwt.Expire,
	})
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
