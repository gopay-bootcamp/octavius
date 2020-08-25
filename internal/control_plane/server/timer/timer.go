package timer

import "time"

var (
	value         = 0
	sleepDuration = time.Second
)



type Timer interface{
	Start
}

type timer struct{
	value int64
	sleepDuration time.Duration
}

func GetTimer() Timer{
	return 
}
func Start() {
	time.Sleep(sleepDuration)
	value++
}

func GetValue() int {
	return value
}

func Reset() {
	value = 0
}
