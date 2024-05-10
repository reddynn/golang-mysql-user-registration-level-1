## repos

* [level-1](https://github.com/reddynn/golang-mysql-user-registration-level-1)
* [level-2](https://github.com/reddynn/golang-mysql-user-registration-level-2)
* [level-3](https://github.com/reddynn/golang-mysql-user-registration-level-3)
* [level-4](https://github.com/reddynn/golang-mysql-user-registration-level-4)
* [level-5](https://github.com/reddynn/golang-mysql-user-registration-level-5)
......................... level-N
## Welcome Note

```
Everyone is welcome for code contributions or code enhancemnets.
please give a star, fork the repo and create PR for your code contributions
```
## db commands

```
create database userdb;
use userdb;
create table users(username varchar(255), password varchar(255));
```
## run the application with go run

```
go run main.go
```
## `Why I started this project`
```
as a beginner if we look into any projects we may feel that its too complex to understand, it happens regularly with me, so that's where this project came to my mind, why can't i build something from basics and increase the standards and it should be documented so that if i go back and look into that i should feel that how the coding levels and standards are increased from very basics to advanced.
```
## `level-1:`

* basic user authentication functionalities like signup and login
* we are using mysql database to store the password.
* we are taking user details like username and password from query parameters

## `level-2:`

* we are taking username and password from request body
* used structs and decoding methods
* added air live reloading.

## `level-3:`

* converted entire application logic into golang folder structure

## `level-4:`

* added environment specific configuration using viper library
*  added validation library to validate the structs.
*  added better decoding methods in golang.

## `level-5:`

* added sign out functionality using golang sessions and cookies.

* ......
* .....
* ...
* ..
* .

## `upcoming levels:`
* storing tokens in databases
* jwt authentication mechanism will be added
* coding standards will be increased.

## `level-N:`
* application should be production ready
