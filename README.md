# Bocker

Implementing Linux containers

## Prequisite

- golang
- docker(recommended)
- debootstrap(use your package manager to install it)

## Building

Prefer running it inside docker container as it is under development stage

```bash
docker build -t bocker .
```

or

```bash
go build -o bocker app/main.go
```

## Running

Commands will be running inside a sandboxed environment which will not affect your host system.

```bash
docker run --rm -it --privileged bocker <ARGUMENTS>
```

or

```bash
./bocker run <ARGUMENTS>
```

## Example

Installing cowsay in bocker and running

```bash
./bocker run apt-get install cowsay
./bocker run cowsay Hello, Bocker!
```
