# spiderbot

Spider Bot based on [RegisHsu's 3D-printable quadraped](https://www.instructables.com/id/DIY-Spider-RobotQuad-robot-Quadruped/).

The current version is written in TinyGo. It runs on a Adafruit Feather nrf-series and supports BLE control. The servo control is done with an [Adafruit PCA9685](https://learn.adafruit.com/16-channel-pwm-servo-driver), which is a 16-channel PWM device controllable via an I2C interface.

The original version was in Python, which ran on a Linux system and did servo control with a [SSC-32U](http://www.lynxmotion.com/p-1032-ssc-32u-usb-servo-controller.aspx) over a Bluetooth UART.

## Disclaimer

This is not an officially supported Google product.
