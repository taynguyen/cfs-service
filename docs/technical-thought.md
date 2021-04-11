# Goal
This seem to be a simple backend service that would would focus on:
 - Nice, simple design
 - Performance
 - Scalability
 - Easy to maintain

# Steps
 - Models + Chose DB
 - Programing language and framework
 - CI, CD
 - Deployment: ?? Should be docker for easy rolling up, down, adding node or working with K8s when perform
 - Test plan:

## Assumption
 - We already have user/authentication service with predefined authorization JWT token with embedded data
```JSON
{
	"u-id":	"uuid", // User id
	"a-id":	"uuid", // Agency id that user belonged
	"e":  "", // Expired at mls
	"r":	"", // Role id if existed.
}
 - Our service would be provided a public key (from key-pair that used by user service) to validate the JWT token.
 - Then we don't 
```

 - Our request 

# Data model
## DB 
 Chose our DB: Consider: 
  - Relational: Postgres, MySQL
	- Non relational: MongoDB
NoSQL is designed for flexibility and speed, not 100% data integrity. Our models is defined and clear, with some business logic related shown by model, not the dynamic data or data that have unclear structure. 
SQL is also proven technology with good developer experience and support.
=> Let's go with SQL

Most popular DB: Postgres and MySQL
- MySQL: extremely fast database for read-heavy workloads, sometimes at the cost of concurrency when mixed with write operations.
- Postgres: performance was more balanced - reads were generally slower than MySQL, but it was capable of writing large amounts of data more efficiently, and it handled concurrency better
With new release, both is improved, and having good performance
=> Performance should not the factor to choose

Both is good, Postgrest is better for writing large data, MySQL is better or reading large data. I also have more experience with MySQL. I think we go with MySQL
=> MySQL

## Model:
 - Would designed this as separated service with user or agency. That would reduce the complexity of the service. And it's only for the events, don't need to care about user, permission, or agency management.

## Storing datetime
 - Storing as type datetime or UNIX: prefer UNIX to use the supported datetime function in SQL, make our query life easier.
 - Store as UTC, NO local datetime stored

# Programming language and framework
 I have 2 choice: golang or node 
 Based on the requirement, our system would need high load of APIs call (reading, writing). Nodejs using event loop, coding is fun, but the concurrent is rather not good. `Go` is designed with concurrent is the strong feature, it has `gorountines` and `channel` really suite with APIs server. 
 => go
 
 Web framework: Goal for a framework: minimal, support good set of middleware, don't need to serve UI, have good performance. After looking around, I saw that `fiber` fit my need, it have good performance and memory allocation.
 => Let's bootstrap with `fiber`

# CICD
 - We follow git-flow standard and Merge request when working with team.
 - Also add linter-check, error check in CI

# Deployment
 2 approaches:
  - Normal binary build, deploy directly on Ubuntu machine. => Simple, a bit better performance for single, small service.
	- Dockerize: using docker image to deploy. => Would take a bit more resource when using docker, but we good ecosystem, easy to rollup, rollout, support k8s, support zero downtime.
	- We could put our service behind a proxy like nghix or ingress (k8s), that would take care of the SSL problem
=> Dockerize the deployment

# Monitoring
 - We need to monitor our system, adding a `healthz` to support 3rd check system
 - Other system health could be implement in deploy machine an set up monitoring using prometheus, grafana, ...

# Test plan
Acceptant criteria:
Features:
 - Adding events
 - Query events:
 - Security:
  + Only user could add
	+ Only get events of your own agency
	+ Only valid authorization could execute command
 - Performance:
	+ Response time
	+ Request per second
	+ Load test

Plan:
 - Support `test` environment for QA teams and for regression
 - Unit test when implementation on feature
 - Develop integration test that cover our core feature feature, execute on our `develop`, `test`, `prod` branch. Also cover load test, stress test.

 - Do we have QA team?

