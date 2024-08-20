create database todoListApp
go

create table USUARIOS(
	id  int identity(1,1) primary key,
	username varchar(30),
	email varchar(30),
	password varchar(90)
)
GO

create table TASKS(
	id  int identity(1,1) primary key,
	title  varchar(30),
	idusuario int not null,
	status  char(1) default 'P',
	create_at datetime default GETDATE(),

	constraint fk_idusuario_tasks foreign key (idusuario) references USUARIOS(id) 
)
GO

select * from usuarios