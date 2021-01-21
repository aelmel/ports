# ports
Copy the repository
```
$ git clone https://github.com/aelmel/ports.git
```

Build services with docker-compose
```
$ docker-compose up -d
```
Call Get port endpoint 
```
$ curl --location --request GET 'localhost:8008/port/ZAJNB'
```
Response **Not found** 
Copy your json file to /var/lib/docker/volumes/ports-infra_client_api/_data client_api will monitor each minute this folder for json files
```
$ cp ports.json /var/lib/docker/volumes/ports-infra_client_api/_data
```
This will populate MongoDb database
Call Get port endpoint 
```
$ curl --location --request GET 'localhost:8008/port/ZAJNB'
```
Response should contain port details
