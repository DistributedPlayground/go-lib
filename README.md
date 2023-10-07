# Go-Lib
This repo is responsible for centralizing some core go components into a single library for use by other services. It handles errors, logging, validation, loading environment variables, and utility functions.
It is a fundamental part of our E-commerce platform within the [DistributedPlayground](https://github.com/DistributedPlayground) project. See the [project description](https://github.com/DistributedPlayground/project-description) for more details.

## Usage
Although it's not necessary to clone this repo to run the Distributed Playground project, `Go-Lib` can be easily integrated into any Go project. Here's an example of how you can import one of it's packages:

```go
import (
    "github.com/DistributedPlayground/go-lib/httperror"
)

// Usage
return httperror.Internal500(ctx)
```