package middleware

import (
	com "breakfaster/controller/v1/common"
	exc "breakfaster/pkg/exception"
	"breakfaster/repository/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LineUIDAuthHeader is the auth header containing line UID
const LineUIDAuthHeader string = "X-Line-Identifer"

// LineUIDAuth authorize a request by X-Line-Identifer
func (auth *AuthChecker) LineUIDAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		lineUID := c.GetHeader(LineUIDAuthHeader)
		if lineUID == "" {
			com.Response(c, http.StatusUnauthorized, exc.ErrUnautorizedCode, exc.ErrUnautorized)
			c.Abort()
			return
		}
		_, err := auth.empRepository.GetEmpID(lineUID)
		switch err {
		case exc.ErrEmployeeNotFound:
			com.Response(c, http.StatusUnauthorized, exc.ErrUnautorizedCode, exc.ErrUnautorized)
			c.Abort()
			return
		case exc.ErrGetEmployee:
			com.Response(c, http.StatusInternalServerError, exc.ErrServerCode, exc.ErrServer)
			c.Abort()
			return
		case nil:
			c.Next()
		}
	}
}

// AuthChecker is the authorization middleware type
type AuthChecker struct {
	empRepository dao.EmployeeRepository
}

// NewAuthChecker is the factory for AuthChecker instance
func NewAuthChecker(empRepository dao.EmployeeRepository) *AuthChecker {
	return &AuthChecker{
		empRepository: empRepository,
	}
}
