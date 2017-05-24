package aws

type AwsEnvironmentService struct {
	SNS_URL   string
	SQS_URL   string
	AwsId     string
	AwsSecret string
	AwsRegion string
	Build     []string
}

func (awsenv AwsEnvironmentService) List() string {
	return ""
}

func (awsenv AwsEnvironmentService) VerifyEnv() string {
	return ""
}

func (awsenv AwsEnvironmentService) VerifyAndBuildEnv() {

}
