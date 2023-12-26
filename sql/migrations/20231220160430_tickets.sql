-- +goose Up
create table if not exists tickets (
    id              int8 primary key generated by default as identity   ,
    user_id         int8            not null                            ,
    movie_id        int8            not null                            ,
    premieres_at    timestamptz     not null                            ,
    created_at      timestamptz     not null    default now()           ,

    foreign key (user_id)   references users(id)                        ,
    foreign key (movie_id)  references movies(id)
);

create table if not exists seats (
    id          int8 primary key generated by default as identity   ,
    ticket_id   int8        not null                                ,
    seat_number int2        not null                                ,

    foreign key (ticket_id) references tickets(id)
);

-- +goose Down
drop table if exists tickets;
drop table if exists seats;
