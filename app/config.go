package app

type Environment struct {
	AwsID     string `yaml:"AwsId"`
	AwsSecret string `yaml:"AwsSecret"`
	AccountID int    `yaml:"accountId"`
	Region    string `yaml:"region"`
	SnsURL    string `yaml:"snsUrl"`
	SqsURL    string `yaml:"sqsUrl"`
	Build     []struct {
		Queues []struct {
			Name string `yaml:"Name"`
		} `yaml:"Queues"`
		Topics []struct {
			Name          string `yaml:"Name"`
			Subscriptions []struct {
				QueueName string `yaml:"QueueName"`
				Raw       bool   `yaml:"Raw"`
			} `yaml:"Subscriptions"`
		} `yaml:"Topics"`
	} `yaml:"Build"`
}

var Environments map[string]Environment

type EnvironmentService interface {
	List() string            // Output = list of topics, queues and Subscriptions
	VerifyEnv() string       // Response is a report with discrepancies
	VerifyAndBuildEnv() bool // true = success, false = failure (Errors should be output separately
}
