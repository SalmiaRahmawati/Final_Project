package middleware

import (
	"net/http"

	"my_gram/pkg"
	"my_gram/pkg/helper"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helper.VerifyToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, pkg.ErrorResponse{
				Message: "unauthorized",
				Errors:  []string{"invalid token"},
			})
			return
		}

		// store token claims in request data
		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
