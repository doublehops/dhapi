package responses

func GetSingleItemResponse(data interface{}) SingleItemResponse {
	return SingleItemResponse{
		Data: data,
	}
}

func GetMultiItemResponse(data interface{}, pagination PaginationType) MultiItemResponse {
	return MultiItemResponse{
		Data:           data,
		PaginationType: pagination,
	}
}

func GetValidationError(code int, errors []ValidationField) ValidationErrorResponse {
	return ValidationErrorResponse{
		Name:    "Validation failed",
		Message: "One or more validation errors occurred",
		Code:    code,
		Status:  "error",
		Errors:  errors,
	}
}
