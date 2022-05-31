drop table if exists users cascade;

create table users (id serial primary key, email varchar, password varchar);

drop table if exists sites cascade;

create table sites (id serial primary key, title varchar, description varchar);

drop table if exists posts cascade;

create table posts (
    id serial primary key,
    body text,
    site_id integer references sites(id),
    created_at timestamp default now(),
    updated_at timestamp default now()
);

drop table if exists pages cascade;

create table pages (
    id serial primary key,
    title varchar,
    body text,
    slug varchar,
    site_id integer references sites(id),
    created_at timestamp default now(),
    updated_at timestamp default now()
);

drop table if exists events cascade;

create table events (
    id serial primary key,
    title varchar,
    body text,
    event_date timestamp,
    site_id integer references sites(id),
    created_at timestamp default now(),
    updated_at timestamp default now()
);
