# Validation

This validation package is designed to make adding validation to API endpoints as clear and
easy to use as possible. Although common validation rules exist for use, it's easy to create
custom ones as well as custom error messages.

## Example Usage

You define the rules for each of the struct properties that need to be tested. Each
validation function should be placed comma-separated in the `Function` property of the
given rule.

```go
import "github.com/doublehops/dhapi/validator"

type Person struct {
    Name         string
    Age          string
    EmailAddress string
}

func main() {

    person := Person{
        Name:         "Jo",
        Age:          "Smith",
        EmailAddress: "jo@example.com",
    }

    rules := []Rule{
        {"Name", person.Name, false, []ValidationFunctions{validator.MinLength(13, "")}},
        {"Name", person.EmailAddress, false, []ValidationFunctions{validator.EmailAddress("My custom error message")}},
    }
    
    errors := RunValidation(rules)
    j, _ := json.Marshal(errors)
    fmt.Println(string(j))
}
```

The response will contain the errors as an array per property as multiple rules for
each could fail. It should be easy for any frontend to consume. An example would look like this:
```json
{
  "emailAddress": [
    "is not a valid email address"
  ],
  "name": [
    "is not the minimum length"
  ]
}
```

## Adding a Custom Validation Rule

