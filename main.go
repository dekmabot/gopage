package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

func HandleRequest(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	response := "<!DOCTYPE html>\n" +
		"<html lang=\"en\">\n" +
		"<head>\n    " +
		"<meta charset=\"UTF-8\">\n    " +
		"<title>Domain is for sale!</title>\n" +
		"<style>html, body{height:100%; width:100%; background-color: #161b22; color: #ccc; font-size: 18px; padding: 0; margin: 0;} div.container {height: 100%; position: relative } div.container p {margin: 0; position: absolute; top: 50%; left: 50%; margin-right: -50%; transform: translate(-50%, -50%) }</style>\n" +
		"</head>\n" +
		"<body>\n" +
		"<div class=\"container\">" +
		"<p>" +
		"This domain is for sale<br>" +
		"<a href=\"mailto:dekmabot@gmail.com\" style=\"color:white; text-decoration: none; display: inline-block; margin: 10px 0;\">dekmabot@gmail.com</a><br>" +
		"<span style=\"font-size:0.8em; color:grey\">" + time.Now().UTC().String() + "</span>" +
		"</p>" +
		"</div>\n" +
		"</body>\n" +
		"</html>"

	return events.APIGatewayProxyResponse{
		Body: response,
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
