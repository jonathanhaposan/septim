package webserver

type Writer interface {
	Success(interface{}) Response
	Failed(string) Response
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

func Failed(e string) Response {
	return Response{
		Error: e,
	}
}
