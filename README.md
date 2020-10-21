# golang-sandbox

This repository was created with the intention of being a simple sandbox for testing Golang. The project itself contains a number of separate 'projects' which use some of the basic features of Go. This is intended to be purely educational. Sp far, this project is following a number of examples as outlined in the course: [Go: The Complete Developer's Guide (Golang)](https://www.udemy.com/course/go-the-complete-developers-guide/). 

## Projects

So far, the sandbox contains two projects:

* helloworld
* cards

Each project's description can be found below:

### helloworld

Just a basic 'Hello World!' program for Go.

To run the project, just `cd` to the helloworld directory and execute the command:

```
go run main.go
```

### cards

A program which tests arrays and slices by simulating a deck of cards. 

The deck type simulates a standard deck of 52 cards. The deck can be saved and loaded from a file, can deal out a variable number of cards to users, and can be shuffled randomly. This small project demonstrates the use of the following Go features:

* Slices
* For Loops
* File I/O
* Random Number Generation
* Packages
* tests

You can run the project by navigating the the cards directory and execute the command:

```
go run main.go deck.go
```

You can also test the project by running:

```
go test
```

### evenorodd

Just a small project where we implement an 'even or odd' function. This simply checks an integer for whether it's even or odd and prints the output. 

To run the project, navigate to the 'evenorodd' directory and run the following:

```
go run main.go
```

## Setup

Make sure you have Golang installed on your computer. If it isn't installed, you can download the full package from [golang.org](https://golang.org/). 

Simply clone the repository to your desired directory:

```
git clone https://github.com/Raelr/golang-sandbox.git
```

Once cloned, navigate to the project that you wish to test and run the golang run command:

```
go run main.go <OTHER_.GO_FILES_IN_DIRECTORY>
```

This should hopefully run the go files for the project!

Enjoy! 
