CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE IF NOT EXISTS users.users(
    id serial not null,
    tg_id int not null,
    unique(tg_id),
    primary key(id)
);

CREATE SCHEMA IF NOT EXISTS drugs;

CREATE TABLE IF NOT EXISTS drugs.drugs(
    id serial not null,
    user_id int not null,
    name text not null,
    desription text,
    created timestamp not null,
    take_today bool not null default false,
    today_count int not null default 0,
    primary key(id),
    foreign key (user_id) references users.users(id) on delete cascade
)