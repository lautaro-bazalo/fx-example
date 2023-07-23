# About project

Hi, I'm Computer Engineer and a few months ago I had to do an apoc of Dependency Injection in Golang.
You probably are thinking now that is a risky move use DY in a language like golang but, I think that just is another "tool" using reflection in background.
Looking out for some DY framework or something like that I found Uber Fx.

# Technologies involved

In this apoc I used the following technologies

• Gorm orm

• Gin-gonic Framework

• Docker compose

• Mysql database


# FX-Example 

In [FX](https://uber-go.github.io/fx/) there are four main concepts:

* `fx.Provide`  Registers all constructors
* `fx.Decorate` Registers all decorators 
* `fx.Invoke`   For function in your modules that must always run  
* `fx.Run`      Executes the app and runs all startup hooks appended to it.

![img.png](img.png)

I think that one big benefit of fx are the hooks to startup and shutdown the serve. You only have to define the functions to run and stop the server and forget to work with channels and signals.

### Running the APOC
To start up, there is a docker-compose.yaml it will create a mysql database.