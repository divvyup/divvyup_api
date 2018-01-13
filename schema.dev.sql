create table users (id integer not null primary key, username text UNIQUE, password text);
delete from users;