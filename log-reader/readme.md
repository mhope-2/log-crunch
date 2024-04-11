# Log Agent
Log Agent is just what it's name says; a log agent. It simulates the generation of log messages and then writes the
generated message to the Kafka brokers.

### Useful Kafka Shell Commands

1. List Available Topics:
```bash
kafka-topics.sh --list --bootstrap-server <bootstrap-server>
```
E.g.
```bash
kafka-topics.sh --list --bootstrap-server kafka1:9092
```

2. Console Consume / List Topic Data:
```bash
kafka-console-consumer.sh --bootstrap-server <bootstrap-server> --topic my-topic --from-beginning
```
E.g.
```bash
kafka-console-consumer.sh --bootstrap-server kafka1:9092 --topic topic --from-beginning
```

### Useful Cassandra Shell Commands
1. Start the shell
```bash
cqlsh
```

2. Create a Keyspace (Database)
```bash
CREATE KEYSPACE IF NOT EXISTS my_keyspace WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '3'}  AND durable_writes = true;
```