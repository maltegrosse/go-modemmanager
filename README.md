Go-ModemManager
================
[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)

Go D-Bus bindings for ModemManager
Additional information: [ModemManager D-Bus Specs](https://www.freedesktop.org/software/ModemManager/api/1.12.0/ref-dbus.html)

Tested with [ModemManager - Version 1.28.8](https://gitlab.freedesktop.org/mobile-broadband/ModemManager) and Go 1.13
with a [SolidRun Hummingboard Edge](https://www.solid-run.com/nxp-family/hummingboard/) on `Debian Buster (armv7)` with `kernel 5.4.x` and `libqmi 1.24.6` and a `Quectel EC25` miniPcie modem.

## Status
Work in Progress, some methods/properties/signals needs to be fixed for initial release of version 0.1

## Todo
Some methods/properties are untested as they are not supported by my modem/lack of how to use them. See `todo` tags in the code.

## Installation

This packages requires Go 1.13 (for the dbus lib). If you installed it and set up your GOPATH, just run:

`go get -u https://github.com/maltegrosse/go-modemmanager`

## Usage

You can find some examples in the [examples](examples) directory.

## License


- **[MIT license](http://opensource.org/licenses/mit-license.php)**
- Copyright 2020 © Malte Grosse.