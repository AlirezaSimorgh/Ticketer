-- +goose Up
create table if not exists movies (
    id          int8 primary key generated by default as identity   ,
    name        text            not null    unique                  ,
    genres      []text          not null                            ,   
    show_time   int2            not null                            ,
    movie_time  int2            not null                            ,   
    actresses   []text          not null                            ,   
    description varchar(512)    not null                            ,   

    constraint movies_show_time_validation  check (0 <= show_time and show_time < 960),
    constraint movies_movie_time_validation  check (0 <= movie_time and movie_time < 960),
);

-- +goose Down
drop table if exists movies;