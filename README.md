# Each Golang Client
A Golang client for [Each](https://eachlabs.ai). Each is a platform for deploying and combining machine learning models as APIs.

> This library supports **only** Each Flow API.


## Installation
```bash
go get github.com/eachlabs/each-go
```

## Usage
```golang
package main

import (
	"context"
	"fmt"

	each "github.com/eachlabs/eachgo"
)

func main() {
	appCtx := context.Background()
	client, err := each.NewClient(each.WithCredential("YOUR_API_KEY"))
	if err != nil {
		panic(err)
	}
}
```

## Flow Methods

Each provides an AI workflow engine to orchestrate multiple models and data sources. You can create a chain flow multiple models and data sources together.

For more information, please refer to the [Each Flow API documentation](https://docs.eachlabs.ai/flows).


### Trigger a Flow
```golang
    flowID := "flow id" 
    inputs := map[string]interface{}{
        "input1": "value1",
        "input2": "value2",
    }
    triggerID, err := client.TriggerFlow(ctx, flowID, inputs)
    if err != nil {
        panic(err)
    }
    fmt.Println(triggerID)
```

### Get Execution 
```golang
    triggerID := "trigger id"
    flowID := "flow id"
    result, err := client.GetExecution(ctx, flowID, triggerID)
    if err != nil {
        panic(err)
    }
    fmt.Println(result)
```

### Get Executions
```golang
    flowID := "flow id"
    results, err := client.GetExecutions(ctx, flowID)
    if err != nil {
        panic(err)
    }
    fmt.Println(results)
```

## TODO
- [ ] Add tests
- [ ] Add more examples
- [ ] Add inference API support
- [ ] CI Integration