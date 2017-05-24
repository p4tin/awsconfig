package main

import (
	"io/ioutil"
	"log"

	"fmt"

	"os"

	"github.com/p4tin/awsconfig/app"
	"gopkg.in/yaml.v2"
)

func LoadFile(config interface{}, filename string) error {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Aws Configurator v1.0")
	var val app.Environment

	err := LoadFile(&val, "config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	val.AwsID = os.Getenv("AWS_ACCESS_KEY")
	val.AwsSecret = os.Getenv("AWS_SECRET_KEY")

	fmt.Println("ID: \t\t", val.AwsID)
	fmt.Println("Secret: \t", val.AwsSecret)
	fmt.Println("Account: \t", val.AccountID)
	fmt.Println("Region: \t", val.Region)
	fmt.Println("SNS Url: \t", val.SnsURL)
	fmt.Println("SQS Url: \t", val.SqsURL)
	fmt.Println("Build:\n\t Queues:")
	for _, q := range val.Build[0].Queues {
		fmt.Println("\t\t", q.Name)
	}
	fmt.Println("\t Topics:")
	for _, t := range val.Build[1].Topics {
		fmt.Println("\t\t", t.Name)
		fmt.Println("\t\t\t Subscriptions:")
		for _, s := range t.Subscriptions {
			fmt.Println("\t\t\t\t", s.QueueName, "(Raw = ", s.Raw, ")")
		}
	}
	fmt.Println("Done!!!")
}
