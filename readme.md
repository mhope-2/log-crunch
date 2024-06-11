# Log Crunch

Log Crunch is a project comprising three different go services; a log agent, a log reader/processor and a log query service.
The `log-agent` service generates random log messages via a Kafka Producer, the `log-reader` service then consumes the log messages
using a Kafka Consumer and proceeds to write it to a Cassandra DB.
The log-query service provides an API interface to query the log data.


## How to set up

---
### With Docker
1. Clone the repository
```bash
git clone https://github.com/mhope-2/simple-commerce.git
```
2. Build services
```bash
docker-compose build
```
3. Startup services using docker
```bash
docker-compose up
``` 

### With Make
1. Clone the repository
```bash
git clone https://github.com/mhope-2/simple-commerce.git
```
2. Build services
```bash
make build
```
3. Startup services using docker
```bash
make up
``` 

### Sample GET request to the log-query service
```bash
curl http://localhost:5003/messages/3f1c084a-3227-4a95-a6ab-2d868e5bea02
```
where `3f1c084a-3227-4a95-a6ab-2d868e5bea02` is the message id from the Cassandra DB

sample response:

```json
{
    "ID": "3f1c084a-3227-4a95-a6ab-2d868e5bea02",
    "LogLevel": "INFO",
    "Value": "Failed to connect to external API",
    "Timestamp": "2024-06-11T10:14:41Z"
}
```

### Note
- Read the readme files in each of the service dirs to find out how to interact with the services in a Docker env. 

### Todo
[x] Ensure reliability and resilience of kafka brokers
