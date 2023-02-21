package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpHeaderKeys() []string {
	return []string{"User-Id", "User-Name", "User-Device-Details", "User-Device-Type", "Content-Type"}
}
func HttpHeaderValues() []string {
	return []string{"application/json"}
}
func ValidateHttpHeaderKeys(header http.Header) bool {
	fmt.Printf("header: %v", header)
	for _, key := range HttpHeaderKeys() {
		if _, ok := header[key]; !ok {
			return false
		}
	}
	return true
}
func ValidateHttpHeaderValues(header http.Header) bool {
	for _, key := range HttpHeaderKeys() {
		if header.Get(key) == "" {
			return false
		}
	}
	return true
}
func ValidateHttpHeader(header http.Header) bool {
	return ValidateHttpHeaderKeys(header) && ValidateHttpHeaderValues(header)
}
func ValidateHttpHeaderDeviceType(header http.Header) bool {
	return header.Get("user-device-type") == "web" || header.Get("user-device-type") == "android" || header.Get("user-device-type") == "ios"
}
func ValidateHttpHeaderDeviceDetails(header http.Header) bool {
	return header.Get("user-device-details") != ""
}
func ValidateHttpHeaderUserId(header http.Header) bool {
	return header.Get("user-id") != ""
}
func ValidateHttpHeaderUserName(header http.Header) bool {
	return header.Get("user-name") != ""
}
func SetMiddlewareJSON(c *gin.Context) {
	if !ValidateHttpHeader(c.Request.Header) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Header"})
		c.Abort()
		return
	}
	if !ValidateHttpHeaderDeviceType(c.Request.Header) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Device Type"})
		c.Abort()
		return
	}
	if !ValidateHttpHeaderDeviceDetails(c.Request.Header) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Device Details"})
		c.Abort()
		return
	}
	if !ValidateHttpHeaderUserId(c.Request.Header) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User Id"})
		c.Abort()
		return
	}
	if !ValidateHttpHeaderUserName(c.Request.Header) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User Name"})
		c.Abort()
		return
	}
	c.Header("Content-Type", "application/json")
	c.Next()
}
func SetMiddlewareFormHeader(c *gin.Context) {
	if !ValidateHttpHeader(c.Request.Header) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Header"})
		c.Abort()
		return
	}
	if !ValidateHttpHeaderDeviceType(c.Request.Header) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Device Type"})
		c.Abort()
		return
	}
	if !ValidateHttpHeaderDeviceDetails(c.Request.Header) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Device Details"})
		c.Abort()
		return
	}
	if !ValidateHttpHeaderUserId(c.Request.Header) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User Id"})
		c.Abort()
		return
	}
	if !ValidateHttpHeaderUserName(c.Request.Header) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User Name"})
		c.Abort()
		return
	}
	c.Header("Content-Type", "multipart/form-data")
	c.Next()
}
