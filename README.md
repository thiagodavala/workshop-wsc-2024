# Workshop Worldskills Brasil
#53 Cloud Computing - 2024

## How to Build

```
go mod init thiagoogeremias.io/workshopcloud2024
go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/dynamodb
go build -o server .
```

## How to execute

```
export DYNAMO_TABLE="tabelateste"
export PORT="80"
```

## Runbook

### Introduction

The Unicorn Rentals website provides cryptographically strong hashes as a service. The
system is comprised of a client application and a standalone server. The system will
always return the same hash value for a given key. The system builds the hash based on a
proprietary, computationally expensive hashing algorithm.

The server is currently deployed as a single, standalone 'go' binary. There is no
ession state to track or external calls to make so the server architecture should
respond well to simple horizontal scaling. As there is expected to be a constrained
set of client keys, the system should respond well to caching techniques.
That's it for hints.


### Scope

This runbook describes the operational theory and practice for the production system
powering the Unicorn Rentals website. The primary audience is the DevOps team running
the site. The DevOps team is responsible for deploying code, scaling the site in response
to load, maintaining our published SLA's (including response time and uptime), disaster
recovery, troubleshooting activities and any monitoring and alerting activities required
to meet these objectives.

#### Tasks

1. Please read the documentation carefully, it will be very helpful
2. Access AWS Academy Leaner Lab, start a Lab.
3. Check the existing configuration in EC2 (Elastic Cloud Computer Server)
4. Check the existing configuration in VPC (VPC, Subnet)
5. Configure the application to automatically scale to handle the increasing load (Auto Scaling Groups, with Launch Configuration)
6. Configure the relevant server dependencies as described in the technical details
7. Your most important task is to launch and ensure that the Server application is available and scalable
8. Configure the necessary application monitoring, metrics and alarms in CloudWatch
9. Monitor costs and do not overscale the infrastructure to minimize penalty points
10. The amount of traffic changes over time and occurs continuously

### Troubleshooting Procedures

If you need help, try:

```
./server --help
```

New application test utility!

```
export BASE_URL=localhost:80
./test_requests.sh
```

## To Do

- How to test HA? Automate DR.
- Refactor infrastructure using container architecture
