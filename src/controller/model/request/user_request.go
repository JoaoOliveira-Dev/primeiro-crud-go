package request

type UserRequest struct {
	Name     string `json:"name"`
	Usuario  string `json:"usuario"`
	Email    string `json:"email"`
	Password string `json:"password"`
}