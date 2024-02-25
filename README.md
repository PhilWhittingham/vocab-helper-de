# German Vocabulary Cataloguer in Golang

Created in an effort to both;
* Demonstrate and practice transferable software engineering skills with Go.
* Provide a "digital flashcards" solution I can use to improve my German language retention

## Description

An elaborate, over-engineered backend solution for cataloging German vocabulary.

This project utilizes tools and architectures that I have become comfortable with during my time as an engineer. It can be used as a minimal example for the packages and concepts used.


## Pre-requisites

Confirmed working using `Go 1.12`, older versions may work.

## Running Instructions

Build the environment using

```go get .```<br />

Run the web service using

```go run .```<br />

Run the tests using

```go test ./...```<br />

## Built With

* [Go 1.21](https://go.dev/)


## Inspired By
[go-clean-architecture](https://github.com/zhashkevych/go-clean-architecture) provided a majority of the template and, if I'm being honest, the code. It's truly a great resource if you're already used to [Uncle Bob's Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html), Hexagonal Architecture, Domain Driven Design and supporting [12 Factor Applications](https://12factor.net/).

[python-http-thing-doer](https://github.com/PhilWhittingham/python-http-thing-doer) for allowing me to map areas of the above repository to my understanding gathered from Python development. Some areas I didn't carry over from go-clean-architecture were because of my preferred approach demonstrated in this Python repository.