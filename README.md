# Hero Social Network

[README in Português](README_pt.md)
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

1. Climb the dependencies
~~~ make 
make docker-up
~~~

2. After uploading the dependencies, run:
~~~
make run
~~~
#### To access the API contracts, run the application and access http://localhost:8080/swagger/index.html#/, or import the [swagger.yaml](/docs/swagger.yaml) file into the [swagger editor](https://editor.swagger.io/) or import the [Heroes-social-network.postman_collection.json](/docs/heroes-social-network.postman_collection.json) file into Postman .For more information about the project, visit: [Wiki](https://github.com/LeandroAlcantara-1997/heroes-social-network/wiki)


## Technologies

* Golang 1.20;
* PostgreSQL;
* Cache;
* Splunk;
* Swagger;


## Entity relationship diagram

![diagram](/docs/assets/heroes-social-network.jpg)