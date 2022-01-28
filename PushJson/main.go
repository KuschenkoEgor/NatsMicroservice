package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	sc, err := stan.Connect("test-cluster", "Admin")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Подключились!")

	jsonFile, err := os.Open("/home/zhora/Desktop/goolang-book/ProjectL0/PushJson/model.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened model.json")

	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println(string(byteValue))

	err = sc.Publish("foo", byteValue)
	if err != nil {
		log.Fatalf("Error during publish: %v\n", err)
	}
	sc.Close()



}
