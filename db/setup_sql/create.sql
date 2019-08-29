create table user (
    id varchar(80) primary key,
    password_hash varchar(255)
);

create table session (
    id varchar(80) primary key,
    passphrase varchar(30),
    expiration_unix_time int,
    foreign key (id) references user(id)
);
