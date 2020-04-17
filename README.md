![Alt Go-ModemManager](./go-modemmanager.png)
================
[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)
![Go](https://github.com/maltegrosse/go-modemmanager/workflows/Go/badge.svg) 
[![Go Report Card](https://goreportcard.com/badge/github.com/maltegrosse/go-modemmanager)](https://goreportcard.com/report/github.com/maltegrosse/go-modemmanager)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/f873c5d0eb514347b01b6f24dd4f7b76)](https://www.codacy.com/manual/maltegrosse/go-modemmanager?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=maltegrosse/go-modemmanager&amp;utm_campaign=Badge_Grade)

Go D-Bus bindings for ModemManager


Additional information: [ModemManager D-Bus Specs](https://www.freedesktop.org/software/ModemManager/api/1.12.0/ref-dbus.html)

Tested with [ModemManager - Version 1.28.8](https://gitlab.freedesktop.org/mobile-broadband/ModemManager) and Go 1.13
with a [SolidRun Hummingboard Edge](https://www.solid-run.com/nxp-family/hummingboard/) on `Debian Buster (armv7)` with `Kernel 5.4.x` and `libqmi 1.24.6` and a `Quectel EC25` miniPcie modem.

## Status
Work in Progress, some methods/properties/signals needs to be fixed for initial release of version 0.1

## Todo
Some methods/properties are untested as they are not supported by my modem/lack of how to use them. See `todo` tags in the code.

- Implement MarshalJson methods
- double check struts
- implement signal methods
- tidy up the code

## Installation

This packages requires Go 1.13 (for the dbus lib). If you installed it and set up your GOPATH, just run:

`go get -u https://github.com/maltegrosse/go-modemmanager`

## Usage

You can find some examples in the [examples](examples) directory.

## License
**[MIT license](http://opensource.org/licenses/mit-license.php)**

Copyright 2020 © Malte Grosse.

Other:
- [ModemManager Logo under GPLv2+](https://gitlab.freedesktop.org/mobile-broadband/ModemManager/-/tree/master/data)

- [GoLang Logo under Creative Commons Attribution 3.0](https://blog.golang.org/go-brand)