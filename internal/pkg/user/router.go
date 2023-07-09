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
	g.GET("/:id", r.getUserByID)

}
func (r *router) createUser(c *gin.Context) {
	var req api.CreateUserRequest

	if err := c.ShouldBind(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if user, err := r.useCase.CreateUser(req); err != nil {
		_ = c.Error(err)
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
		c.JSON(500, toApiError(errHandler))
	} else {
		c.JSON(http.StatusOK, ToAPI(*user))
	}
}
