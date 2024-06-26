# Heroes Social Network

[README in Portuguese](README_pt.md)
<img align="right" width="220px" src="docs/assets/ironman.png">

## About

Heroes social network is a project created to make life easier for superhero fans, offering resources such as:

1. Search for character characteristics;
2. Team to which characters belong;
3. HQ data;
4. Games, Movies and Series Data;


## Here you will find characters from companies such as:

<img align="center" width="140px" src="docs/assets/marvel.png">
<img align="center" width="80px" src="docs/assets/DC_Comics_logo.png">



## How to execute

1. Run:
~~~ make 
make docker-up
~~~

## To configure splunk
To configure [Splunk](http://localhost:8000/), access the platform in settings -> data inputs -> HTTP Event Collector -> Global Settings and disable SSL.

|[ENVIRONMENTS](build/.env.example)        | EXAMPLE VALUES                                   |
|--------------------|--------------------------------------------------|
|API_PORT                    | 8080                                     |
|DB_NAME                     | heroes                                   |
|DB_USER                     | user                                     |
|DB_PASSWORD                 |passw0rd                                  |
|DB_HOST                     |postgres-database                         |
|DB_PORT                     |5432                                      |
|SPLUNK_PASSWORD             |passw0rd                                  |
|SPLUNK_HOST                 |http://splunk:8088                        |
|SPLUNK_TOKEN                |e26c8b3f-fe44-490c-82f9-22c5bf8ab3b5      |
|SPLUNK_ASSYNC               |false                                     |
|CACHE_HOST                  |redis-cache                               |
|CACHE_PORT                  |6379                                      |
|CACHE_PASSWORD              |passw0rd                                  |
|CACHE_READ_TIMEOUT          |2                                         |
|CACHE_WRITE_TIMEOUT         |2                                         |
|ALLOW_ORIGINS               |http://localhost:3000,https://localhost:8080|
|OTEL_EXPORTER_OTLP_ENDPOINT |http://jaeger:4318                        |
|ENVIRONMENT                 |local                                     |
|SERVICE_NAME                |heroes-social-network                     |

#### To access the API contracts, run the application and access http://localhost:8080/swagger/index.html#/, or import the [swagger.yaml](/docs/swagger.yaml) file into the [swagger editor](https://editor.swagger.io/) or import the [Heroes-social-network.postman_collection.json](/docs/heroes-social-network.postman_collection.json) file into Postman .For more information about the project, visit: [Wiki](https://github.com/LeandroAlcantara-1997/heroes-social-network/wiki)


## Technologies

* Golang 1.20;
* PostgreSQL;
* Redis as cache;
* Splunk;
* Traces with Jaeger;
* OpenTelemetry;
* Swagger;
* Graceful shutdown;
* Dev Container with Docker;


## Entity relationship diagram

![diagram](/docs/assets/heroes-social-network.jpg)