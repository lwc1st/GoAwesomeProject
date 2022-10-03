package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type ApplePhone struct {
}

func (applePhone ApplePhone) call() {
	fmt.Println("I am Apple, I can call you!")
}

type XiaoMiPhone struct {
}

func (xiaoMiPhone XiaoMiPhone) call() {
	fmt.Println("I am XiaoMi, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(ApplePhone)
	phone.call()

	phone = new(XiaoMiPhone)
	phone.call()
}