create table text_record (
    id varchar(255) primary key,
    last_update_date timestamp not null,
    user_login int,

    text text,

    constraint fk_text_record_user
        foreign key(user_login)
            references users(login)
);

create table bank_card_record (
    id varchar(255) primary key,
    last_update_date timestamp not null,
    user_login int,

    card_number varchar(16) not null,
    month int not null,
    day int not null,
    code int not null,

    constraint fk_bank_record_user
        foreign key(user_login)
            references users(login)
);

create table binary_record (
    id varchar(255) primary key,
    last_update_date timestamp not null,
    user_login int,

    "binary" bytea,

    constraint fk_binary_record_user
        foreign key(user_login)
            references users(login)
);

create table login_password_record (
    id varchar(255) primary key,
    last_update_date timestamp not null,
    user_login int,

    login varchar(255) not null,
    password varchar(255) not null,

    constraint fk_login_password_record_user
        foreign key(user_login)
            references users(login)
);
