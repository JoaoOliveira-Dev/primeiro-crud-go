package response

type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Usuario  string `json:"usuario"`
	Email    string `json:"email"`
	Password string `json:"password"`
}