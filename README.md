# Data Dashboard

## About

This program is meant to work in combination with the [SubStore](https://github.com/GerhardusC/SubStore) program. It allows you to visualise/retrieve the data captured from an MQTT broker via a REST api.

## Installation

Currently, if you are running an x86 machine on linux, you may choose to install this program using the dedicated [MQTTui](https://github.com/GerhardusC/MQTTui) project, otherwise you must build from source.

Ensure you have the rust toolchain installed and run:
`cargo build --release`

You will find the executable output as `./target/release/substore`

## Usage

-d string
      Database path (default "./data.db")
-s string
      Serve directory, the directory served as the frontend. (default "frontend/dist/")

## API Docs

### GET -> `/measurements/since?timestamp=x`

Retrieve all measurements after the timestamp "x" (UNIX epoch).

### GET -> `/measurements/between?start=x&stop=y`

Retrieve all measurements between the start "x" timestamp and stop "y" timestamp (UNIX epoch).

### GET -> `/measurements`

Retrieve all measurements. Careful, this can read a large amount of data.

