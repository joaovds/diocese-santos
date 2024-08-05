package main

import (
	"fmt"

	"github.com/joaovds/diocese-santos/internal/liturgy"
)

func main() {
	liturgyUsecases := liturgy.NewLiturgyUsecases()
	fmt.Println(liturgyUsecases.GetCurrentLiturgicalInfo())
}
