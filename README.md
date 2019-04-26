 # Token Service
 
 [![Build Status](https://travis-ci.com/p0sixEDfalls/Token-Service.svg?branch=master)](https://travis-ci.com/p0sixEDfalls/Token-Service)

# What is the project?
This is simple implementation of service, that allow to buy ERC20 tokens using Ethereum cryptocurrency blockchain.

 # API Documentation
Visit to [SwaggerHub](https://app.swaggerhub.com/apis/p0sixEDfalls/Token-Service/1.0.0-oas3) to see more information about service API.

 # How to build
 If you are on Linux open the terminal (also these commands will succefully work on WindowsNT and MacOS) and type:
 
 Install *gorilla/mux* package:
 
 	$ go get -u github.com/gorilla/mux
  
 Install *ethereum/go-ethereum* package:
 
 	$ go get -u github.com/ethereum/go-ethereum
  
  Install *jinzhu/gorm* package:
 
 	$ go get -u github.com/jinzhu/gorm
  
  Install *dgrijalva/jwt-go* package:
 
 	$ go get -u github.com/dgrijalva/jwt-go
  
  Install *lib/pq* package:
 
 	$ go get -u github.com/lib/pq
  
  Clone the repository:

	$ git clone https://github.com/p0sixEDfalls/Token-Service.git
  
  Open the folder where cloned repository is located and build:
  
    $ go build ./src/main/main.go
    
   # How to build docker images
   Clone the repository:

	$ git clone https://github.com/p0sixEDfalls/Token-Service.git
  
  Open the folder where cloned repository is located and use docker-compose:
  
    $ docker-compose up .
    
   # Any questions?
   Contact me:

	xdb7493c6e70@yandex.ru
   
   
   
  