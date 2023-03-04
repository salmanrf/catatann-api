package models

type CustomHttpErrors struct {
	Messages []string
	Code int
}

func CreateCustomHttpError(code int, error_messages interface{}) (*CustomHttpErrors) {
	var cust_error CustomHttpErrors 
	
	cust_error.Code = 500
	cust_error.Messages = []string{"internal server error"}

	err_message, is_error := error_messages.(error) 
	
	if is_error {
		cust_error.Messages = []string{err_message.Error()}
	}

	message, is_string := error_messages.(string) 
	
	if is_string {
		cust_error.Messages = []string{message}
	}

	messages, is_array := error_messages.([]string)

	if is_array {
		cust_error.Messages = messages
	}

	if code != 0 {
		cust_error.Code = code
	}
	
	return &cust_error
}