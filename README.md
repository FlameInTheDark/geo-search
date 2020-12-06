# geo-search
Geonames.org search api

Install:

`go get github.com/FlameInTheDark/geo-search`

Usage:

```go
package main

import (

gs "github.com/FlameInTheDark/geo-search"
"time"
)

func main() {
    client := gs.New("username", "en", 10, time.Second * 5)

    // Return an array of 10 elements with the New York and the nearest places
    places, err := client.Search("us new-york")

    // Return an array of 10 places which names contains 'New' substring
    places, err := client.ByName("New")

    // Return an array of 10 places which names exactly contains 'London'
    places, err := client.ByNameEquals("London")
}
```