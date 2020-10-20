package exception

import "net/http"

const (
	// General error code 4xx
	ErrInvalidParamCode = http.StatusBadRequest
	ErrUnautorizedCode  = http.StatusUnauthorized
	ErrInvalidSetCode   = http.StatusNotFound

	// General error code 5xx
	ErrServerCode = http.StatusInternalServerError

	// Custom common error code 1xxx
	ErrDateFormatCode = 1001

	// Food error code 2xxx
	ErrGetFoodCode      = 2001
	ErrFoodNotFoundCode = 2002

	// Employee error code 3xxx
	ErrGetEmployeeCode       = 3001
	ErrEmployeeNotFoundCode  = 3002
	ErrUpsertEmployeeIDsCode = 3003

	// Order error code 4xxx
	ErrGetOrderCode              = 4001
	ErrOrderNotFoundCode         = 4002
	ErrCreateOrderCode           = 4003
	ErrGetlineUIDWhenConfirmCode = 4004
	ErrGetOrderWhenConfirmCode   = 4005
	ErrUpdateOrderStatusCode     = 4006

	// Send message error code 5xxx
	ErrSendMsgCode = 5001
)
