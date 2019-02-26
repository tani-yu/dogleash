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

### Enabling shell command completions

DogLeash supports shell command completions for Bash and Zsh.

#### On macOS, Using Bash

If you are using macOS and Bash, you need to install `bash-completion` using [Homebrew](https://brew.sh/).

```bash
## If running Bash 3.2 included with macOS
brew install bash-completion

## or, if running Bash 4.1+
brew install bash-completion@2
```

Follow the "caveats" section of brew's output to add the appropriate bash completion path to your local `~/.bashrc`.

If you have installed DogLeash manually, you need add the completion settings to bash completion directory.

```bash
dogleash completion bash > $(brew --prefix)/etc/bash_completion.d/dogleash
```

#### Using Zsh

If you are using Zsh, you can enable autocompletion adding the following codes to `~/.zshrc`.

```zsh
if [ $commands[dogleash] ]; then
  source <(dogleash completion zsh)
fi
```

## Usage

Get all monitor information at standard output. (JSON format)

```bash
dogleash monitor show_all
```

Display all monitor information in JSON format.

```bash
dogleash monitor show_all
```

Export all monitor information in JSON File. If you want to specify path, you can use `d` option.

```bash
dogleash monitor export
dogleash monitor export -d /tmp/
```

Import monitor information from JSON file.

```bash
dogleash monitor import -i JSON_FILE_PATH
```

## Contributing

See the [contributing guidelines](CONTRIBUTING.md).

## License

[Apache-2.0](LICENSE)
