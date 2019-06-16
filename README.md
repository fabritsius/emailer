# emailer

[![GoDoc](https://godoc.org/github.com/fabritsius/emailer?status.svg)](https://godoc.org/github.com/fabritsius/emailer)

This Go module simplifies a process of sending HTML emails to multiple users.

## Example

Simple use case.

You want to send an email to each one of your subscribers (here just one). Here is how you can do that:

```go
package main

import (
    "fmt"

    "github.com/fabritsius/emailer"
)

func main() {
    // Get an email template
    temp := simpleTemplate()

    // Example user data
    user := map[string]string{
        "NAME": "Anton",
        "MAIL": "user@example.com",
        "PET":  "cat",
    }

    // Create a slice of recipients (here is just one)
    recipients := []map[string]string{user}

    // Create a config for smtp connection
    cfg := emailer.Config{
        Name:     "Program",
        Mail:     "sender@example.com",
        Password: "unbreakable",
        Server:   "smtp.example.com",
        Port:     "465",
    }

    // Create Mail object with template and subject
    mail := emailer.New(temp, "A letter from a Program")

    // Send emails to recipients and collect errors
    errors := mail.SendToMany(recipients, &cfg)
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

You can find a better example without hard coded variables in [this gist](https://gist.github.com/fabritsius/3f4b0a1b3a6a275c9411eb74e3ed2830).

## Features

- use `New(template, subject)` to create `Mail` object
- use `SendTo` or `SendToMany` functions on the `Mail` object
- `SendTo(recipient, config)` allows to send an email to a recipient
- `SendToMany(recipients, config)` sends emails to multiple recipients
- both methods can take `ChangeUserFields` function as an optional parameter

Full method documentation can be found on [GoDoc page](https://godoc.org/github.com/fabritsius/emailer).

## TODO

- [x] Add core features
- [x] Make send loop concurrent
- [ ] Add testing

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
