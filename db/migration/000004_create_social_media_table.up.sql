begin;
create table if not exists social_media(
    Id SERIAL primary key,
    name text ,
    social_media_url text,
    user_id int not null references public.users(id)
    
);

commit;