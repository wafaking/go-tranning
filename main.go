package main

import (
	"fmt"
	"gotrainning/infra"

	"github.com/asaskevich/EventBus"
)



func main() {
	bus := EventBus.New();
	if err :=bus.Subscribe(infra.ToPicCalculation, infra.Calculator);err!=nil{
		panic(fmt.Sprintf("subErr: %s", err))
	}

	if err := bus.Unsubscribe(infra.ToPicCalculation, infra.Calculator);err!=nil{
		panic(fmt.Sprintf("UnSubErr: %s", err))
	}
	select {

	}
}