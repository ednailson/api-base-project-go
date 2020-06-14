package handlers

import "github.com/ednailson/httping-go"

func ExampleHandler(request httping.HttpRequest) httping.IResponse {
	return httping.OK("OK!")
}
