package response

type LoginRes struct {
	Status           string `json:"status"`
	Type             string `json:"type"`
	CurrentAuthority string `json:"currentAuthority"`
}

type CurrentRes struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	UserID string `json:"userid"`
	Email  string `json:"email"`
}
