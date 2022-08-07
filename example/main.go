package main

import (
	"fmt"
	"net/url"

	"github.com/deitrix/go-form"
	"github.com/k0kubun/pp/v3"
)

type Request struct {
	UserID string
	Age    int8
	Factor float64
}

func main() {
	var req Request
	fields := form.Fields{
		"userId": form.String(&req.UserID),
		"age":    form.Int(&req.Age),
		"factor": form.Float(&req.Factor),
	}
	values := url.Values{
		"userId": []string{"1234"},
		"age":    []string{"24"},
		"factor": []string{"3.14159"},
	}
	if err := form.Decode(fields, values); err != nil {
		panic(err)
	}
	pp.Println(req)
	fmt.Println(form.Encode(fields).Encode())
}
