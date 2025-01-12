create table api_data
(
    id         varchar(255) not null
        primary key,
    content    text         null,
    popularity int          null,
    token      varchar(255) null,
    ip_address varchar(255) null,
    cache_at   datetime     null
);

create table favorites
(
    user_id    varchar(255)                        not null,
    poem_id    varchar(255)                        not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    primary key (user_id, poem_id)
);

create table match_tags
(
    id      int auto_increment
        primary key,
    data_id varchar(255) null,
    tag     varchar(255) null,
    constraint match_tags_ibfk_1
        foreign key (data_id) references api_data (id)
            on delete cascade
);

create index data_id
    on match_tags (data_id);

create table notes
(
    id           int auto_increment
        primary key,
    title        varchar(255)                        not null,
    content      text                                not null,
    image_base64 mediumtext                          null,
    username     varchar(255)                        not null,
    created_at   timestamp default CURRENT_TIMESTAMP null
);

create table origin
(
    id      int auto_increment
        primary key,
    data_id varchar(255) null,
    title   varchar(255) null,
    dynasty varchar(255) null,
    author  varchar(255) null,
    constraint origin_ibfk_1
        foreign key (data_id) references api_data (id)
            on delete cascade
);

create index data_id
    on origin (data_id);

create table origin_content
(
    id      int auto_increment
        primary key,
    data_id varchar(255) null,
    content text         null,
    constraint origin_content_ibfk_1
        foreign key (data_id) references api_data (id)
            on delete cascade
);

create index data_id
    on origin_content (data_id);

create table origin_translate
(
    id        int auto_increment
        primary key,
    data_id   varchar(255) null,
    translate text         null,
    constraint origin_translate_ibfk_1
        foreign key (data_id) references api_data (id)
            on delete cascade
);

create index data_id
    on origin_translate (data_id);

create table users
(
    id         int auto_increment
        primary key,
    username   varchar(50)                         not null,
    password   varchar(255)                        not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    constraint username
        unique (username)
);


