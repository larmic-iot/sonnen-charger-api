# sonnenBatterie-api

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go build](https://github.com/larmic-iot/sonnen-charger-api/actions/workflows/go-build.yml/badge.svg)](https://github.com/larmic-iot/sonnen-charger-api/actions/workflows/go-build.yml)
[![Docker build and push](https://github.com/larmic-iot/sonnen-charger-api/actions/workflows/docker-build-and-push.yml/badge.svg)](https://github.com/larmic-iot/sonnen-charger-api/actions/workflows/docker-build-and-push.yml)
[![Docker hub image](https://img.shields.io/docker/image-size/larmic/sonnen-charger-api?label=dockerhub)](https://hub.docker.com/repository/docker/larmic/sonnen-charger-api)
![Docker Image Version (latest by date)](https://img.shields.io/docker/v/larmic/sonnen-charger-api)

*NOTE* this application is still in development and a very early version.

A REST api client (adapter) for the [sonnenCharger](https://sonnen.de/ladestation-elektroauto/). Sonnen does not suppoert
the REST api of the charger. So this app is using 
[Modbus](https://en.wikipedia.org/wiki/Modbus#:~:text=Modbus%20TCP%2FIP%20or%20Modbus,layers%20already%20provide%20checksum%20protection)
REST endpoints documented in [open api 3.1](api/open-api-3.yaml).

This project inspired by [tp-link-hs110-api written in go](https://github.com/larmic/tp-link-hs110-api) and
improves my Go knowledge.

## Versioning

[Semantic Versioning 2.x](https://semver.org/) is used. Version number **MAJOR.MINOR.PATCH** with

* **MAJOR** version increase on incompatible API changes
* **MINOR** version increase on adding new functionality in a backwards compatible manner
* **PATCH** version increase on backwards compatible bug fixes or documentation

## Usage

The easiest way is to use the docker image. Otherwise, the artifact will have to be built by yourself.

```sh 
$ docker pull larmic/sonnen-charger-api
$ docker run -d -p 8080:8080 --rm \
 -e SONNEN_CHARGER_IP='<my-charger-ip>' \
 --name larmic-sonnen-charger-api larmic/sonnen-charger-api
```

## Example requests

See [open api 3 specification](api/open-api-3.yaml) for further information.

```sh 
$ curl http://localhost:8080/api             # Open Api 3.1 specification
$ curl http://localhost:8080/api/settings    # Charger settings
```

## Build application by yourself

### Requirements

* Docker
* Go 1.21.x (if you want to build it without using docker builder)

### Build and run it

See [Makefile](Makefile)!

```sh 
$ make              # prints available make goals
```