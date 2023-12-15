package main

import "fmt"

func logging() {
	fmt.Println("Logging App...")

	message := recover()
	fmt.Println("terjadi error ", message)
}

func runApplication(err bool) {
	defer logging()
	fmt.Println("Application RUN")
	if err {
		panic("Error BOS")
	}
}

func main() {
	runApplication(false)
}
