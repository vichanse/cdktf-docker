# cdktf-docker

This repository allows you to deploy applications running in Docker containers, with CDK for Terraform (CDKTF).

- a backend, in Go, that display a whale ^^
- a frontend, in NodeJS, that call the backend and ... display the whale :)

`backend` and `frontend` containers will run in a network named `my_network`.

## Prerequisites

- Install cdktf CLI

## Docker

Deploy our stack:

```bash
$ cdktf get
$ cdktf deploy
```

Check everything is corerctly deployed:

```bash
$ docker image ls
$ docker container ls
$ docker network ls
$ curl localhost:8000
```

Destroy your stack:

```bash
$ cdktf destroy
```
