$ docker exec -it {container_id} bash

$ mysql -uroot -p goexpert

mysql> create table products(id varchar(255), name varchar(80), price decimal(10,2), primary key(id));

mysql> show tables;

mysql> select * from products;