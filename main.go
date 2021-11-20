package main

import (
	"bytes"
	"context"
	"embed"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/goodsign/monday"
	"text/template"
	"time"
)

var (
	//go:embed resources
	res   embed.FS
	pages = map[string]string{
		"/": "resources/views/welcome.gohtml",
	}
)

func HandleRequest(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var err error
	var body bytes.Buffer

	page, ok := pages["/"]
	if !ok {
		panic(err)
	}

	tpl, err := template.ParseFS(res, page)
	if err != nil {
		panic(err)
	}

	loc, err := time.LoadLocation("Europe/Moscow")
	timeString := monday.Format(time.Now().In(loc), time.RFC1123Z, monday.LocaleRuRU)

	data := map[string]interface{}{
		"email": "dekmabot@gmail.com",
		"time":  timeString,
	}

	if err := tpl.Execute(&body, data); err != nil {
		panic(err)
	}

	response := events.APIGatewayProxyResponse{
		Body: body.String(),
		Headers: map[string]string{
			"Content-Type": "text/html; charset=utf-8",
		},
		StatusCode: 200,
	}

	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}
