package constants

var SuperuserForbiddenEndpoints = []string{
	"POST /api/v1/users/change-password",
	"POST /api/v1/orders/checkout",
	"PUT /api/v1/customers/update",
	"POST /api/v1/payments",
	"POST /api/v1/feedbacks",
	"PUT /api/v1/customers/update",
	"POST /api/v1/loyalties/redeems/transaction",
	"POST /api/v1/payments/{id}/cancel",
	"PUT /api/v1/customers/update-photo",
}
