awsifaces
===========

A go module containing generated interfaces and mocks for services in [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2).

### Installation

```sh
go get github.com/jbreindel/awsifaces
```

While go will eliminate any unused code, you can simply copy any of these interfaces into your project if you don't want additional dependencies.

### Usage

```go
import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	awsifaces "github.com/jbreindel/awsifaces/service"
)

func NewSfn(cfg aws.Config) awsifaces.SfnClient {
	return sfn.NewFromConfig(cfg)
}
```

```go
mockSqsClient := new(mocks.SqsClient)
mockSqsClient.On("SendMessage", mock.Anything, mock.Anything, mock.Anything).Return(&sqs.SendMessageOutput{
  ...
}, nil).Once()
```

for additional usage options see [mockery](https://github.com/vektra/mockery).