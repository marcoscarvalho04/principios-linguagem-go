package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	inicio := time.Now()
	var controle sync.WaitGroup

	controle.Add(3)

	go calcular(9.37, &controle)
	go calcular(6.94, &controle)
	go calcular(42.57, &controle)

	controle.Wait()

	fmt.Printf("Finalizado em %s ", time.Since(inicio))
}

func calcular(base float64, controle *sync.WaitGroup) {
	defer controle.Done()
	n := 0.0

	for i := 0; i < 100000000; i++ {
		n += base / math.Pi * math.Sin(2)
	}
	fmt.Println(n)
}
