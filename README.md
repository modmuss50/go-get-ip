
# go-get-ip  
A simple library to get the current ipv4 or ipv6 address using [ident.me](https://api.ident.me/)
This project does not use any external libraries and all requests are made over https.
  
# Usage:

To download this library just run:

```console
go get github.com/modmuss50/go-get-ip
```

The api is extremely easy to use, see the example below: 

```go
package main

import (
"fmt"
"github.com/modmuss50/go-get-ip"
)

func  main() {
    ip, err  := go_get_ip.GetIPV4()
    if err !=  nil {
        //Failed to get the ip address, this may happen if you do not have internet access
        panic(err)
    } else {
        fmt.Println("IPv4: "  + ip)
    }
    
    ipv6, err  := go_get_ip.GetIPV6()
    if err !=  nil {
        //Failed to get the ip address, this may happen if you do not have internet access
        //This will only work if the network is properly configured to work with ipv6
        panic(err)
    } else {
        fmt.Println("IPv6: "  + ipv6)
    }
}
```
### Things to note:

`GetIPV6` will only work when the network is configured to work with ipv6, so don't forget to gracefully handle the errors.

If you need help feel free to ask questions on our discord server. [![](https://img.shields.io/badge/Discord-TeamReborn-738bd7.svg)](https://discord.gg/teamreborn)

