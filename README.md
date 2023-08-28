# A Kubernetes CLI parser with pretty output 

:wave: A full fledged **Kubernetes CLI**.

As of now, it can parse _Pods, Namespaces, Clusters_ (in table output), _Nodepools_ (in default list output), and partially most important _metadata._ Output can be customized, default is a simple listing, otherwise it is in a form of table + prompt.

Please refer to the official Google docs for more extensive metadata parsing and response properties related to clusters/nodepools. Authentication is handled locally via _kubeconfig_ and _gcloud_.

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
