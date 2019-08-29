# user

## user

- id : varchar(80), primary key
- password_hash : varchar(255)

## session

- id : varchar(80), primary key, foreign key references column user(id)
- passphrase : varchar(30)
- expiration_unix_time : int

## command

- id : int, primary key, auto_increment
- user_id : varchar(80), primary key, foreign key references column user(id)
- name : varchar(80)
- command : varchar(255)
- use_ok : int
