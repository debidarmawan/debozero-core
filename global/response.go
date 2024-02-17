package global

type Response[T any] struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

var ResponseStatus = struct {
	FailedResponse  string
	SuccessResponse string
	RetryResponse   string
}{
	FailedResponse:  "FAILED",
	SuccessResponse: "OK",
	RetryResponse:   "RETRY",
}
