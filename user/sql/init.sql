create database user_service;
# 用户信息表
create table user_info
(
    id          bigint unsigned auto_increment primary key,
    code        varchar(32) default ''                not null comment '身份识别码',
    nickname    varchar(32) default ''                not null comment '昵称',
    password    varchar(32) default ''                not null comment '密码',
    salt        varchar(4)  default ''                not null comment '盐',
    group_id    bigint      default ''                not null comment '群组id',
    create_time timestamp   default current_timestamp not null,
    update_time timestamp   default current_timestamp not null,
    delete_time timestamp   default null
);
# 用户群组表
create table user_group
(
    id             bigint unsigned auto_increment primary key,
    name           varchar(32)     default ''                not null,
    create_user_id bigint unsigned default 0                 not null,
    create_time    timestamp       default current_timestamp not null,
    update_time    timestamp       default current_timestamp not null,
    delete_time    timestamp       default null
);
# 用户组权限表
# 用户组的创建者默认拥有增删权限
create table user_group_auth
(
    id          bigint unsigned auto_increment primary key,
    group_id    bigint    default ''                not null comment '群组id',
    is_delete   tinyint   default 0                 not null comment '',
    is_add      tinyint   default 0                 not null comment '',
    create_time timestamp default current_timestamp not null comment '',
    update_time timestamp default current_timestamp not null comment '',
    delete_time timestamp default null
);

# 用户组