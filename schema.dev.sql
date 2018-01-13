create table users (id integer not null primary key, username text UNIQUE, password text);
create table groups (id integer not null primary key, name text);
create table membership (userid integer not null, groupid integer not null);
create table receipts (id integer not null primary key, groupid integer not null);
create table items (id integer not null primary key, receiptid integer not null, name text);
create table rules (id integer not null primary key, groupid integer not null, name text);
delete from users;
delete from groups;
delete from membership;
delete from receipts;
delete from items;
delete from rules;