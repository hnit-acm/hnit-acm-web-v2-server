create database hnit_acm_system_service;

create table system_page_configure
(
    id     bigint unsigned auto_increment primary key auto_increment,
    name   varchar(256) not null default '' comment '页面名称',
    banner varchar(256) not null default '' comment '页面横幅图片'
)