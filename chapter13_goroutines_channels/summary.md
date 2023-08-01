# Summary of Goroutines and Channels
> - Attention:This chapter is a little difficult at the first time to learn.

## The book I have read of this chapter:
- Head First Go


## What is goroutines:
- Goroutines are functions that are run concurrently.

## How to new a goroutines?
- It is very easy.New goroutines are started with a ***go*** statement:an ordinary function call preceded by the "go" keyworld. 
 
**Example**:
```go  
go myFunction()
go otherFunction("argument")
```
> Every Go program runs at least one goroutine,named main routine(the routine where the main function is started )

## Using goroutines
Suppose we have two functions now named a() and b(),now look at the code below,we will run them in different ways.  
- without keyword:go,both the func a and b are run in the same routine:main routine
    ```go
    func main(){
        a()
        b()
        fmt.Println("end main()")
    }
    ```
- Using keyword go,the func a and b will run at their own routine,they can run concurrently.
    ```go 
   func main(){
        go a()
        go b()
        fmt.Println("end main()")
   } 
    
    ```

> Go programs stop running as soon as the main **goroutine** ends,even if other goroutines are still running
### But there are still some problems:
- in the program above,the go program will go to an end,and the func a and b have no chance to run,because the main func go to an end quickly,thus making the main goroutine finish.  
- We can't get return value from the go func
   ```go
    size = go responseSize("https://example.com/")//You are saying "go run this;I am not going to wait.
    fmt.Println(size)//but then what is the return value?
   ```
Therefore,we need another feature of Go called channels.
 
## Channels:a way to communicate between goroutines
> Not only do channels allow you to send values from one goroutine to another,they ensure the sending routine has send the value before the accepting routine attempts to use it.  
### Things we need to demonstrate channels:
1. Create a channel
2. Write a func that receive a channel as a parameter.We will run this func at a separate goroutine,and use it to send values over the channel.
3. Receive the sent values in our original goroutine.

### How to declare and create a channel?
```go
var myChannel chan int//Type of values the channel will carry
myChannel = make(chan int)

or:

myChannel:=make(chan int)
```
### How to send and receive values with channel?
```go
myChannel <- 3.14//send the value to the channel
<- myChannel //receive the value from the channel
```
### Syschronizing goroutine with channel
We mentioned that channels also ensure the sending goroutine has sent the value before the receiving goroutine attempt to use it.  
Channels do that by pausing all further operations in the current goroutine.A send operation block the sending goroutine until the receiving goroutine executes a receive operation.And vice versa.
