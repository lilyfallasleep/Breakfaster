package router

import (
	com "breakfaster/controller/v1/common"
	cs "breakfaster/controller/v1/schema"
	exc "breakfaster/pkg/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEmployeeByEmpID doc
// @Summary Get employee line UID
// @Description Get employee line UID by querying employee ID
// @Tags Employee
// @Produce json
// @Param emp-id query string true "Employee ID"
// @Success 200 {object} schema.JSONEmployee
// @Failure 400 {object} common.ErrResponse
// @Failure 404 {object} common.ErrResponse
// @Failure 500 {object} common.ErrResponse
// @Router /api/v1/employee/line-uid [get]
func (r *Router) GetEmployeeByEmpID(c *gin.Context) {
	var employeeEmpID cs.FormEmployeeEmpID
	if err := c.ShouldBindQuery(&employeeEmpID); err != nil {
		com.Response(c, http.StatusBadRequest, exc.ErrInvalidParamCode, exc.ErrInvalidParam)
		return
	}

	employee, err := r.empSvc.GetEmployeeByEmpID(employeeEmpID.EmpID)
	switch err {
	case exc.ErrEmployeeNotFound:
		com.Response(c, http.StatusNotFound, exc.ErrEmployeeNotFoundCode, exc.ErrEmployeeNotFound)
		return
	case exc.ErrGetEmployee:
		com.Response(c, http.StatusInternalServerError, exc.ErrGetEmployeeCode, exc.ErrGetEmployee)
		return
	case nil:
		c.JSON(http.StatusOK, employee)
	default:
		com.Response(c, http.StatusInternalServerError, exc.ErrServerCode, exc.ErrServer)
	}
}

// GetEmployeeByLineUID doc
// @Summary Get employee ID
// @Description Get employee ID by querying employee line UID
// @Tags Employee
// @Produce json
// @Param line-uid query string true "Employee Line UID"
// @Success 200 {object} schema.JSONEmployee
// @Failure 400 {object} common.ErrResponse
// @Failure 404 {object} common.ErrResponse
// @Failure 500 {object} common.ErrResponse
// @Router /api/v1/employee/emp-id [get]
func (r *Router) GetEmployeeByLineUID(c *gin.Context) {
	var employeeLineUID cs.FormEmployeeLineUID
	if err := c.ShouldBindQuery(&employeeLineUID); err != nil {
		com.Response(c, http.StatusBadRequest, exc.ErrInvalidParamCode, exc.ErrInvalidParam)
		return
	}

	employee, err := r.empSvc.GetEmployeeByLineUID(employeeLineUID.LineUID)
	switch err {
	case exc.ErrEmployeeNotFound:
		com.Response(c, http.StatusNotFound, exc.ErrEmployeeNotFoundCode, exc.ErrEmployeeNotFound)
		return
	case exc.ErrGetEmployee:
		com.Response(c, http.StatusInternalServerError, exc.ErrGetEmployeeCode, exc.ErrGetEmployee)
		return
	case nil:
		c.JSON(http.StatusOK, employee)
	default:
		com.Response(c, http.StatusInternalServerError, exc.ErrServerCode, exc.ErrServer)
	}
}

// UpsertEmployeeByIDs doc
// @Summary Insert an employee by employee ID and line UID
// @Description Insert an employee. If the employee ID or line UID exists, update the corresponding field
// @Tags Employee
// @Produce json
// @Param employee body schema.PostEmployee true "Add Employee"
// @Success 200 {object} common.SuccessMessage
// @Failure 400 {object} common.ErrResponse
// @Failure 500 {object} common.ErrResponse
// @Router /api/v1/employee [post]
func (r *Router) UpsertEmployeeByIDs(c *gin.Context) {
	var employee cs.PostEmployee
	if err := c.ShouldBindJSON(&employee); err != nil {
		com.Response(c, http.StatusBadRequest, exc.ErrInvalidParamCode, exc.ErrInvalidParam)
		return
	}
	err := r.empSvc.UpsertEmployeeByIDs(employee.EmpID, employee.LineUID)
	switch err {
	case exc.ErrUpsertEmployeeIDs:
		com.Response(c, http.StatusInternalServerError, exc.ErrUpsertEmployeeIDsCode, exc.ErrUpsertEmployeeIDs)
		return
	case nil:
		c.JSON(http.StatusOK, com.OkMsg)
	default:
		com.Response(c, http.StatusInternalServerError, exc.ErrServerCode, exc.ErrServer)
	}
}
