package exception

import "errors"

var (
	// Gerneral errors
	ErrInvalidParam = errors.New("invalid parameter")
	ErrUnautorized  = errors.New("unauthorized")
	ErrInvalidSet   = errors.New("resource not found")
	ErrServer       = errors.New("server error")

	// Custom common erors
	ErrDateFormat = errors.New("wrong date format; should be YYYY-MM-DD")

	// Food errors
	ErrGetFood      = errors.New("server error while getting food")
	ErrFoodNotFound = errors.New("food not found")

	// Employee errors
	ErrGetEmployee       = errors.New("server error while getting employee")
	ErrEmployeeNotFound  = errors.New("employee not found")
	ErrUpsertEmployeeIDs = errors.New("server error while upserting employee by IDs")

	// Order errors
	ErrGetOrder              = errors.New("server error while getting order")
	ErrOrderNotFound         = errors.New("order not found")
	ErrCreateOrder           = errors.New("server error while creating order")
	ErrDeleteOrder           = errors.New("server error while deleting order")
	ErrGetlineUIDWhenConfirm = errors.New("server error while getting line UID for confirmation")
	ErrGetOrderWhenConfirm   = errors.New("server error while confirming order")
	ErrUpdateOrderStatus     = errors.New("server error while updating order status")

	// Message errors
	ErrSendMsg = errors.New("server error while sending message")

	// Redis errors
	ErrRedisCmdNotFound = errors.New("redis command not foundl supports SET and DELETE")
)
