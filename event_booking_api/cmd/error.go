package main

import "github.com/gin-gonic/gin"

func ErrorResponse(err error) gin.H {
	return gin.H{"status": false, "message": err}
}

func SuccessResponse(data interface{}) gin.H {
	return gin.H{"status": true, "data": data}
}
