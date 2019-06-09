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
		// 	MAIL_NAME – sender name
		// 	MAIL_ADDR – sender email address
		// 	MAIL_PASS – sender email password
		// 	MAIL_SERV – email server address
		// 	MAIL_PORT – email server port
		panic(err)
	}
	
	// Create Mail object with template and subject
	mail := emailer.New(temp, "Just a letter")

	// Send emails to recipients
	if err := mail.SendToMany(recipients, &cfg); err != nil {
		panic(err)
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

## TODO

- [x] Add core features
- [ ] Add testing
- [ ] Make send loop asynchronous
- [ ] Add support for more use cases