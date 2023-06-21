create database client;
create user client_infos;
alter user client_infos with encrypted password '1231';
grant all privileges on database client to client_infos;
\c client;
create table client_infos (id serial primary key, cpf text, private bool, incompleto bool, data_da_ultima_compra text, ticket_medio decimal, ticket_da_ultima_compra decimal, loja_mais_frequente text, loja_da_ultima_compra text);
grant all privileges on all tables in schema public to client_infos;
grant all privileges on all sequences in schema public to client_infos;