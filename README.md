
## Table of contents

* [General info](#general-info)
* [Setup](#setup)
* [Routes](#routes)
  

## General info

This project is an example of simple service that utilizes few mechanics.
* user account management, router for user endpoints ->[user_router.go](https://github.com/UnoHouse/api/blob/master/http/user_routes.go)
* some miscellaneous endpoints that works on different schema and service
[other_routes.go](https://github.com/UnoHouse/api/blob/master/http/other_routes.go)

  

## Setup

To run this project, You will need to 

 - download dependencies for project
 - create database with two schemas (DDL is located in [build/db.sql](https://github.com/UnoHouse/api/blob/master/build/db.sql))
	 - unohouse
	 - users
 - prepare .env file in project root directory keys structure

```
LISTEN=0.0.0.0:8001
APP_ENV=develop
MYSQL_USER=
MYSQL_PASSWORD=
MYSQL_HOSTS=
MYSQL_PORT=
MYSQL_MAIN_DATABASE=
MYSQL_USERS_DATABASE=
```
- build or run app via main.go file
  
  Let's hope everything will work as expected :)

## TODO

* add goroutines for

	* notifications sendout
	* faker generators
		* more example users
		* more devices

* add tests
	* at least for core endpoints

* add swagger

* add travis

* add circle-ci

* improve this documentation
