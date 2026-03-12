package errs

import "net/http"

const (
	// Global Status Codes
	CodeOK               = 0    // Success
	CodeInternalError    = 1000 // Internal server error
	CodeUnauthorized     = 1001 // Unauthorized: No login or invalid token
	CodeForbidden        = 1002 // Forbidden: Insufficient permissions
	CodeTooManyRequests  = 1003 // Rate limit exceeded
	CodeInvalidSignature = 1004 // Invalid request signature
	CodeNotImplemented   = 1005 // Feature not implemented

	// User & Authentication Codes
	CodeUserNotFound      = 2000 // User does not exist
	CodeUserAlreadyExists = 2001 // User already exists
	CodeInvalidPassword   = 2002 // Incorrect password
	CodeUserDisabled      = 2003 // User account is disabled
	CodeTokenExpired      = 2004 // Authentication token expired

	// Resource Management Codes
	CodeResourceNotFound      = 3000 // Resource not found
	CodeResourceAlreadyExists = 3001 // Resource already exists
	CodeConflict              = 3002 // Resource conflict (e.g., version mismatch)
	CodeResourceLocked        = 3003 // Resource is locked or occupied
	CodeNoContent             = 3004 // Request processed but no content returned
	CodeDuplicate             = 3005

	// Request & Parameter Validation Codes
	CodeInvalidParam      = 4000 // General parameter error
	CodeMissingParam      = 4001 // Missing required parameter
	CodeValidationFailed  = 4002 // Parameter validation failed
	CodeUnsupportedFormat = 4003 // Unsupported data format
	CodeOutOfRange        = 4004 // Value or pagination out of range
	CodeCacheMiss         = 4005 // Cache miss

	// Infrastructure & System Codes
	CodeDatabaseError    = 5000 // Database operation failed
	CodeCacheError       = 5001 // Redis or cache failure
	CodeNetworkError     = 5002 // Network request failure
	CodeExternalAPIError = 5003 // External API service error
	CodeFileSystemError  = 5004 // File system operation error
	CodeStorageError     = 5005
)

func MapToHTTPStatus(code int) int {
	switch code {
	case CodeOK:
		return http.StatusOK
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodeTooManyRequests:
		return http.StatusTooManyRequests
	case CodeInvalidSignature, CodeInvalidParam, CodeMissingParam, CodeValidationFailed:
		return http.StatusBadRequest
	case CodeResourceNotFound, CodeUserNotFound:
		return http.StatusNotFound
	case CodeConflict, CodeResourceAlreadyExists, CodeCacheMiss, CodeDuplicate:
		return http.StatusConflict
	case CodeNotImplemented:
		return http.StatusNotImplemented
	case CodeDatabaseError, CodeCacheError, CodeNetworkError, CodeExternalAPIError, CodeFileSystemError, CodeStorageError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
