package main

import (
	"fmt"
	"sync"
)

func printSalam() {
	fmt.Println("Hai")
}

var itemsChannel = make(chan string)
var treasureChannel = make(chan *string)

func main() {
	//go printSalam()
	//fmt.Println("Hello")

	//Tutor 1
	//laut := []string{"batu", "harta", "pasir", "tanah", "harta", "ikan", "kelapa", "kapal", "harta"}

	//go menyelam(laut)
	//go membersihkan()
	//go menyimpan()

	//time.Sleep(500 * time.Millisecond)

	//Tutor 2
	var wg sync.WaitGroup

	wg.Add(2)
	go printText("Hallo", &wg)
	go printText("Hai", &wg)
	wg.Wait()
}

func menyelam(items []string) {
	for _, item := range items {
		if item == "harta" {
			fmt.Println("Berhasil get " + item)
			itemsChannel <- item
		}
	}
}

func membersihkan() {
	for rawItem := range itemsChannel {
		fmt.Println("Berhasil clear " + rawItem)
		treasureChannel <- &rawItem
	}
}

func menyimpan() {
	for rawItem := range treasureChannel {
		fmt.Println("Berhasil simpan", *rawItem)
	}
}

func printText(text string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	wg.Done() //memberi tahu wg bahwa goroutine ini sudah selesai
}
