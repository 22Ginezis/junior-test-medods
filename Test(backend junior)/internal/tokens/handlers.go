package tokens

import (
	"fmt"

	"github.com/NetworkPy/TestTaskJuniorBackDev/internal/handlers"
	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	T *Tokens
}

func NewHandler(T *Tokens) handlers.Handler {
	return &TokenHandler{T: T}
}

func (th *TokenHandler) Register(router *gin.Engine) {
	router.GET("/login/:guid", th.GenerateToken)
	router.GET("/refresh/:guid/:refresh", th.RefreshMiddleware(th.GenerateToken))
}

func (th *TokenHandler) GenerateToken(c *gin.Context) {
	guid := c.Param("guid")
	at, u, err := th.T.generateTokens(guid)
	if err != nil {
		respondWithError(c, 500, "Server Error")
		return
	}
	err = th.T.createDb(u, at)
	fmt.Println(err)
	if err != nil {
		respondWithError(c, 500, "Server Error")
		return
	}
	c.JSON(200, at)
}
