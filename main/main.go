package main

import (
	"Cave-CTF/log"
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello,Cave-CTF")
	log.NewLog(log.DEBUG, 200, errors.New("tes"), "sss")
	time.Sleep(1 * time.Second)
}
