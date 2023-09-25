# TL;DR

How do I use this thing. You can create messages, keys, DID and DIDocuments with this library, but you need to publish them yourself. This is a library, not a service.

## DID*

```go
package main

import (
  "fmt"

  "github.com/bahner/ma-go/did"
)

subEthaDID = did.New("space", subEthaKey.IPNSName.String())
  fmt.Println(subEthaDID)

```
