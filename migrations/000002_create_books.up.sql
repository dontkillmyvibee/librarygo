begin;

create table library.books
(
    id           uuid primary key      default gen_random_uuid(),
    title        varchar(100) not null,
    description  varchar(1000),
    is_available bool         not null default true,
    created_at   timestamptz  not null default now(),
    updated_at   timestamptz  not null default now()

        constraint books_updated_at_check
            check (updated_at >= created_at)
);

create table library.book_authors
(
    book_id   uuid not null,
    author_id uuid not null,

    primary key (book_id, author_id),

    foreign key (book_id)
        references library.books (id)
        on delete cascade,

    foreign key (author_id)
        references library.authors (id)
        on delete cascade
);

create index book_authors_author_id_idx
    on library.book_authors (author_id);

commit;