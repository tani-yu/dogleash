# DOGLEASH

dogleash is a command line tool to import and export data from [Datadog](https://www.datadoghq.com/).

Maybe other funcion will be added.

## INSTALL

To build this tool from source code, Go 1.11 is required.

```
export GO111MODULE=on && go get github.com/tani-yu/dogleash
```

## SETTING

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

If you have installed dogleash manually, you need add the completion settings to bash completion directory.

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

## USAGE

get all monitor information at standard output. (json format)

```
dogleash monitor show_all
```

if you want to save the output result, you should add `-p` and specify a location to path.

```
dogleash monitor show_all -p /tmp/
```

import monitor information from json file.

```
dogleash monitor import -i JSON_FILE_PATH
```
