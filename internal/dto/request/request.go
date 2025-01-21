package request

type CreateUserRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
