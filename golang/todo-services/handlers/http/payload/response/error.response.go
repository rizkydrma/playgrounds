package response

type BaseResponse struct {
	Message		string	`json:"responseMessage"`
	Code			string	`json:"responseCode"`
	Data			any			`json:"data"`
}

const (
	SUCCESS_CODE 										= "200"
	ERROR_CODE											= "500"
	NOT_FOUND_CODE									= "404"
	EMAIL_ALREADY_EXIST_CODE				= "402"

	INTERNAL_SERVER_ERROR_CODE			= "500"
	INVALID_REQUEST_PAYLOAD_CODE		= "400"
	
	FAILED_GET_DATA_CODE						= "500"
	FAILED_STORE_DATA_CODE					= "500"

	UNAUTHORIZED_CODE								= "503"
)

const (
	SUCCESS_MESSAGE									= "Success"
	ERROR_MESSAGE										= "Error"
	NOT_FOUND_MESSAGE								= "Not Found"
	EMAIL_ALREADY_EXIST_MESSAGE			= "Email already exist"

	INTERNAL_SERVER_ERROR_MESSAGE		= "Internal server error"
	INVALID_REQUEST_PAYLOAD_MESSAGE	= "Invalid request payload"
	
	FAILED_GET_DATA_MESSAGE					= "Failed to get data"
	FAILED_STORE_DATA_MESSAGE				= "Failed to store data, bad request"

	UNAUTHORIZED_MSG								= "Unauthorized"
)