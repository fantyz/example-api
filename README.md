![main](https://github.com/fantyz/example-api/actions/workflows/ci.yml/badge.svg)

# Example API

The purpose of this repository is to provide an example of an as-complete-as-possible basic API written in Go.

It contains the following parts:

* Basic API with:
  - Config
  - Logging (optionally structured)
* Unit tests
* Integration tests
* Containerization
* Working CICD that:
  - Lints
  - Vets
  - Run unit tests
  - Run integration tests

## Running tests

Regular unit tests and integration tests have been separated for the sake of speed. Running integration tests require that you got a working Postgres DB that can be used. Often they also takes a bit longer to run than regular unit tests.

Regular unit tests are run the usual way: `go test .`. Integration tests (along with the regular unit tests) can be run with: `go test . -tags=integration`.

## Containerization

The example api is containerized using docker and relies on vendoring dependencies. You can build and run it by:

1. Run `go mod vendor`
1. Run `docker build -t example-api .`
1. Run `docker run --rm -p 8000:8000 example-api`

## TODO

Additional features that I'd like to include:

* Example endpoint that actually use a database + configure and connect. It's both akward having an integration test when the example itself doesn't use a database and it makes the copy-paste exercise to use this for something real easier.
* Configuration file that works with docker image as well
* Datadog integration (APM+tracing)
* Use a test suite for integration tests to only connect once and allow easier setup/teardown of the individual tests
* Improve the integration test setup to work for local development as well (grab config from environment/use defaults that are intended for local)
* Add example of using fixtures for integration tests
* Branch based deployment of service to <somewhere> (need infrastructure for this - example-infra repo to examplify how to work with infrastructure as well?)
