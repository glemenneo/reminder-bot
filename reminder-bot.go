package main

import (
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/spf13/viper"
)

func getEnvVariable(key string) string {
  viper.SetConfigFile(".env")
  err := viper.ReadInConfig()

  if err != nil {
    log.Fatalf("Error while reading config file %s", err)
  }
	
  value, ok := viper.Get(key).(string)

  if !ok {
    log.Fatalf("Invalid type assertion")
  }

  return value
}

type ReminderBotStackProps struct {
	awscdk.StackProps
}

func NewReminderBotStack(scope constructs.Construct, id string, props *ReminderBotStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	lambdaHandler := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("myGoHandler"), &awscdklambdagoalpha.GoFunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2(),
		Entry:   jsii.String("./lambda-handler"),
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			GoBuildFlags: jsii.Strings(`-ldflags "-s -w"`),
		},
	})

	awsapigateway.NewLambdaRestApi(stack, jsii.String("lambdaApi"), &awsapigateway.LambdaRestApiProps{
		Handler: lambdaHandler,
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewReminderBotStack(app, "ReminderBotStack", &ReminderBotStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
	 Account: jsii.String(getEnvVariable("AWS_ACCOUNT_ID")),
	 Region:  jsii.String("ap-southeast-1"),
	}
}
