# Simple go backend server
for now just starting out

## Go progress report no:2
1) Project structure improved with much better scalability
2) First HTML template added, now site has a simple home page

##Go progress repost no:3
passed, sorry for this

##Go progress report no:4
Templating added, now html has less code and is more stuctured.

##Mideterm explanation

* First you need to install mysql or mariadb, every OS installation is different, so better google it.
* Then you need to go into mysql and create some tables and a database
```
sudo mysql
create database ecommerce
use ecommerce
create table products(id integer not null primary key auto_increment, name varchar(100) not null, description text not null, price integer not null, created datetime not null);
create table users(username varchar(64) not null, password varchar(64) not null);
```
* After you created these tables and populated them as you like, create a user named *web* with password *pass*, and git it an access to your database

```
create user 'web'@'localhost';
grant select, insert, update, delete on ecommerce.* to 'web'@'localhost';
alter user 'web'@'localhost' identified by 'pass';
```

* While inside a project folder, get the database driver
`go get github.com/go-sql-driver/mysql`

* Now run the server 
`go run ./cmd/web/`
