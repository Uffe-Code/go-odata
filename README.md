## Go OData [![GoDoc](https://godoc.org/github.com/Uffe-Code/go-odata?status.svg)](https://godoc.org/github.com/Uffe-Code/go-odata)

```go
import "github.com/Uffe-Code/go-odata"
```

go-odata is a library that desires to generate a fully functional OData API Client for Go.

### Setup
To setup the library, add a go file to be executed on it's own. It should look something
like this.

```go
package main

import (
	"fmt"
	"github.com/Uffe-Code/go-odata/modelGenerator"
	"path/filepath"
)

func main() {
	directoryPath, err := filepath.Abs("dataModel/")
	if err != nil {
		fmt.Printf("error while looking up folder: %s", err.Error())
		return
	}
	generator := modelGenerator.Generator{
		ApiUrl:        "https://services.odata.org/TripPinRESTierService/(S(c0y0kjlx4yjoxry4otnmoxf4))/",
		DirectoryPath: directoryPath,
	}
	err = generator.GenerateCode()
	if err != nil {
		fmt.Printf("error while generating code: %s", err.Error())
		return
	}

	fmt.Printf("code generated successfully")
}
```

When executed, it will generate a file in the specified directory with all model definitions.
You can then use the client.

### Initialize the client
```go
client := odataClient.New("https://services.odata.org/TripPinRESTierService/(S(c0y0kjlx4yjoxry4otnmoxf4))/")
```

#### Client wrapper
If you want to implement your own logic around the client, implement the Wrapper interface by adding a 
method called "ODataClient()" that will return the raw client. If you do, then the wrapper can be used
to fetch data sets.

### Get data set for model
```go
client := odataClient.New("https://services.odata.org/TripPinRESTierService/(S(c0y0kjlx4yjoxry4otnmoxf4))/")
peopleCollection := dataModel.NewPeopleCollection(client)
dataSet := peopleCollection.DataSet()
```

### List data from API
```go
data, errs := dataSet.List(odataClient.ODataFilter{
	filter: "Name eq 'Foo'"
})
for err := range errs {
    fmt.Printf("%s", err.Error())
    return
}
for model := range data {
    fmt.Println(model.Name)
}
```

### Create new record
```go
person := dataModel.Person{
	Name: "Foo",
}
insertedPerson, err := dataSet.Insert(person)
fmt.Printf("%d", insertedPerson.PersonId)
```

### Update a record
```go
person := dataModel.Person{
	Name: "Foo",
}
updatedPerson, err := dataSet.Update(5, person)
fmt.Printf("%d", updatedPerson.PersonId) // 5
```
