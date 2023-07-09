# FX-Example 

APOC about framework of dependency injection of uber [FX](https://uber-go.github.io/fx/) 

In Fx, there are four main concepts:

* `fx.Provide`  Registers all constructors
* `fx.Decorate` Registers all decorators 
* `fx.Invoke`   For function in your modules that must always run  
* `fx.Run`      Executes the app and runs all startup hooks appended to it.

![img.png](img.png)



#### Links of reference
* [Youtube channel Panpito](https://www.youtube.com/watch?v=UnrAF8FwfXU&t=1751s&ab_channel=Panpito)
* [Meduim Post](https://luannt2909.medium.com/apply-dependency-injection-with-uber-fx-golang-365d914189c1)

### Running the APOC
To start up, you may run Torwart's local db. You can use the configuration of Torwart repository to create your Docker image and run it. 