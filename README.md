# Internal Certification Challenge

A project to challenge ourself with specific web stack to achieve the Internal Certification. 🚀

The requirement of this project is about extracting large amounts of data from the Google search result page.

## Web Application
- [Staging](https://go-challenge-staging.herokuapp.com/)
- [Production](https://go-challenge.herokuapp.com/)

## Prerequisite
* [Go - 1.15](https://golang.org/doc/go1.15)

## Usage

#### Clone the repository
```sh
$ git clone https://github.com/Lahphim/go-crawler-challenge.git
```

#### Install development dependencies
```sh
$ make install-dependencies
```
All dependencies:
- [Bee - Bee CLI](https://github.com/beego/bee)
- [Forego - Foreman in Go](https://github.com/ddollar/forego)

#### Run the application with development mode

Prepare the database and install some necessary packages.
```sh
$ make envsetup
```

Start the application.
```sh
$ make dev
```
Visiting http://localhost:8080/ with a web browser will display the application. ✨

#### Run tests
````sh
$ make test
````

## License
This project is Copyright (c) 2014-2021 Nimble. It is free software,
and may be redistributed under the terms specified in the [LICENSE] file.

[LICENSE]: /LICENSE

## About
![Nimble](https://assets.nimblehq.co/logo/dark/logo-dark-text-160.png)

This project is created to complete **Web Certification Path** using **Go** at [Nimble][nimble]

[nimble]: https://nimblehq.co
