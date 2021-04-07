package handlers

type healthcheck struct {
	Ping string `json:"ping"`
}

type helloWorld struct {
	Hello string `json:"hello"`
}

type notFound struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}
