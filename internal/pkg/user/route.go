package user

import (
	"fxdemo/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Router interface {
	Handler(gGroup gin.IRouter)
}

type router struct {
	useCase UseCase
}

func NewRouter(useCase UseCase) Router {
	return &router{useCase: useCase}
}

func (r *router) Handler(gGroup gin.IRouter) {
	g := gGroup.Group("/users")
	g.POST("", r.createUser)
	g.PATCH("", r.updataByName)
	g.GET("/:id", r.getUserByID)
	g.GET("/name/:name", r.getUserByName)

}
func (r *router) createUser(c *gin.Context) {
	var req api.CreateUserRequest

	if err := c.ShouldBind(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if user, err := r.useCase.CreateUser(req); err != nil {
		c.JSON(translateErrorCode(err), err)
		return
	} else {
		c.JSON(http.StatusOK, ToAPI(*user))
	}

}

func (r *router) getUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if user, errHandler := r.useCase.GetUserByID(id); errHandler != nil {
		c.JSON(translateErrorCode(errHandler), toApiError(errHandler))
	} else {
		c.JSON(http.StatusOK, ToAPI(*user))
	}
}
func (r *router) getUserByName(c *gin.Context) {
	userName := c.Param("name")

	if user, errHandler := r.useCase.GetUserByName(userName); errHandler != nil {
		c.JSON(translateErrorCode(errHandler), toApiError(errHandler))
	} else {
		c.JSON(http.StatusOK, ToAPI(*user))
	}
}

func (r *router) updataByName(c *gin.Context) {
	var req api.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if user, err := r.useCase.updateByName(req); err != nil {
		c.JSON(translateErrorCode(err), toApiError(err))
	} else {
		c.JSON(http.StatusOK, ToAPI(*user))
	}
}
