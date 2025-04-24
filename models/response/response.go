package response

type Response struct{
	Status int `json:"status"`
	Message string `json:"message"`
	Data any `json:"data"`
	Error any `json:"error"`
}