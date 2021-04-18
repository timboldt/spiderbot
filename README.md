# spiderbot

Spider Bot based on [RegisHsu's 3D-printable quadruped](https://www.instructables.com/id/DIY-Spider-RobotQuad-robot-Quadruped/).

The current version runs on a Adafruit Feather nrf-series and (planned; not implemented yet) supports BLE control. The servo control is done with an [Adafruit PCA9685](https://learn.adafruit.com/16-channel-pwm-servo-driver), which is a 16-channel PWM device controllable via an I2C interface.

There are two implementations of the current version:
* `arduino/` contains an Arduino/C++ implementation. See [the readme](arduino/README_Arduino.md) in that folder for details. 
* `tinygo/` contains a Tiny-Go implementation. It can be flashed with `tinygo flash -target=feather-nrf52840 cmd/main.go` from within that directory.

There are two older implementations in `Older/`, which ran on a Linux system (an RPi 3B+, in my case) and did servo control with a [SSC-32U](http://www.lynxmotion.com/p-1032-ssc-32u-usb-servo-controller.aspx) over a Bluetooth UART. The Python version is the more mature of the two, but neither version is really complete. In particular, I didn't understand coordinate reference frames very well when I designed them.

## Disclaimer

This is not an officially supported Google product.
