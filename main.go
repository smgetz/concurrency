package main

import (
	"fmt"
	"time"
)

//^^^^^^ Lines 112-130 is the completed example of how to use concurreny and channels the correct way, comment out code as need to run and see results in CLI ^^^^^//

//Placing a go statement before a funciton call starts a concurrent execution of that function, this concurrent thread of executino is called a goroutine
func main() {
	//by using the go contruct the function execution can be made concurrent
	go my_sleep_func()
	fmt.Println("Control doesnt reach here till my_sleep_func finishes executing")
	fmt.Println("the intution is similar to being asynchrinous without callback!")

	//sleeping in main to make sure the main doesnt end
	time.Sleep(6 * time.Second)
}

func my_sleep_func() {
	//sleeps for 5 seconds
	time.Sleep(5 * time.Second)
	fmt.Println("My func out of sleep")
}

//
//
//

//Trying to synchronize and pass messages between these concurrently running Go routines
//my_func is called concurrently and it sleeps for amilisecond in every iteration and this gives the main function enough time ot iterate 10 times and finish
//Golang channels can be vizualized as a bridge between Goroutines, a bridge which can load a finite number(if 10 then no more than 10, only until one person to the other end)
func main() {

	go my_func()
	for i := 0; i < 10; i++ {
		fmt.Println("in the main function")
	}
	time.Sleep(1 * time.Second)
}

func my_func() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println("in the Go routine")
	}
}

//
//
//

func main() {
	//building a bridge
	//important to specify the data type the bridge would carry
	bridgeForOne := make(chan int, 1) //specifying the size of the bridge to be one

	//lets onboard a integer onto the bridge
	//this is ow you onboard a person "bridge <-1" and this is how to get the person off the bridge "<-bridge"
	//will block if more elements are listed than the int of channel
	bridgeForOne <- 1

	fmt.Println("integer sent onto the bridge")
}

//
//
//

//Goal of channels is to synchronize and safely share data between concurrently runnin goroutines
//Trick is to share the reference of the channel with the goroutines those you desire to sync and communicate
//Steps
//Create a channel of desired capacity
//Call a funciton in a new go routine and shre the reference of the channel with the goroutine/goroutines with which you wish to comm with
//Start sending the data through the channels through one or many goroutines and receive it in the other

func main() {
	//step 1: create bridge
	bridgeForOne := make(chan int, 1)
	//call the function in a new goroutine
	//and send the reference of channel to the new goroutine
	//using this reference the new goroutine will
	//communicate and sync with main
	go testRunConcurrent(bridgeForOne)
	//lets onboard a int onto the bridge
	bridgeForOne <- 1122

	fmt.Println("Int recieved at the other end of the bridge in the new go routine and hence the main unblocks and control comes here")
}

func testRunConcurrent(bridgeReferenceFromMain chan int) {
	//Inside the new Go routine
	fmt.Println("Inside the new goroutine")

	//recieving the integerof the channel bridge

	takingTheIntegerOffTheBridge := <-bridgeReferenceFromMain

	//only after this recieve from the other end of the shared channel,
	//the send through the channel on main function unblocks and continues execution

	fmt.Printf("\nHere is the number sent across the bridge from main: %d\n", takingTheIntegerOffTheBridge)
	//since the function called from your new goroutine ends execution here, the new goroutines also gracefully ends
}

//
//
//

func main() {
	bridge := make(chan int) //creating channel
	go my_func(bridge)       //creating a new go routine and sharing the reference of the channel
	for i := 0; i < 10; i++ {
		bridge <- 1 //sending int through the bridge(channel), this is blocked until recieved in the goroutine
		fmt.Println("in the main function")
	}
	time.Sleep(1 * time.Second)
}

func my_func(referenceToMainBridge chan int) {

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		<-referenceToMainBridge //the recieved value is not used, its just disposed
		//main is blacked till its recieved here
		fmt.Println("in the go routine")
	}
}
