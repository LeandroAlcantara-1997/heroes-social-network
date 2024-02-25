# Heroes Social Network

<img align="right" width="220px" src="docs/assets/ironman.png">

## Sobre


Heroes social network é um projeto criado com a finalidade de facilitar a vida de fãs de heróis, no qual oferece recursos como:

1. Busca de personagens;
2. Busca de equipes no qual os personagens pertencem;
3. Dados de HQ's;
4. Dados de Jogos, Filmes e Séries;


## Aqui você vai encontrar dados de personagens de empresas como:

<img align="center" width="140px" src="docs/assets/marvel.png">
<img align="center" width="80px" src="docs/assets/DC_Comics_logo.png">



## Como executar

1. Suba as dependências com o comando:
~~~ make 
make docker-up
~~~

2. Configure as tabelas no banco de dados com o seguinte comando:
~~~
make migration-up
~~~

3. Após subir as dependências, inicialize a aplicação com o comando:
~~~
make run
~~~

## Para configurar Splunk

Para configurar [Splunk](http://localhost:8000/), acesse a plataforma em settings -> data inputs -> HTTP Event Collector -> Global Setting e desabilite o SSL.

|[ENVIRONMENTS](build/.env.example)                  | EXAMPLE VALUES                           |
|----------------------------|------------------------------------------|
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
|ALLOW_ORIGINS               |http://localhost:3000,https:/localhost:8080|
|OTEL_EXPORTER_OTLP_ENDPOINT |http://jaeger:4318                        |
|ENVIRONMENT                 |local                                     |
|SERVICE_NAME                |heroes-social-network                     |


#### Para acessar os contratos da API, execute a aplicação e acesse http://localhost:8080/swagger/index.html#/, ou importe o arquivo [swagger.yaml](/docs/swagger.yaml) no [swagger editor](https://editor.swagger.io/) ou importe o arquivo [Heroes-social-network.postman_collection.json](/docs/heroes-social-network.postman_collection.json) em seu Postman. Para mais informações sobre o projeto, visite: [Wiki](https://github.com/LeandroAlcantara-1997/heroes-social-network/wiki)


## Tecnologias utilizadas:

* Golang 1.20;
* PostgreSQL;
* Redis as cache;
* Splunk;
* Trace com Jaeger;
* OpenTelemetry;
* Swagger;
* Graceful Shutdown;
* Dev Container com Docker;


## Diagrama de entidade relacionamento:

![diagram](/docs/assets/heroes-social-network.jpg)