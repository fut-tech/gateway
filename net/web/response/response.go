package response

import "github.com/gin-gonic/gin"

const (
	UnauthCode     = 401
	ForbiddenCode  = 403
	NotFoundCode   = 404
	BadGatewayCode = 502
	OKCode         = 200
)

// Response Response struct
type Response struct {
}

// OKResponse Return success response
func (r *Response) OKResponse(ctx *gin.Context, data interface{}) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(OKCode, data)
}

// FailResponse Return fail response
func (r *Response) FailResponse(ctx *gin.Context, code int, err error) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(code, gin.H{
		"error": err.Error(),
	})
}
