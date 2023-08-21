# A Kubernetes CLI parser with pretty output 

A full fledged Kubernetes CLI. As of now, it can parse _Pods, Namespaces, Clusters_ and _metadata._ Output can be customized, default is with a prompt and table.

Please refer to the official Google docs for more extensive metadata parsing and response properties..

**More features TBD.**

## TODO ##
- _Deployment_
  
## Usage ##

### Manual ###
```console
$ go build . && go install
$ ./app
$ ./app [COMMANDS] [FLAGS]
$ ./app --help
```

### Install binary ###
```console
$ go install .
```

### Build ###
```console
$ make build
$ cd ./bin
$ ./kubeparser [ARGUMENTS] [FLAGS]
```

### Run ###
```console
$ make run [COMMANDS] [FLAGS]
$ make run --help
```

### If you want to use new packages ###
```console
$ make vendor
```

### Test & format ###
```console
$ make fmt
$ make test
```
