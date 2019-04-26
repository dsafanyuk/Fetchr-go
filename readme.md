# Fetchr-go

## Quick Info

This is my attempt to learn Go by reimplementing my senior project [Fetchr](github.com/dsafanyuk/fetchr) backend. (originally written in Node.js).

Fetchr is a service that has 2 main functions: 

- users  browse products, add items to carts, checkout
- couriers  browse open orders, accept open orders, deliver their assigned orders

## Code Structure


This project will try to follow Domain Driven Design and Hexagonal Architecture.

<details>

<summary> File structure </summary>

```
├── cmd // not sure about using this as the way to run packages yet
│   └── user 
│       └── main.go
├── config // database config goes here for now
│   └── config.go
└── pkg
   ├── database
   │   └── psql
   │       └── user.go
   ├── middleware
   │   └── auth.go
   └── user // user domain
       ├── handler.go // handles http 
       ├── model.go   // user struct
       ├── repo.go    // user interface
       └── service.go // user service

```

</details>
