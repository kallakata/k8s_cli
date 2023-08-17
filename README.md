# A Kubernetes CLI parser with pretty output 

A full fledged Kubernetes CLI. As of now, it can parse _Pods, Namespaces, Clusters_ and _metadata._ Output can be customized, default is with a prompt and table.

**More features TBD.**

## TODO ##
- _Deployment_
- _Nodes_/_Nodepools_
  
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

### Make ###
```console
$ make build
$ make run [COMMANDS] [FLAGS]
$ make run --help
```

### If you want to use new packages ###
```console
$ make vendor
```

