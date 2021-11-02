# Booking request manager

## Requirements

In order to execute this application in your machine you'll need to have the following software installed:

- Go 1.17 or higher https://golang.org/doc/install (only if you want to execute the application or tests in local)
- Docker: https://docs.docker.com/get-docker/
- Docker-compose: https://docs.docker.com/compose/install/

## Init the application

In order to have all the dockerized environment in local in order to execute, test and benchmark the application:
```
make init
```

## Run the application

In order to execute the booking-request-manager app you'll need to execute the makefile command:
```
make server-run
```

## Execute tests:
```
make tests
```

## Execute benchmark:
```
make benchmark
```
The result of the benchmark execution will be stored in a bench.new file inside the root project. So the idea is to
execute the benchmark with the current code, rename the bench.new to for example bench.old, change the code and run again the benchmark to compare the two files.

If you want to compare the two benchmarks you can install this tool locally (you'll need to have go installed locally):

```
go get -u golang.org/x/perf/cmd/benchstat
```
After doing this you can execute the following tool to see the performance gain of the changes:
```
benchstat bench.old bench.new
```


If you use the intellij IDEA (intellij ultimate or only goland) you can execute both tests and the booking-request-manager application through the run configurations stored in the .run folder.
This is configured to use the go installed in local machine, so you'll need to have go installed

## Architecture overview:
- This application was developed using the hexagonal architecture tactical approach of the Domain Driven Design.

## Folder structure:
- The entry point of the application lives in the cmd/api folder


- All the code of the application lives in the internal folder, It's separated by modules (booking folder) and inside each module we can find this structure:
    - domain: Here we find all the value objects, and we will place the entities, repositories and domain services if needed.
      Here we have all the domain logic related to guard the consistency of our models (bookingRequests, stats...)


- application: Here we find the command handlers (use cases)


- infrastructure/persistence: Here we'll find the repository implementations, in this case the persistence layer is implemented using mongodb


- infrastructure/io: Here we place all the specific ways to expose our application layer (command handlers). Now as we're exposing the "stats and maximize command handlers" using a http rest server we can find the following controllers:
    - infrastructure/io/http/rest/controller/health_check
    - infrastructure/io/http/rest/controller/stats
    - infrastructure/io/http/rest/controller/maximize

##TODO

- Improve the dockerization in order not to need to install go in local (I leave this as improvement because having go in local is very straightforward)
- Improve the approach to maximize stats and run the benchmark in order to see the gain