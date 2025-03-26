package middlewares

import (
	"explore/common"
	"explore/router/response"

	"github.com/gin-gonic/gin"
)

// 用户校验
func Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从认证头取token
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			response.FailWithCode(c, "无法验证您的身份", response.LoginErrorCode)
			c.Abort()
			return
		}
		claims, err := common.ParseToken(token)

		if err != nil {
			response.FailWithCode(c, "无效的token", response.LoginErrorCode)
			c.Abort()
			return
		}

		c.Set("session", common.Claims{
			UserId:   claims.UserId,
			UserName: claims.UserName,
		})
	}
}
