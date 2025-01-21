package ecode

var (
	NotFound       = BadRequest(404, "Not Found")
	TooManyRequest = BadRequest(429, "Too many request")
	InternalErr    = BadRequest(500, "Internal Server Error")
	BadGatewayErr  = BadRequest(502, "Bad Gateway Error")
	InvalidParams  = NewInvalidParamsErr("Invalid Params")
	InvalidHashID  = BadRequest(601, "invalid hash id")
)

func NewInvalidParamsErr(msg string) *Error {
	return BadRequest(1000, msg)
}
