package router

import (
	"my_gram/internal/handler"

	"github.com/gin-gonic/gin"
)

type UserRouter interface {
	Mount()
}

type userRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.UserHandler
}

func NewUserRouter(v *gin.RouterGroup, handler handler.UserHandler) UserRouter {
	return &userRouterImpl{v: v, handler: handler}
}

func (u *userRouterImpl) Mount() {
	u.v.POST("/register", u.handler.UserRegister)
	u.v.POST("/login", u.handler.UserLogin)
	u.v.GET("/:id", u.handler.GetUsersById)
	u.v.PUT("/:id", u.handler.PutUsersById)
	// u.v.DELETE("/:id", u.handler.DeleteUsersById)
}
