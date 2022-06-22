begin;
create table if not exists photos(
    Id SERIAL primary key,
    Title text not null,
    Caption text not null,
    Url text not null,
    user_id int not null references public.users(id),
    Created_date date,
    Updated_date date
);

commit;