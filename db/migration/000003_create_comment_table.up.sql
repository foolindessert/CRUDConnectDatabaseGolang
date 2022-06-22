begin;
create table if not exists comment(
    Id SERIAL primary key,
    user_id int not null references public.users(id),
    photo_id int not null references public.photos(id),
    message text,
    Created_date date,
    Updated_date date
);

commit;