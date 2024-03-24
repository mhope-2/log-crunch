# Log Agent
Log Agent is just what it's name says; a log agent. It simulates the generation of log messages and then writes the
generated message to the Kafka brokers.

### Useful Kafka Shell Commands


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
