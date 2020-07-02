// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/timboldt/spiderbot/ssc32u"
)

func main() {
	const deviceName = "/dev/rfcomm0"
	const baudRate = 9600

	ssc := ssc32u.New()
	if !ssc.Connect(deviceName, baudRate) {
		fmt.Printf("Failed to connect to '%s' at %d baud\n", deviceName, baudRate)
		os.Exit(1)
	}
	//defer ssc.Close()

	ssc.AddServo(0, "tibia-1")
	ssc.AddServo(15, "tibia-2")
	ssc.AddServo(31, "tibia-3")
	ssc.AddServo(16, "tibia-4")
	ssc.AddServo(1, "femur-1")
	ssc.AddServo(14, "femur-2")
	ssc.AddServo(30, "femur-3")
	ssc.AddServo(17, "femur-4")
	ssc.AddServo(2, "coxa-1")
	ssc.AddServo(13, "coxa-2")
	ssc.AddServo(29, "coxa-3")
	ssc.AddServo(18, "coxa-4")

	// Sit down.
	ssc.Servo("femur-1").SetAngleDegrees(135)
	ssc.Servo("femur-2").SetAngleDegrees(45)
	ssc.Servo("femur-3").SetAngleDegrees(135)
	ssc.Servo("femur-4").SetAngleDegrees(45)
	ssc.Servo("coxa-1").SetAngleDegrees(45)
	ssc.Servo("coxa-2").SetAngleDegrees(135)
	ssc.Servo("coxa-3").SetAngleDegrees(45)
	ssc.Servo("coxa-4").SetAngleDegrees(135)
	ssc.Commit(100)
	time.Sleep(2 * time.Second)

	// Stand up.
	ssc.Servo("femur-1").SetAngleDegrees(90)
	ssc.Servo("femur-2").SetAngleDegrees(90)
	ssc.Servo("femur-3").SetAngleDegrees(90)
	ssc.Servo("femur-4").SetAngleDegrees(90)
	ssc.Servo("coxa-1").SetAngleDegrees(45)
	ssc.Servo("coxa-2").SetAngleDegrees(135)
	ssc.Servo("coxa-3").SetAngleDegrees(45)
	ssc.Servo("coxa-4").SetAngleDegrees(135)
	ssc.Commit(500)
	time.Sleep(2 * time.Second)

	// Sit down.
	ssc.Servo("femur-1").SetAngleDegrees(135)
	ssc.Servo("femur-2").SetAngleDegrees(45)
	ssc.Servo("femur-3").SetAngleDegrees(135)
	ssc.Servo("femur-4").SetAngleDegrees(45)
	ssc.Servo("coxa-1").SetAngleDegrees(45)
	ssc.Servo("coxa-2").SetAngleDegrees(135)
	ssc.Servo("coxa-3").SetAngleDegrees(45)
	ssc.Servo("coxa-4").SetAngleDegrees(135)
	ssc.Commit(500)
	time.Sleep(2 * time.Second)

	fmt.Printf("%v\n", ssc)

	ssc.Close()
}
