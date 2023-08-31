![Version](https://img.shields.io/badge/version-0.0.0-orange.svg)

# Maps Utils

A basic golang package for demonstration purpose. Package currently contains only one function:

---

## Installation

Go to your project root, where `go.mod` file exists, than grab the library via:

```bash
go get github.com/SenRecep/mapsutils
```

---

## Usage

```go
package main

import (
	"fmt"
	"github.com/SenRecep/mapsutils"
	"log"
)

func main() {
	m1 := map[string]int{
		"key1": 12,
	}

	m2 := map[string]int{
		"key1": 12,
	}
	isEqual, err := mapsutils.DeepEqualMap(m1, m2)
	if err != nil {
		log.Fatal(err)
	}
	if isEqual {
		fmt.Println("Maps is equal")
	}
}
```

---

## Contributor(s)

* [recepsen](https://github.com/SenRecep) - Creator, maintainer

---

## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/SenRecep/mapsutils/fork)
1. Create your `branch` (`git checkout -b my-feature`)
1. `commit` yours (`git commit -am 'add some functionality'`)
1. `push` your `branch` (`git push origin my-feature`)
1. Than create a new **Pull Request**!

---

## License

This project is licensed under MIT

---

This project is intended to be a safe, welcoming space for collaboration, and
contributors are expected to adhere to the [code of conduct][coc].

[coc]: https://github.com//mapsutils/blob/main/CODE_OF_CONDUCT.md