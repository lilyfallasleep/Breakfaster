package router

import (
	com "breakfaster/controller/v1/common"
	cs "breakfaster/controller/v1/schema"
	exc "breakfaster/pkg/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetFoodAll godoc
// @Summary Get all foods
// @Description Retrieve foods for each day in the given time interval
// @Tags Food
// @Produce json
// @Param start query string true "Start date in format YYYY-MM-DD"
// @Param end query string true "End date in format YYYY-MM-DD"
// @Success 200 {array} schema.NestedFood
// @Failure 400 {object} common.ErrResponse
// @Failure 404 {object} common.ErrResponse
// @Failure 500 {object} common.ErrResponse
// @Router /api/v1/foods [get]
func (r *Router) GetFoodAll(c *gin.Context) {
	var foodDate cs.FormFoodDate
	if err := c.ShouldBindQuery(&foodDate); err != nil {
		com.Response(c, http.StatusBadRequest, exc.ErrDateFormatCode, exc.ErrDateFormat)
		return
	}
	startDate := foodDate.StartDate
	endDate := foodDate.EndDate
	if startDate == "" || endDate == "" {
		com.Response(c, http.StatusBadRequest, exc.ErrDateFormatCode, exc.ErrDateFormat)
		return
	}

	foods, err := r.foodSvc.GetFoodAll(startDate, endDate)
	switch err {
	case exc.ErrDateFormat:
		com.Response(c, http.StatusBadRequest, exc.ErrDateFormatCode, exc.ErrDateFormat)
		return
	case exc.ErrFoodNotFound:
		com.Response(c, http.StatusNotFound, exc.ErrFoodNotFoundCode, exc.ErrFoodNotFound)
		return
	case exc.ErrGetFood:
		com.Response(c, http.StatusInternalServerError, exc.ErrGetFoodCode, exc.ErrGetFood)
		return
	case nil:
		c.JSON(http.StatusOK, foods)
	default:
		com.Response(c, http.StatusInternalServerError, exc.ErrServerCode, exc.ErrServer)
	}
}
