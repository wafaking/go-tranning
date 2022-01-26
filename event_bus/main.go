package main

import (
	"github.com/asaskevich/EventBus"
	"gotrainning/infra"
)

func main() {
	for i:=0;i<10;i++{
	bus := EventBus.New();
	//bus.Subscribe("main:calculator", calculator);
	bus.Publish(infra.ToPicCalculation, 20, 40);
	//bus.Unsubscribe("main:calculator", calculator);
	}
}
