create database system_service;
create table system_page_configure
(
    id     bigint unsigned auto_increment primary key auto_increment,
    name   varchar(256) not null default '',
    banner varchar(256) not null default ''
)