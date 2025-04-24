package response

type ResponseUser struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}