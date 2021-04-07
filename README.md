# Internal Certification Challenge

A project to challenge ourself with specific web stack to achieve the Internal Certification. 🚀

The requirement of this project is about extracting large amounts of data from the Google search result page.

## Web Application
- [Staging](https://go-challenge-staging.herokuapp.com/)
- [Production](https://go-challenge.herokuapp.com/)

## Prerequisite
* [Go - 1.16](https://golang.org/doc/go1.16)

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

#### Background Tasks
Using [tasks](https://beego.me/docs/module/toolbox.md#tasks) which is provided in the BeeGo's [toolbox](https://beego.me/docs/module/toolbox.md) module.

The mechanism of this module is very similar to cron jobs 🍀.

So we can create a task and assign the schedule of time to the task, then we can do whatever after the task is triggered by the time we set.

Initializing all the tasks from here `conf/initializers/task.go` and addressing those tasks within this path: `/tasks/*_task.go`

Example:
Setting up the task to run in every minute (https://beego.me/docs/module/toolbox.md#spec-in-detail).
```golang
searchKeywordTask := SearchKeywordTask{Name: "search_keyword_task", Schedule: "0 * * * * *"}
searchKeywordTask.Setup()

...
```

Add the task then all of them will be executed with `StartTask()`.
```golang
task.AddTask(searchKeywordTask.Name, searchKeywordTask.Task)
task.AddTask(***, ***)
task.AddTask(***, ***)

task.StartTask()
```

## API Documentation
- [Postman](https://documenter.getpostman.com/view/105704/TzCMc7RN)

## License
This project is Copyright (c) 2014-2021 Nimble. It is free software,
and may be redistributed under the terms specified in the [LICENSE] file.

[LICENSE]: /LICENSE

## About
![Nimble](https://assets.nimblehq.co/logo/dark/logo-dark-text-160.png)

This project is created to complete **Web Certification Path** using **Go** at [Nimble][nimble]

[nimble]: https://nimblehq.co
