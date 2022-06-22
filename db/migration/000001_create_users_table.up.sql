begin;

create table if not exists users(
id SERIAL primary key,
username text unique not null,
email text unique not null,
password text not null,
age int not null,
created_date date,
updated_date date
);

commit;