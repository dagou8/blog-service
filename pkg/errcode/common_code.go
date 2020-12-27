package errcode

var (
	Success                   = NewError(0, "Success")
	ServerError               = NewError(10000000, "Server internal error")
	InvalidParams             = NewError(10000001, "Input parameter error")
	NotFound                  = NewError(10000002, "Not found")
	UnauthorizedAuthNotExist  = NewError(10000003, "Authentication failed, the corresponding AppKey and AppSecret could not be found")
	UnauthorizedTokenError    = NewError(10000004, "Authentication failed, token error")
	UnauthorizedTokenTimeout  = NewError(10000005, "Authentication failed, token timeout")
	UnauthorizedTokenGenerate = NewError(10000006, "Authentication failed, token generation failed")
	TooManyRequests           = NewError(10000007, "Too many requests")
)
