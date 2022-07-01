package main

import (
	"fmt"
	"time"
)

type Err struct {
	_time time.Time
	_err  string
}

func (e *Err) Error() string {
	return fmt.Sprintf("At:%v Error :%v", e._time, e._err)
}

func fun() error {
	return &Err{time.Now(), "Not working"}
}
func main() {
	if err := fun(); err != nil {
		fmt.Println(err)
	}
}
