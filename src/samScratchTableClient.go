package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
	"log"
	"os"
)

var Info *log.Logger

type Servant struct {
	ServantId string `json:"ServantId"`
	Name      string `json:"Name"`
	Class     string `json:"Class"`
}

type Servants []Servant

func Init() {
	Info = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (servants Servants) toJson() string {
	output := "{\"Servant\":["
	for _, servant := range servants {
		output += servant.toJson()
	}
	output += "]}"
	return output
}

func (servant Servant) toJson() string {
	output := "{"

	output += "\"ServantId\":\"" + servant.ServantId + "\","
	output += "\"Name\":\"" + servant.Name + "\","
	output += "\"Class\":\"" + servant.Class + "\""
	output += "},"

	return output
}

func getStringFromItems(items []map[string]*dynamodb.AttributeValue) string {
	servants := Servants{}

	err := dynamodbattribute.UnmarshalListOfMaps(items, &servants)
	if err != nil {
		Info.Print("Got error unmarshalling item:", err)
		return ""
	}

	output := servants.toJson()
	return output
}

func GetAllServants() string {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// https://docs.aws.amazon.com/ja_jp/sdk-for-go/v1/developer-guide/configuring-sdk.html
	svc := dynamodb.New(
		sess,
		aws.NewConfig().WithRegion("ap-northeast-1"))

	input := &dynamodb.ScanInput{
		TableName: aws.String("Servant"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		Info.Print("Got error scanning table:", err)
		return ""
	}

	output := getStringFromItems(result.Items)

	return output
}

// For testing from the Local PC.
// Please change the name of this function as "main" when you want to test it.
func test() {
	Init()
	output := GetAllServants()
	fmt.Println(output)
}
