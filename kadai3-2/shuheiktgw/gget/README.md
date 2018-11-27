gget
====

gget is a wget like command to download file, but downloads a file in parallel.

## Usage
```
gget [options...] URL

OPTIONS:
  --parallel value, -p value  specifies the amount of parallelism (default: the number of CPU)
  --help, -h                  prints help

```

## Install

```
go build
./gget [options...] URL
```