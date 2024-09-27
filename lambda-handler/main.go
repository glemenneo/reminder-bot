package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	ginLambda = ginadapter.New(r)
}

func init() {
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	lambda.Start(handleRequest)
}