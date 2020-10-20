package router

import (
	com "breakfaster/controller/v1/common"
	cs "breakfaster/controller/v1/schema"
	exc "breakfaster/pkg/exception"
	ss "breakfaster/service/schema"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetOrder(c *gin.Context, order *ss.JSONOrder, err error) {
	switch err {
	case exc.ErrDateFormat:
		com.Response(c, http.StatusBadRequest, exc.ErrDateFormatCode, exc.ErrDateFormat)
		return
	case exc.ErrOrderNotFound:
		com.Response(c, http.StatusNotFound, exc.ErrOrderNotFoundCode, exc.ErrOrderNotFound)
	case exc.ErrGetOrder:
		com.Response(c, http.StatusInternalServerError, exc.ErrGetOrderCode, exc.ErrGetOrder)
		return
	case nil:
		c.JSON(http.StatusOK, order)
	default:
		com.Response(c, http.StatusInternalServerError, exc.ErrServerCode, exc.ErrServer)
	}
}

// GetOrder doc
// @Summary Get an order
// @Description Get an order by employee ID / access card number and date
// @Tags Order
// @Produce json
// @Param type query string true "Query type; should be 'eid' or 'card'"
// @Param payload query string true "Payload"
// @Param date query string true "Date in format YYYY-MM-DD"
// @Success 200 {object} schema.JSONOrder
// @Failure 400 {object} common.ErrResponse
// @Failure 404 {object} common.ErrResponse
// @Failure 500 {object} common.ErrResponse
// @Router /api/v1/order [get]
func (r *Router) GetOrder(c *gin.Context) {
	var params cs.FormGetOrder
	if err := c.ShouldBindQuery(&params); err != nil {
		com.Response(c, http.StatusBadRequest, exc.ErrInvalidParamCode, exc.ErrInvalidParam)
		return
	}
	switch params.Type {
	case "eid":
		order, err := r.orderSvc.GetOrderByEmpID(params.Payload, params.Date)
		handleGetOrder(c, order, err)
	case "card":
		order, err := r.orderSvc.GetOrderByAccessCardNumber(params.Payload, params.Date)
		handleGetOrder(c, order, err)
	}
}

// CreateOrders doc
// @Summary Create orders for an employee
// @Description Create orders of next week for an employee. Overwite an order if it already exists
// @Tags Order
// @Produce json
// @param X-Line-Identifer header string true "Line UID Authorization"
// @Param orders body schema.AllOrders true "Add Orders"
// @Success 200 {object} common.SuccessMessage
// @Failure 400 {object} common.ErrResponse
// @Failure 401 {object} common.ErrResponse
// @Failure 500 {object} common.ErrResponse
// @Router /api/v1/orders [post]
func (r *Router) CreateOrders(c *gin.Context) {
	var orders ss.AllOrders
	if err := c.ShouldBindJSON(&orders); err != nil {
		com.Response(c, http.StatusBadRequest, exc.ErrInvalidParamCode, exc.ErrInvalidParam)
		return
	}
	if len(orders.Foods) == 0 {
		com.Response(c, http.StatusBadRequest, exc.ErrInvalidParamCode, exc.ErrInvalidParam)
		return
	}
	err := r.orderSvc.CreateOrders(&orders)
	switch err {
	case exc.ErrDateFormat:
		com.Response(c, http.StatusBadRequest, exc.ErrDateFormatCode, exc.ErrDateFormat)
		return
	case exc.ErrCreateOrder:
		com.Response(c, http.StatusInternalServerError, exc.ErrCreateOrderCode, exc.ErrCreateOrder)
		return
	case exc.ErrGetlineUIDWhenConfirm:
		com.Response(c, http.StatusInternalServerError, exc.ErrGetlineUIDWhenConfirmCode, exc.ErrGetlineUIDWhenConfirm)
		return
	case exc.ErrGetOrderWhenConfirm:
		com.Response(c, http.StatusInternalServerError, exc.ErrGetOrderWhenConfirmCode, exc.ErrGetOrderWhenConfirm)
		return
	case exc.ErrSendMsg:
		com.Response(c, http.StatusInternalServerError, exc.ErrSendMsgCode, exc.ErrSendMsg)
		return
	case nil:
		c.JSON(http.StatusOK, com.OkMsg)
	default:
		com.Response(c, http.StatusInternalServerError, exc.ErrServerCode, exc.ErrServer)
	}
}

// SetPick doc
// @Summary Pick an order
// @Description Pick an order, setting the picked status true
// @Tags Order
// @Produce json
// @Param order body schema.PutPickOrder true "Pick Order"
// @Success 200 {object} common.SuccessMessage
// @Failure 400 {object} common.ErrResponse
// @Failure 500 {object} common.ErrResponse
// @Router /api/v1/order/pick [put]
func (r *Router) SetPick(c *gin.Context) {
	var pickOrder cs.PutPickOrder
	if err := c.ShouldBindJSON(&pickOrder); err != nil {
		log.Print(err)
		com.Response(c, http.StatusBadRequest, exc.ErrInvalidParamCode, exc.ErrInvalidParam)
		return
	}
	err := r.orderSvc.SetPick(pickOrder.EmpID, pickOrder.Date)
	switch err {
	case exc.ErrDateFormat:
		com.Response(c, http.StatusBadRequest, exc.ErrDateFormatCode, exc.ErrDateFormat)
		return
	case exc.ErrUpdateOrderStatus:
		com.Response(c, http.StatusInternalServerError, exc.ErrUpdateOrderStatusCode, exc.ErrUpdateOrderStatus)
		return
	case nil:
		c.JSON(http.StatusOK, com.OkMsg)
	default:
		com.Response(c, http.StatusInternalServerError, exc.ErrServerCode, exc.ErrServer)
	}
}
