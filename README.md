# YayGo

Golang API wrapper for [Yay!](https://yay.space)

<img align="right" alt="Gravity" src="https://github.com/user-attachments/assets/8e9c4c9c-47e2-43a2-9006-8846356f4f02" width="400">

## Getting Started

### Installation

```sh-session
go get "github.com/popo-nyan/yaygo"
```

### Usage

Import the package into your project.

```go
import "github.com/popo-nyan/yaygo"
```

Construct a new YayGo client session.

```go
yay, err := yaygo.New("your_email", "your_password")
```

See Documentation and Examples below for more detailed information.

### Example

```go
package main

import (
    "github.com/popo-nyan/yaygo"
)

func main() {
    yay, err := yaygo.New("your_email", "your_password")
    if err != nil {
        panic(err)
    }

    resp, err := yay.Post.CreatePost(&CreatePostParams{
        Text: "Hello with YayGo!",
    })
    if err != nil {
        panic(err)
    }
    fmt.Println(resp.Result)
}
```

## Documentation

Coming soon!
