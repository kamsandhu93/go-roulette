# Go Roulette

* [Installation](#installation)
  - [Pre-reqs](#pre-reqs)
  - [Install Dependencies](#install-dependencies)
  - [Run the webserver on localhost:8080](#run-the-webserver-on-localhost-8080)
  - [Run tests](#run-tests)
* [API](#api)
  + [Healthcheck](#healthcheck)
  + [Roulette](#roulette)
* [Roadmap](#roadmap)

Go implementation of a backend web server for the popular casino game Roulette. 

Technical stack:
- Go - A modern statically typed programming language, that can be compiled into small portable binaries. 
It has a simple C like syntax, with an inbuilt garbage collector, and native support for concurrency - which makes it a good candidate for microservices. 
- Gin - A webserver framework witten in Go. It provides a fast, highly performant web server with a number of valuable features. These include JSON validation
and middleware support. Using Gin reduces the amount of boilerplate code required to get a reliable and scalable webserver
up and running in Go. This allows developers to focus on implementing business cases, ultimately delivering more value to the
end user. Gin is a popular Go framework, and therefore has extensive community support.



## Installation
#### Pre-reqs
- Make ([Mac](https://formulae.brew.sh/formula/make), [Ubuntu](https://linuxhint.com/install-make-ubuntu/))
- [Go version 1.17](https://go.dev/dl/) 


#### Install Dependencies
```bash
make install
```
#### Run the webserver on localhost:8080
```bash
make run 
```
#### Run tests
```bash
make tests
```

## API

The server exposes a RESTful API that accepts JSON http requests. The following endpoints are supported:

### Healthcheck
A simple health check endpoint.

`/heatlh`

####Methods:
- GET

####Request fields: 
None

####Response
```json
"OK"
```
### Roulette
Runs a round of roulette for bets that are submitted to this endpoint. Winnings are calculated and returned to 
the caller as a JSON payload.

`/v1/roulette`

####Methods:
- POST

####Request fields:
- (string) correlation_id - Unique identifier to correlate traffic across client and server.
- (string) user_id - Unique identifier for the user.
- (array) bets - Array of bet objects with the following schema:
    - (string) id - Unique identifier for the bet
    - (string) BetType - Type of bet placed, see bellow for valid values.
    - (int) Size - Value of the bet placed - a decimal number. 
 
####Example Request:
```json
{
    "user_id" : "1",
    "correlation_id" : "1",
    "bets": [
        {
            "id": "1",
            "size" : 1,
            "type" : "6"
        },
        {
            "id" : "2",
            "size" : 5,
            "type" : "even"
        },
        {
            "id" : "3",
            "size" : 10,
            "type" : "black"
        }
    ]
}
```

####Response fields:
- (string) correlation_id - Unique identifier to correlate traffic across client and server.
- (int) winning_number - Winning number from the round of roulette.
- (int) winnings - Total winnings derived from submitted bets.

####Example response:
```json
{
    "correlation_id" : "1",
    "winning_number" : 10,
    "winnings" : 0
}
```

####Supported bet types:

Bet types represent the location of a chip placed on a roulette board. The following table details all valid bet types.

| Bet Type    | Valid Values (string)                                                                                     | Implemented           |
|-------------|-----------------------------------------------------------------------------------------------------------|-----------------------|
| Straight    | <code>0&#124;1&#124;2&#124;3&#124;4&#124;5&#124;6&#124;7...34&#124;35&#124;36</code>                      | :white_check_mark:    |
| Half Board  | <code>odd&#124;even&#124;high&#124;low&#124;red&#124;black</code>                                         | :white_check_mark:    |
| Third Board | <code>first-third&#124;second-third&#124;third-third&#124;first-col&#124;second-col&#124;third-col</code> | :x:                   |
| Split       | <code>1-2&#124;2-3...</code>                                                                              | :x:                   |
| Street      | <code>1-2-3&#124;4-5-6...</code>                                                                          | :x:                   |
| Square      | <code>1-2-4-5&#124;2-3-5-6...</code>                                                                      | :wavy_dash: Partially |

## Roadmap

- Finish the square bet implementation.
- Implement remaining bet types documented in the table above. 
- Add supplemental unit tests to increase coverage, and also more granular unit tests that test at the function level.
- Create a Dockerfile for a container that runs the Gin server. The dockerfile should use a builder layer to build the binary.
- Generate a UUID request ID and return it with the header on all responses.
- Deploy the built docker images to a hosting service such as AWS fargate, with a https load balancer placed in front of it such as AWS ALB. 
The load balancer should terminate the SSL connection and forward to the docker container.
- Implement CI/CD steps that run jobs on feature branches (on push to remote, or merge requests):
  - Static analysis such as unit tests, coverage, security analysis.
  - Go/Docker build.
- Implement CI/CD steps that run jobs on the main branch (on feature branch merges into main):
  - Build artefacts - Docker image.
  - Run tests against the new docker image.
  - Promote to docker repository (such as ECR), with an appropriate tag such as the current git sha.
  - Deployment of newly built image to hosting environment.
  - Run environment based integration/end-to-end tests.

