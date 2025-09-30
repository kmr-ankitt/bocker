# Bocker

Implementing Linux containers

## Prequisite

- golang
- docker(recommended)

## Building

Prefer running it inside docker container as it is under development stage

```bash
docker build -t bocker .
```

or

```bash
go build app/main.go
```

## Running

```bash
docker run --rm -it --privileged bocker <ARGUMENTS>
```

or

```bash
./main <ARGUMENTS>
```
