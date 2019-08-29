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

create table command (
    id int auto_increment,
    user_id varchar(80),
    name varchar(80),
    command varchar(255),
    use_ok int,
    foreign key (user_id) references user(id),
    primary key (id, user_id)
);