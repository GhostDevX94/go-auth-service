create table users
(
    id                bigserial
        primary key,
    name              varchar(255) not null,
    surname           varchar(255) not null,
    phone             varchar(255),
    email             varchar(255) not null
        constraint users_email_unique
            unique,
    password          varchar(255) not null
);
