package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

func main() {
	for i := 0; i < 4; i++ {
		fmt.Printf("%v\n", service())
	}
}

func service() error {
	err := model()
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2)%2 == 0 {
		return errors.Wrap(err, "service wrap ")
	}
	return err
}

func model() error {
	err := db()
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2)%2 == 0 {
		return errors.Wrap(err, "model wrap ")
	}
	return err
}

func db() error {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2)%2 == 0 {
		return errors.New("db error")
	}
	return nil
}
