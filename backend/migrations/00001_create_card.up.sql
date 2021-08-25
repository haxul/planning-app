create table cards
(
    id          UUID         not null primary key,
    state       varchar(50)  not null,
    title       varchar      not null,
    description varchar(500) not null,
    tag         varchar(20)  not null,
    update_on   timestamp    not null
)