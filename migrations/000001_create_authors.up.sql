begin;

create schema library;

create table library.authors
(
    id          uuid primary key      default pg_catalog.gen_random_uuid(),
    first_name  varchar(100) not null,
    last_name   varchar(100) not null,
    middle_name varchar(100),
    created_at  timestamptz  not null default now(),
    updated_at  timestamptz  not null default now(),

    constraint authors_update_at_check
        check (updated_at >= created_at)
);

commit;