# DogLeash

[![Build Status](https://travis-ci.com/tani-yu/dogleash.svg?token=gx8YmzzZyXuG4grwEWXa&branch=master)](https://travis-ci.com/tani-yu/dogleash)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/tani-yu/dogleash/blob/master/LICENSE)

DogLeash is a command line tool to import and export data from [Datadog](https://www.datadoghq.com/).

Maybe other funcion will be added.

## Install

To build this tool from source code, Go 1.11 is required.

```bash
export GO111MODULE=on && go get github.com/tani-yu/dogleash
```

## Settings

Go to [Datadog API settings](https://app.datadoghq.com/account/settings#api) and generate an API key and Application key.
Create a minimal `~/.dogrc` that looks like this:

```ini
[Connection]
apikey = YOUR_API_KEY
appkey = YOUR_APP_KEY
```

You can skip this step if you already use [dogshell](https://docs.datadoghq.com/developers/faq/dogshell-quickly-use-datadog-s-api-from-terminal-shell/) and have `~/.dogrc` file.

You can also use environment variables `DATADOG_API_KEY` and `DATADOG_APP_KEY`.
In this case, the credential read from `.dogrc` file will be overwritten by the environment variables.


## Usage

get all monitor information at standard output. (json format)

```bash
dogleash monitor show_all
```

if you want to save the output result, you should add `-p` and specify a location to path.

```bash
dogleash monitor show_all -p /tmp/
```

import monitor information from json file.

```bash
dogleash monitor import -i JSON_FILE_PATH
```

## Contributing

See the [contributing guidelines](CONTRIBUTING.md).

## License

[Apache-2.0](LICENSE)
