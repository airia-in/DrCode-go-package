# DrCode - Go Package

`drcode` is a Go package that provides easy integration with error tracking services, allowing you to capture and report errors or panics in your Go web applications. It is designed to simplify error handling by providing middleware and utility functions that automatically send error reports to your error tracking service.

## Features

- **Initialize Error Tracking**: Set up your error tracking service with your project ID and public key.
- **Panic Handling**: Automatically capture and report panics with the provided middleware.
- **Error Reporting**: Manually report errors to your error tracking service.

## Installation

First, ensure you have Go installed on your system. Then, you can install the `drcode` package using `go get`:

```bash
go get github.com/airia-in/DrCode-go-package
```

Import the package into your Go project:

```bash
import drcode "github.com/airia-in/DrCode-go-package"
```

## Usage

### 1. Initialize Error Tracking

Before using any error handling or reporting functions, you must initialize your error tracking service with your project ID and public key.

```bash
package main
import (
    "log"
    drcode "github.com/airia-in/DrCode-go-package"
)
func main() {
    err := drcode.Initialize("your-project-id", "your-public-key")
    if err != nil {
        log.Fatalf("Failed to initialize error tracking: %v", err)
    }
    // Your application code here
}
```

## 2. Panic Handling Middleware

Use the ErrorHandler middleware to automatically capture and report panics in your HTTP handlers.

```bash
package main

import (
    "net/http"

    drcode "github.com/airia-in/DrCode-go-package"
)

func main() {
    err := drcode.Initialize("your-project-id", "your-public-key")
    if err != nil {
        log.Fatalf("Failed to initialize error tracking: %v", err)
    }

    http.HandleFunc("/panic", drcode.ErrorHandler(func(w http.ResponseWriter, r *http.Request) {
        panic("This is a test panic!")
    }))

    http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
        err := fmt.Errorf("this is a test error")
        drcode.ReportError(err)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("An error occurred and was reported"))
    })

    log.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## 3. Manual Error Reporting

You can manually report errors to your error tracking service using the ReportError function.

```bash
package main

import (
    "fmt"
    "net/http"

    drcode "github.com/airia-in/DrCode-go-package"
)

func main() {
    err := drcode.Initialize("your-project-id", "your-public-key")
    if err != nil {
        log.Fatalf("Failed to initialize error tracking: %v", err)
    }

    http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
        err := fmt.Errorf("this is a test error")
        drcode.ReportError(err)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("An error occurred and was reported"))
    })

    log.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
