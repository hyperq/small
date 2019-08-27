package member

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

// Account godoc
// @Description 获取用户信息
// @Tags Member
// @Summary 用户信息
// @Accept  json
// @Produce  json
// @Param id query int true "Account ID"
// @Success 200 {object} model.Member "用户信息"
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /member/account?id={id} [get]
func (s *Member) Account(c *gin.Context) {
	id := c.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return

	}
	account, err := s.d.QueryMmeberByID(aid)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, account)
}
