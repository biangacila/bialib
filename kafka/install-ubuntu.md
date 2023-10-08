

wget https://downloads.apache.org/kafka/<version>/kafka_2.13-<version>.tgz

wget https://downloads.apache.org/kafka/3.5.1/kafka_2.13-2.8.0.tgz

tar -xzf kafka-3.5.0-src.tgz

cd kafka_2.13-<version>


bin/kafka-server-start.sh config/server.properties

============ docker =========================================

docker pull confluentinc/cp-kafka

docker run -d --name kafka -p 9092:9092 --network host -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 confluentinc/cp-kafka

docker run -d --name kafka -p 9092:9092  -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 confluentinc/cp-kafka

sudo apt install kafka-client


======== install kalka from docker-compose ===================

docker-compose up -d
docker-compose ps
docker-compose down



Golang-based message broker (e.g., NSQ) for real-time event processing and streaming.


Implement Querying and Projections:

Build query mechanisms to retrieve event streams and aggregate data from events. You can use query patterns like CQRS (Command Query Responsibility Segregation).
Create projections or views that store the current state of your application based on events. This can help in faster data retrieval.