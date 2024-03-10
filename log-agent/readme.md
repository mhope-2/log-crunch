# Log Agent
Log Agent is just what it's name says; a log agent. It simulates the generation of log messages and then writes the
generated message to the Kafka brokers.

### Useful Kafka Shell Commands

1. Create a Topic:  
```bash
kafka-topics.sh --create --bootstrap-server <bootstrap-server> --replication-factor <replication-factor> --partitions <num-partitions> --topic <topic-name>
```
E.g.
```bash
kafka-topics.sh --create --topic topic --bootstrap-server kafka1:9092 --replication-factor 3 --partitions 1
```
2. List Available Topics:  
```bash
kafka-topics.sh --list --bootstrap-server <bootstrap-server>
```
E.g. 
```bash
kafka-topics.sh --list --bootstrap-server kafka1:9092
```

3. Console Consume / List Topic Data:
```bash
kafka-console-consumer.sh --bootstrap-server <bootstrap-server> --topic my-topic --from-beginning
```
E.g.
```bash
kafka-console-consumer.sh --bootstrap-server kafka1:9092 --topic topic --from-beginning
```

4. Delete a Topic
```bash
kafka-topics.sh --delete --topic <topic-name> --bootstrap-server <bootstrap-server>
```
E.g. 
```bash
kafka-topics.sh --delete --topic pic --bootstrap-server kafka1:9092
```
