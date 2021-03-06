# demo02: Hello Service
>用 `micro` 命令生成相应的文件

这个 README.md 也是自动生成的

Generated with

```
micro new --namespace=go.micro --type=service hello
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.hello
- Type: service
- Alias: hello

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./hello-service
```

Build a docker image
```
make docker
```