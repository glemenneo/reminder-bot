# Reminder Bot

This is a simple project that I'm using to learn about the AWS CDK.

It is written in fully in Go using the experimental AWS CDK Lambda Golang Library.

The goal is to write a Telegram Bot will send out reminders according to a predefined schedule.

## Requirements

A .env file in the root of the project with the follow enviroment variables:

- `AWS_ACCOUNT_ID` account id of the AWS account that you would like to deploy this application to

## Useful commands

- `cdk deploy` deploy this stack to your default AWS account/region
- `cdk diff` compare deployed stack with current state
- `cdk synth` emits the synthesized CloudFormation template
- `go test` run unit tests
