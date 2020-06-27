package main

import (
	"fmt"

	"github.com/timboldt/spiderbot/ssc32u"
)

func main() {
	ssc := ssc32u.New("/dev/tty1", 9600)
	defer ssc.Close()

	ssc.AddServo(1, "test1")
	ssc.AddServo(14, "fourteen")

	ssc.Servo(1).SetPosition(1300)

	fmt.Printf("%v\n", ssc)
}
