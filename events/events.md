### setting up environment variables
```shell
export kafkaURL="127.0.0.1"
export topic=user_onboarding
export groupID=mongo-ingestion-group

export mongoURL="mongodb://127.0.0.1:27017"
export dbName=db
export collectionName=user

export MONGO_CONTAINER_NAME=mongo_db
export MONGO_HOST="127.0.0.1"
export MONGO_PORT=27017
export MONGO_INIT_DB_ROOT_USERNAME=user
export MONGO_INIT_DB_ROOT_PASSWORD=password
```


### run kafka server in using docker compose
```shell
git clone git@github.com:conduktor/kafka-stack-docker-compose.git
```

```shell
cd kafka-stack-docker-compose
```

```shell
docker-compose -f zk-single-kafka-single.yml up -d
``` 


### start producer
```shell
go run main.go 
```


### send message to producer - curl

```shell
curl --location 'http://localhost:8080/produce' \
--header 'Content-Type: application/json' \
--data '{"name": "sahil", "last_name": "kariyania"}'
```


### starting mongoDB using docker
```shell
docker pull mongo
```

```shell
docker run -d \
--name $MONGO_CONTAINER_NAME \
-p $MONGO_PORT:27017 \
-e MONGO_INIT_DB_ROOT_USERNAME=$MONGO_INIT_DB_ROOT_USERNAME \
-e MONGO_INIT_DB_ROOT_PASSWORD=$MONGO_INIT_DB_ROOT_PASSWORD \
mongo
```