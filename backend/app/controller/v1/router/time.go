package router

import (
	"breakfaster/pkg/ordertime/schema"
	"breakfaster/service/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNextWeekInterval doc
// @Summary Get next week date interval
// @Description Get the starting and ending date of next week (in local time)
// @Tags Time
// @Produce json
// @Success 200 {object} schema.JSONTimeInterval
// @Failure 500 {object} common.ErrResponse
// @Router /api/v1/next-week [get]
func (r *Router) GetNextWeekInterval(c *gin.Context) {
	start, end := r.timer.GetNextWeekInterval()
	dateInterval := schema.JSONTimeInterval{
		StartDate: start.Format(constant.DateFormat),
		EndDate:   end.Format(constant.DateFormat),
	}
	c.JSON(http.StatusOK, dateInterval)
}
