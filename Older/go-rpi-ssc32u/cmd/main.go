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

	"github.com/timboldt/spiderbot/spider"
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

	ssc.AddServo(0, "wrist-0")
	ssc.AddServo(15, "wrist-1")
	ssc.AddServo(31, "wrist-2")
	ssc.AddServo(16, "wrist-3")
	ssc.AddServo(1, "knee-0")
	ssc.AddServo(14, "knee-1")
	ssc.AddServo(30, "knee-2")
	ssc.AddServo(17, "knee-3")
	ssc.AddServo(2, "hip-0")
	ssc.AddServo(13, "hip-1")
	ssc.AddServo(29, "hip-2")
	ssc.AddServo(18, "hip-3")

	b := spider.NewBody()
	b.MoveLegAbsolute(spider.FrontRight, 40, 40, -80)
	//l := b.GetLeg(spider.FrontRight)
	// wrist, knee, hip := l.GetAngles()
	//_, _, hip := l.GetAngles()
	// ssc.Servo("wrist-1").SetAngleDegrees(wrist)
	// ssc.Servo("knee-1").SetAngleDegrees(knee)
	//ssc.Servo("hip-1").SetAngleDegrees(-hip)
	ssc.Commit(500)
	time.Sleep(2 * time.Second)

	//fmt.Printf("%v\n", ssc)

	ssc.Close()
}
