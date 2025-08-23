package auth

type ConnectRequest struct {
	Phone string `json:"phone" binding:"required"`
}

type VerifyRequest struct {
	SessionId string `json:"sessionId" binding:"required"`
	Code      int    `json:"code" binding:"required"`
}

type ConnectResponse struct {
	SessionId string `json:"sessionId"`
}

type VerifyResponse struct {
	Token string `json:"token"`
}
