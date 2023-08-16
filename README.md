# A Kubernetes CLI parser with pretty output 

A full fledged Kubernetes CLI. As of now, it can parse Pods, Namespaces, Clusters and metadata. Output can be customized.

**More features TBD.**

## TODO ##
- _Deployment_
- _Nodes_
  
## Usage ##

### Manual ###
```console
$ go build . && go install
$ ./app
$ ./app [COMMANDS] [FLAGS]
$ ./app --help
```

### Make ###
```console
$ make build
$ make install
$ make run [COMMANDS] [FLAGS]
$ make run --help
```

### If you want to use new packages ###
```console
$ make vendor
```

