package handler

import (
	"net/http"
	"strconv"

	"my_gram/internal/model"
	"my_gram/internal/service"
	"my_gram/pkg"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserHandler interface {
	UserRegister(ctx *gin.Context)
	UserLogin(ctx *gin.Context)
	GetUsersById(ctx *gin.Context)
	// DeleteUsersById(ctx *gin.Context)
}

type userHandlerImpl struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) UserHandler {
	return &userHandlerImpl{
		svc: svc,
	}
}

// UserRegister godoc
//
//	@Summary		Register new user
//	@Description	will register new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			model.UserSignUp body model.UserSignUp{} true "register user"
//	@Success		201	{object}	model.UserSignUpCreate
//	@Failure		400	{object}	pkg.ErrorResponse
//	@Failure		404	{object}	pkg.ErrorResponse
//	@Failure		500	{object}	pkg.ErrorResponse
//	@Router			/users/register [post]
func (u *userHandlerImpl) UserRegister(ctx *gin.Context) {
	userSignUp := model.UserSignUp{}

	if err := ctx.Bind(&userSignUp); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := u.svc.UserRegister(ctx, userSignUp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	userResponse := model.UserSignUpCreate{
		DefaultColumn: user.DefaultColumn,
		Username:      user.Username,
		Email:         user.Email,
		DoB:           user.DoB,
	}
	ctx.JSON(http.StatusCreated, userResponse)
}

// UserLogin godoc
//
//	@Summary		User Login
//	@Description	User Login
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			model.UserSignIn body model.UserSignIn{} true "login user"
//	@Success		200	{object}	model.UserSignUpCreate
//	@Failure		400	{object}	pkg.ErrorResponse
//	@Failure		404	{object}	pkg.ErrorResponse
//	@Failure		500	{object}	pkg.ErrorResponse
//	@Router			/users/login [post]
func (u *userHandlerImpl) UserLogin(ctx *gin.Context) {
	userSignIn := model.UserSignIn{}

	if err := ctx.Bind(&userSignIn); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	token, err := u.svc.UserLogin(ctx, userSignIn)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, model.SignInOutput{
		Token: token,
	})
}

// ShowUsersById godoc
//
//	@Summary		Show users detail
//	@Description	will fetch 3rd party server to get users data to get detail user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	model.User
//	@Failure		400	{object}	pkg.ErrorResponse
//	@Failure		404	{object}	pkg.ErrorResponse
//	@Failure		500	{object}	pkg.ErrorResponse
//	@Router			/users/{id} [get]
func (u *userHandlerImpl) GetUsersById(ctx *gin.Context) {
	// get id user
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	user, err := u.svc.GetUsersById(ctx, uint64(id))
	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// UpdateUserById godoc
//
//	@Summary		Update User
//	@Description	Update User
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//
// @Param model.User formData model.UserCreateInputSwagger true "update user"
//
//	@Success		200	{object}	model.User
//	@Failure		400	{object}	pkg.ErrorResponse
//	@Failure		404	{object}	pkg.ErrorResponse
//	@Failure		500	{object}	pkg.ErrorResponse
//	@Router			/users/{id} [put]
func (u *userHandlerImpl) UpdateUserById(ctx *gin.Context) {
	// get id user
	id, err := strconv.Atoi(ctx.Param("id"))
	userUpdate := model.UserUpdate

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	userUpdate.ID = uint64(Id)

	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}

	userFileHeader, _ := ctx.FormFile("user")
	user, err := u.svc.UserUpdate(userUpdate, userFileHeader)
	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "user not found"})
		return
	}

	updateResponse := model.UserUpdateOutput{
		DefaultColumn: user.DefaultColumn,
		Username:      user.Username,
		Email:         user.Email,
		DoB:           user.DoB,
	}
	ctx.JSON(http.StatusOK, updateResponse)
}
