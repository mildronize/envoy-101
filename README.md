# Enjoy Playground

https://www.youtube.com/watch?v=40gKzHQWgP0


Terminology: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/intro/terminology
connection_pooling: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/connection_pooling#arch-overview-conn-pool


## How to run

1. Run `docker-compose up` to start the services

Go to `http://localhost:8080/` to see the application running.

The result will be swapped between the two services.

```
App ID: 1111, Home Page
```

and

```
App ID: 2222, Home Page
```