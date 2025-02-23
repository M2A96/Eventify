package domain

type Hello struct {
	Name    string
	Message string
}

type HelloRequest struct {
	Name string
}

type HelloResponse struct {
	Message string
}
