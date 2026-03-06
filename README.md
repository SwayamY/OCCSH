# OneClickCloudSelfHeal

A small Go microservice designed to demonstrate basic service lifecycle management, observability, and failure handling in a containerized environment.

This project focuses on operational thinking rather than just writing API endpoints.

---

## Features

- Health endpoint (`/health`)
- Status endpoint (`/status`)
- Metrics endpoint (`/metrics`)
- Crash simulation endpoint (`/crash`)
- Request logging
- Request counting
- Service uptime tracking
- Dockerized deployment
- Environment-based configuration

---

## Endpoints

### Health Check

GET /health


Returns `OK` if the service is alive.

---

### Status

GET /status


Returns basic service status.

Example response:

```json
{
  "status": "running"
}```

---

### Metrics
GET /metrics

Returns service metrics.

Example:

{
  "uptime_seconds": 120,
  "request_count": 15
}

---

###Crash Simulation
GET /crash

Intentionally crashes the service to demonstrate failure handling and restart strategies.
----

#Running Locally

Start the service:

#go run main.go

Test endpoints:

#curl localhost:8080/health
#curl localhost:8080/status
#curl localhost:8080/metrics
Custom Port

You can run the service on a custom port:

#PORT=9000 go run main.go

---

Docker

Build image:

#docker build -t occsh .

Run container:

#docker run -p 8080:8080 occsh

 ---

Why This Project

The goal of this project is to demonstrate concepts useful for backend and infrastructure roles:
-service observability
-failure simulation
-containerized deployment
-runtime configuration
-structured logging
