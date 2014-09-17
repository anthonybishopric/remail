## Remail, an STMP rerouter

Remail is a simple command that launches an SMTP server on the configured port and emits a JSON blob for each received email.

## Command-line Usage

```
$ remail
{"from":"ab@squareup.com","to":["bab@squareup.com"],"addr":"[::1]:52661","content":"From: abishopric@squareup.com\r\nTo: lab@squareup.com\r\nSubject: test message\r\nDate: Right now\r\nMessage-Id: somestring\r\n\r\n This is a test message\r\n"}
{"from":"cd@squareup.com","to":["dab@squareup.com"],"addr":"[::1]:52661","content":"From: abishopric@squareup.com\r\nTo: jab@squareup.com\r\nSubject: test message\r\nDate: Right now\r\nMessage-Id: somestring\r\n\r\n This is a test message\r\n"}
```

The default invocation of remail simply listens on port 25 and emits a json document for each email received. To override the port, use the SMTP_PORT environment variable.

`$ SMTP_PORT=:2500 remail`

Since each line is passed on standard out, client programs should be written with the goal of handling each line as an individual email.

## As a Go Library

The SMTP server can be started by invoking `remail.Serve` with a callback for each email received. 

```go
package foo

import "github.com/anthonybishopric/remail/pkg"

func StartFakeSMTPServer() {
    remail.Serve(func(rec remail.ReceivedEmail, jsonContent []byte) error {
        if rec.From == "someone@important.com" {
            // do something useful
        }
        return nil
    })
}

```

## License
Apache 2.0
