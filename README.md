# go-hyde

1hyde is equal to 156cm.


## Installation

Provides the `hyde` command.

```
go get github.com/kaneshin/go-hyde/cmd/hyde
```

or Installs to import `hyde` with the go get command.

```
go get github.com/kaneshin/go-hyde/hyde
```


## Usage

### Command

```
$ hyde 1 #=> 156
```

### Import

```
import (
	"fmt"
	"github.com/kaneshin/go-hyde/hyde"
)

func main() {
	fmt.Println(hyde.ToHyde(156))
	fmt.Println(hyde.FromHyde(1))
}
```


## License

[The MIT License (MIT)](http://kaneshin.mit-license.org/)


## Author

Shintaro Kaneko <kaneshin0120@gmail.com>

