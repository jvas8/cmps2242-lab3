# Lab #3: Structs, Methods, and Web Applications
Course: CMPS2242  
Student Name: Jeimy Vasquez 

## Overview
This repository contains the solution for Lab #3. The project is divided into two distinct Go modules demonstrating core language features:
1. Shapes: A standalone library and executable demonstrating Go structs, value/pointer receivers, and unit testing.
2. Web API: A server demonstrating routing, handlers, and table-driven integration tests.

## Project Structure
lab3-username/
├── shapes/                 # Part 1: Geometric Shapes
│   ├── go.mod
│   ├── main.go             # Entry point for shapes demo
│   ├── shapes.go           # Struct definitions and method logic
│   └── shapes_test.go      # Unit tests for shapes
├── web-api/                # Part 2: Web Application
│   ├── go.mod
│   └── cmd/
│       └── api/
│           ├── main.go     # Server and route handlers
│           └── main_test.go # Table-driven HTTP handler tests
└── README.md