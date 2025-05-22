package types

type LoginPayload struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


type Response struct {
	Message string `json:"message"`
	Data any `json:"data"`
}

type SignupPayload struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}