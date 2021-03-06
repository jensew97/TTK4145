package main

import (
	"fmt"
	"time"
)

type Slave struct {
	IP     string
	number int
}

type Master struct {
	slaves map[Slave]time.Time
	num    int
}

func foo(chan_test chan [][]int) {
	test := <-chan_test
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			test[i][j] = 2
		}
	}
	chan_test <- test
}

func main() {
	master := Master{}
	slave := Slave{"123", 1}
	slave.IP = "456"
	newborn := Slave{"789", 3}
	master.slaves = make(map[Slave]time.Time)

	chan_test := make(chan [][]int)
	foo(chan_test)
	// fmt.Println(master.slaves)
	// master.slaves[slave] = time.Now()
	// master.slaves[newborn] = time.Now()
	// // for s := range master.slaves {
	// // 	fmt.Println(s.IP)
	// // }
	// // time.Sleep(1 * time.Second)
	// // const DEADLINE = 1 * time.Millisecond
	// // for s, last_time := range master.slaves {
	// // 	fmt.Println(last_time)
	// // 	if time.Since(last_time) > DEADLINE {
	// // 		delete(master.slaves, s)
	// // 	}
	// // }
	// baby := Slave{"000", 0}
	// master.slaves[baby] = time.Now()
	// fmt.Println(master.slaves)
}
