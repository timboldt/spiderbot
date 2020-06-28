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
	defer ssc.Close()

	ssc.AddServo(1, "test1")
	ssc.AddServo(14, "fourteen")

	ssc.Servo(1).SetPosition(1300)
	ssc.Commit(100000)

	fmt.Printf("%v\n", ssc)
}
