<hr>
<div align="center">

This is a port of [the original](https://github.com/RienNeVaPlus/human-id) to Go.

</div>
<hr>


## Nice Identifiers

NiceID generates readable strings by chaining common short English words in a semi-meaningful way.
The result is concatenated of `adjective + adjective + noun + verb + adverb` resulting in a pool size of **54,174,285,000** possible combinations.

- **SFW**: no bad words; family friendly results
- No dependencies


## Examples

- some-loose-roses-crash-safely
- late-moody-stars-slide-brightly
- old-true-baths-relate-busily
- orange-floppy-lines-wink-freely


## Usage

```bash
go get github.com/austinjp/niceid
```

```go
package main

import (
	"fmt"
	"github.com/austinjp/niceid"
)

func main() {
	id := niceid.ID()
	fmt.Println(id)
}

```
