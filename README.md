# go-mail

[![GoDoc](https://godoc.org/github.com/codykaup/go-mail?status.svg)](https://godoc.org/github.com/codykaup/go-mail)

`go-mail` creates a simple wrapper around [net/mail.Message] to extend common
functionality.

## Installing

```
$ go get -u github.com/codykaup/go-mail
```

## Example

### Reading a message

```golang
import (
	"bufio"
	"fmt"
	"os"

	mail "github.com/codykaup/go-mail"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	message, _ := mail.ReadMessage(reader)

	fmt.Printf("Headers: %+v\n", message.Header)
	fmt.Printf("Body: %+v\n", message.Body)
}
```

### Append a header

```golang
import (
	"bufio"
	"fmt"
	"os"

	mail "github.com/codykaup/go-mail"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	message, _ := mail.ReadMessage(reader)

	fmt.Printf("Headers: %+v\n", message.Header)

	entry, _ := mail.NewHeaderEntry("Cc", "john@example.com")
	message.AppendHeader(entry)

	fmt.Printf("New Headers: %+v\n", message.Header)
}
```

### Get the full message

```golang
import (
	"bufio"
	"fmt"
	"os"

	mail "github.com/codykaup/go-mail"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	message, _ := mail.ReadMessage(reader)

	fmt.Println(message.Join())
}
```
