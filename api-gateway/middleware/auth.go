package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"project/api-gateway/helper"
	pbAuth "project/auth-service/proto"
)

func (m *Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")
		if len(tokenValue) == 0 {
			helper.BadResponse(c, "invalid authorization header", http.StatusUnauthorized)
			c.Abort()
			return
		}

		authConn := helper.MustConnect(m.authServiceUrl)
		defer authConn.Close()

		authClient := pbAuth.NewAuthServiceClient(authConn)
		req := &pbAuth.ValidateTokenRequest{Token: tokenValue}
		res, err := authClient.ValidateToken(context.Background(), req)
		if err != nil {
			helper.BadResponse(c, fmt.Sprintf("server error : %s", err.Error()), http.StatusInternalServerError)
			c.Abort()
			return
		}

		if res.Email == "" {
			helper.BadResponse(c, "unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("email", res.Email)
		c.Next()
	}
}
