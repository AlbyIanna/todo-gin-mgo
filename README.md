# ToDo Gin Mgo
This is a sample To-Do-List application implementing a RESTful API with [Gin Gonic](https://github.com/gin-gonic/gin) framework and [mgo](https://labix.org/mgo), a MongoDB driver for the [Go](https://golang.org) language.


# Setup
First thing, you have to install go. If you don't know how to do it, you can follow the instructions on the [Go installation page](https://golang.org/doc/install).

Next, we are going to use a MongoDB local instance, so you'll have to install it as well. Follow the [MongoDB installation page](https://docs.mongodb.com/manual/installation/), selecting your operating system.

Now you can clone this repo. Open a terminal and go to your GOPATH folder (usually the GOPATH in linux is `$HOME/go`; anyway, if you don't know what it is, I suggest you to read [this page](https://github.com/golang/go/wiki/GOPATH)). Run this commands:
```
cd src
git clone https://github.com/AlbyIanna/todo-gin-mgo.git
```

Finally we have to get the dependecies we are using. To do that, ensure you are in `$GOPATH/src`, then just run this commands:
```
cd todo-gin-mgo
go get
```
This will clone the repositories in we need in the **src** folder, so that we can use them just writing `import "NAME"` in our program.

# Run
Before running the API, we have to start the mongodb service, so that we can create and use our database. To do that, run this command:
```
sudo service mongod start
```
Verify that the `mongod` process has started successfully by checking the contents of the log file at `/var/log/mongodb/mongod.log`. If you open that file (or print the content using `cat /var/log/mongodb/mongod.log`), the last line should say something like `[initandlisten] waiting for connections on port 27017`. This means mongod is ready. If you want to stop mongod just run `sudo service mongod stop`.

Now cd into `$GOPATH/src/todo-gin-mgo` and run **main.go** like this
```
go run main.go
```
Now you can send http request to **create, list, update and delete** todos in your database.

# Test
To test the API use tools like [Postman](https://www.getpostman.com/) or `curl`.
