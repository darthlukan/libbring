# LibBring v1.0

[![Build Status](https://drone.io/github.com/darthlukan/libbring/status.png)](https://drone.io/github.com/darthlukan/libbring/latest)

> Author: Brian Tomlinson <brian.tomlinson@linux.com>


## Description

> LibBring is a Go wrapper for some of the most commonly used functions of the Bring API
> located at http://developer.bring.com.


## Usage

> Example:

```
    import (
        "fmt"
        "github.com/darthlukan/libbring"
    )

    func main() {

        data, err := libbring.Track("TESTPACKAGE-AT-PICKUPPOINT")

        if err != nil {
            panic(err)
        }

        fmt.Printf("Tracking info: %v", data)
    }
```

> For more information on how to use this library, see the code in ```libbring/libbring.go```, the
> functions in that file are very much self-explanatory.  [Bring API Documentation](http://developer.bring.com/index.html).


## License

> The Go code in this repo falls under GPLv2, see the LICENSE file.  The Bring API itself falls under separate
> licensing conditions but is provided free of charge.  For more information, [visit their developer site.](http://developer.bring.com/index.html)
