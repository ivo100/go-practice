## Goal: 

### Create a simple microservice API that can help us understand and track how a particular person's flight path may be queried. The API should accept a request that includes a list of flights, which are defined by a source and destination airport code. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.


Required JSON structure:
* [["SFO", "EWR"]]  => ["SFO", "EWR"]
* [["ATL", "EWR"], ["SFO", "ATL"]] => ["SFO", "EWR"]
* [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]


### Specifications:
* Your miscroservice must listen on port 8080 and expose the flight path tracker under the /calculate endpoint.

* Create a private GitHub repo and add https://github.com/taariq as a collaborator to the project. Please only add the collaborators when you are sure you are finished.

* Define and document the format of the API endpoint in the README.

* Use Golang and/or any tools that you think will help you best accomplish the task at hand.

* When you are done with the assignment, follow up and reply-all to the email that directed you to this document. Include your private github link and an estimate of how long you spent on the task and any interesting ideas you wish to share.

## Main assumption about the API endpoint and the input:

REST using POST to http://localhost:8080/calculate
JSON payload accoding to above structure

Output of the service is not the whole flight path but only start and end points.

The input is for a single traveller / single flight path. Time order, flight numbers, state, multiple travelers etc. are out of scope.
"Return" flights are out of scope.
There are no cycles (code should detect cycles and treat them as input error)

Implementation handles some trivial edge cases like 2 fligths A->B, B->A as single flight A->B. 
General cycles A->B->C->A cannot be resolved as any point on a circle can be start and end point at the same time.  
Additional information like time or point of origin is required in such cases.

# Implementation 

Because I had sufficient time I am providing 3 different implementations:

1. One using graph library github.com/dominikbraun/graph
2. One using simplified algorithm with hash table using the fact that the flight path is not general DAG but more of "linear" source->destination chain. Detection of origin point can be done by simply finding the node without incoming flights (in-degree = 0).
3. General Kahn's algorithm. 

Dominik Braun's library is very dood - it uses generics and some graph optimizations like transitive reduction.

Which implementation to use is controlled via environment variable exported before running the service.

The default (if not specified) is graphlib.

One additional env.var (DEBUG) controls logging level of the service - default is INFO but it can be set to DEBUG via export DEBUG=TRUE
There is e2e.sh script which builds and runs the service and executes some curl sample commands.

Each package has unit tests - from the rest service on the top to graph with 3 separate packages for different implementations - glib, kahn and simple.

There is one additional health check endpoint which is GET /health which returns 200 OK and version of the service.

/calculate endpoint detects cycles and returns 400 invalid responses for any cycle beyond trivial a->b->a case.

Additionally it validates that codes are exactly 3 non-blank characters, makes them UPPERCASE, removes whitespace, doesn't allow a->a etc.

The service is too simple for swagger file.

Snippet from e2e.sh

```shell


export DEBUG=TRUE
export USE_GRAPH=graphlib
#export USE_GRAPH=simple
#export USE_GRAPH=kahn

curl  -X  POST http://localhost:8080/calculate \
      -H  "Content-Type: application/json" \
      -d  '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]'

sample output

2023/04/10 16:09:52 Using graph library
{"time":"2023-04-10T16:09:52.199123-07:00","id":"","remote_ip":"127.0.0.1","host":"localhost:8080","method":"POST","uri":"/calculate","user_agent":"curl/7.87.0","status":200,"error":"","latency":46167,"latency_human":"46.167µs","bytes_in":64,"bytes_out":14}
["SFO","EWR"]

```
