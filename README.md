# Ticket Engine

[ ![Codeship Status for TimHugh/ticket_service](https://app.codeship.com/projects/4c23f2b0-02fd-0136-93ef-3244a1cc3029/status?branch=master)](https://app.codeship.com/projects/280193)
[![Maintainability](https://api.codeclimate.com/v1/badges/84dd612c3de4bb20e86d/maintainability)](https://codeclimate.com/github/timhugh/ticket_service/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/84dd612c3de4bb20e86d/test_coverage)](https://codeclimate.com/github/timhugh/ticket_service/test_coverage)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/ef9c7c104e30475ab75a09a679a25b3b)](https://www.codacy.com/app/theuett/ticket_service?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=timhugh/ticket_service&amp;utm_campaign=Badge_Grade)

## Usage

I use [govend](https://github.com/govend/govend) to manage dependencies:

```
$ go get -u github.com/govend/govend
$ govend -v
```

There are some required environment variables:
```
MONGODB_URI=mongodb://your-host/database

ROLLBAR_TOKEN=server_token

NEWRELIC_TOKEN=license_key
NEWRELIC_APP=Your App Name
```

Finally, it's easiest to run everything in Docker, using the included Dockerfile and docker-compose.yml. When I'm developing, I usually do something like this:

```
$ docker-compose up -d mongo
$ docker-compose build web
$ docker-compose up web
```
