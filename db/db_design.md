# user

## user

- id : varchar(80), primary key
- password_hash : varchar(255)

## session

- id : varchar(80), primary key, foreign key references column user(id)
- passphrase : varchar(30)
- expiration_unix_time : int
