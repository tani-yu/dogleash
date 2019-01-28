# DOGLEASH

dogleash is a command line tool to import and export data from [Datadog](https://www.datadoghq.com/).

Maybe other funcion will be added.

## INSTALL

```
go get github.com/tani-yu/dogleash
```

## SETTINGS


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
