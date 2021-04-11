drop database if exists cfsservice;

create database if not exists cfsservice CHARACTER SET utf8 COLLATE utf8_general_ci;
create user if not exists cfsdev@'localhost' identified with mysql_native_password by 'cfsdev';
grant all on cfsservice.* to cfsdev@'localhost';