# Hero Social Network

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

2. Após subir as dependências, inicialize a aplicação com o comando:
~~~
make run
~~~

#### Para acessar os contratos da API, execute a aplicação e acesse localhost:port/swagger/index.html#/, ou acesse [swagger editor](https://editor.swagger.io/) ou importe o arquivo Heroes-social-network.postman_collection.json em seu Postman. Para mais informações sobre o projeto, visite: [Wiki](https://github.com/LeandroAlcantara-1997/heroes-social-network/wiki)


## Tecnologias utilizadas:

* Golang 1.20;
* PostgreSQL;
* Cache;
* Splunk;


## Diagrama de entidade relacionamento:

![diagram](/docs/assets/heroes-social-network.jpg)