# Kafka Service Broker

This service broker shares a large Apache Kafka/Apache ZooKeeper amongst many users via the Open Service Broker API.

It offers two service plans:

```
Service Name         Plan Name  Description
starkandwayne-kafka  shared     Create your own topics on shared Kafka
~                    topic      Share a single topic on shared Kafka
```

## Installation

You can install the `kafka-service-broker` CLI various ways:

* MacOS/Homebrew

    ```
    brew install starkandwayne/cf/kafka-service-broker
    ```

* Debian/Ubuntu

    ```
    wget -q -O - https://raw.githubusercontent.com/starkandwayne/homebrew-cf/master/public.key | apt-key add -
    echo "deb http://apt.starkandwayne.com stable main" | tee /etc/apt/sources.list.d/starkandwayne.list
    apt-get update
    apt-get install kafka-service-broker
    ```

* BOSH - see https://github.com/cloudfoundry-community/kafka-service-broker-boshrelease

* Golang from source

    ```
    go get -u github.com/starkandwayne/kafka-service-broker/cmd/broker
    mv $GOPATH/bin/{broker,kafka-service-broker}
    ```

## Run broker

```
kafka-service-broker run-broker
```

## Configuration

The following environment variables can be used to configure the broker:

* `PORT` is the broker listen port for HTTP traffic, defaults to `8100`
* `BROKER_USERNAME` and `BROKER_PASSWORD` are required to setup basic auth authorisation to the API
* `ZOOKEEPER_PEERS` - ZooKeeper cluster used to discover the current Kafka cluster; a comma separated list of `host1:port,host2:port,host3:port`, defaults to `localhost:2181`

## Catalog

The default service catalog is at `data/assets/catalog.json`.

You can override the entire JSON response by setting `$BROKER_CATALOG_JSON` to the JSON string.

Alternately you can make minor adjustments to some of the globally unique attributes:

* `BROKER_SERVICE_GUID` - to change the GUID of the service
* `BROKER_SERVICE_NAME` - to change the name of the service
* `BROKER_PLAN0_GUID`, `BROKER_PLAN1_GUID` - to change the GUID of the service plan (first, second, etc)

## Development

To only clone this branch:

```
git clone https://github.com/starkandwayne/kafka-service-broker --single-branch
```

To run from source code:

```
go run cmd/broker/main.go run-broker
```

When adding/updating `data/assets/`, remember to run `go-bindata` to embed the changes into `data/data.go`:

```
go-bindata --pkg data -o data/data.go data/assets/...
```
