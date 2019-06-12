# emailer

[![GoDoc](https://godoc.org/github.com/fabritsius/emailer?status.svg)](https://godoc.org/github.com/fabritsius/emailer)

This Go module simplifies a process of sending HTML emails to multiple users.

## Example

Simple use case.

You have a `recipients.csv` file with user emails (imagine there are more lines):

```csv
ID, NAME, MAIL, PET
1,Anton,anton@example.com,cat
```

And you want to send an email to each one of them. Here is how you can do that:

```go
package main

import (
    "fmt"

    "github.com/fabritsius/emailer"
    "github.com/fabritsius/envar"
    "github.com/fabritsius/csvier"
)

func main() {
    // Get a simple email template
    temp := simpleTemplate()

    // Get a slice of recipients from a file
    recipients, err := csvier.ReadFile("recipients.csv")
    if err != nil {
        panic(err)
    }
	
    cfg := emailer.Config{}
    // Fill mail config using environment variables
    if err := envar.Fill(&cfg); err != nil {
        // All envs has to be set:
        //  MAIL_NAME – sender name
        //  MAIL_ADDR – sender email address
        //  MAIL_PASS – sender email password
        //  MAIL_SERV – email server address
        //  MAIL_PORT – email server port
        panic(err)
	}

    // Create Mail object with template and subject
    mail := emailer.New(temp, "A letter from a Program")

    // Change fields which are used to set recipient Name and Address
    //  this is optional and in this case default values are used
    userFields := emailer.ChangeUserFields("NAME", "MAIL")

    // Send emails to recipients and collect errors
    errors := mail.SendToMany(recipients, &cfg, userFields)
    for i, err := range errors {
        fmt.Printf("[error %d] %s\n", i, err)
    }
}

func simpleTemplate() string {
    return `
        <p>Dear, <span style="border-bottom: 2px solid rgb(148, 66, 255)">{{ .NAME }}</span></p>
        <p>I hope you and your {{ .PET }} are doing fine =)</p>
        <p>Best wishes to you, <br>Program</p>`
}
```

Received email:

```
Dear, Anton

I hope you and your cat are doing fine =)

Best wishes to you, 
Program
```

You can find this example with data in [this gist](https://gist.github.com/fabritsius/3f4b0a1b3a6a275c9411eb74e3ed2830).

## TODO

- [x] Add core features
- [x] Make send loop concurrent
- [ ] Add testing

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
