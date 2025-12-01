package session

type BodyRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}