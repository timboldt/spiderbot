# Arduino version of Spider Bot

The Arduino "HAL" is a very convenient platform for hobbyist projects like this one.

With a little bit of extra work, you can also have a semi-professional tooling setup.

Here's what I am using:
* M1 Macbook Pro.
* VSCode for the IDE.
* The Arduino CLI.
* AUnit for unit testing on the target hardware.

## Building

### Pre-reqs

These instructions are for the Arduino CLI, but you should be able to do the equivalent in the IDE.

On a Mac, you can install the Arduino CLI with Homebrew:

```
brew install arduino-cli
```

To use an Adafruit core, add this to your Arduino config file (mine is in `~/Library/Arduino15/arduino-cli.yaml`):

```
board_manager:
  additional_urls:
    - https://adafruit.github.io/arduino-board-index/package_adafruit_index.json
```

And then add your specific core, like this:

```
arduino-cli core install adafruit:nrf52
```

Here are the specific libraries I was using when I wrote this readme:

```
Adafruit_BusIO                    1.7.2
Adafruit_PWM_Servo_Driver_Library 2.4.0
AUnit                             1.5.4
```

### Compiling and Flashing

To compile and flash with the Arduino CLI:

```
arduino-cli compile --fqbn adafruit:nrf52:feather52840 -u -p /dev/cu.usbmodemXXXX
```

Note: replace `cu.usbmodemXXX` with an appropriate TTY or COM device for your platform.

### Running Unit Tests

To run the unit tests, use the same command, but from inside the `src/spider` directory.

After flashing, connect to the serial console to see the results of the unit tests.