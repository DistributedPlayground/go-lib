# Go-Lib
This repo is responsible for centralizing some core go components into a single library for use by the Distributed Playground projects. It handles errors, logging, validation, loading environment variables, and utility functions.

## Usage
Although it's not necessary to clone this repo to run the Distributed Playground project, `Go-Lib` can be easily integrated into any Go project. Here's an example of how you can import one of it's packages:

```go
import (
    "github.com/DistributedPlayground/go-lib/httperror"
)

// Usage
return httperror.Internal500(ctx)
```

# Distributed Playground
The purpose of this repo is to practice the development of distributed systems
## Project 1 - Ecommerce Platform
For my first project, I'll create an ecommerce platform. This platform will allow users to purchase goods that are maintained by us. This is *not* a two sided marketplace with our users being both companies and consumers -- at least for now. To simplify, we'll centrally define the products.

### Architecture

#### Functional Requirements
- Users should be able to make purchases using card payments
- Users should be able to see available products, including their remaining stock
- Users should be able to sort and filter products

#### Quality Attributes
- Should be designed to handle >10M users per day
- Should have 99.9% uptime
- Should allow for concurrent development from a large distributed team

#### Constraints
- The core services must be completed by 1 engineer (me) in < 1 month. 

#### Design
- We will use Golang as the primary backend language.

- We will provide an internal CP db optimized for writes in postgres
    - This db will be exposed to our systems through a REST API
- Implementing a psudo CQRS (Command Query Responsibility Segregation)
    - "Psudo" here because we really just want to segregrate the most read heavy users (customers) to a separate db and api optimized for their needs
    - This customer read db will be updated through a service that reads from a Kafka queue

#### Services
- **API Gateway**: *Not Implemented*
- [Products](https://github.com/DistributedPlayground/products): A REST API with a postgres db. It allows sellers to manage their products and product collections. This service publishes writes to a kafka queue.
- [Product Search](https://github.com/DistributedPlayground/product-search): A GraphQL API with mongodb. It allows customers to query the current products and collections. It reads updates from kafka to update the mongodb database.
- [Inventory](https://github.com/DistributedPlayground/inventory): A gRPC API with redis. It is intended to maintain the most up-to-date state of product inventory. It reads updates to product inventory made by sellers from kafka, and also writes updates to the product inventory made by customers through purchases.
- **Orders**: *Not Implemented* Orchestrator for *Payments*, *Fufillment*, and *Notifications*, state management with kafka
- **Order Recovery**: *Not Implemented*