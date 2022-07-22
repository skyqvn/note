create database mydatabase;

drop database mydatabase;

use mydatabase;

show tables;

show databases;

show full processlist;

set persist max_connections = 500;

create table if not exists `mytable`
(
    `id`              int unsigned auto_increment,
    `title`           varchar(100) not null default '',
    `author`          varchar(40)  not null default '',
    `submission_date` date                  default current_timestamp,
    primary key (`id`),
    unique key (`title`)
) engine = innodb
  default charset = utf8;

drop table mytable;

desc mytable;

show columns from mytable;

insert into mytable (title, author)
values ('golang', 'google');

insert into mytable(title)
values ('java'),
       ('mysql'),
       ('python');

select id, title
from mytable
where id = 1;

update mytable
set title='new book'
where id = 1;

select *
from mytable
where author like '%com';

select *
from mytable
where author like '%com'
limit 1;


delete
from mytable
where id = 1;

alter table mytable
    drop member_length;

alter table mytable
    add member_length int;

alter table mytable
    modify title char(10);

alter table mytable
    change member_length j bigint;

alter table mytable
    alter member_length set default 1000;

alter table mytable rename to mytable2;

select *
from mytable
order by submission_date asc;

select *
from mytable
order by submission_date desc;

CREATE INDEX myindex
    ON mytable (title);

CREATE INDEX myindex
    ON mytable (title desc);

CREATE INDEX myindex
    ON mytable (title, submission_date);

CREATE UNIQUE INDEX myindex
    ON mytable (title);

ALTER TABLE mytable
    ALTER INDEX myindex INVISIBLE;

ALTER TABLE mytable
    ALTER INDEX myindex VISIBLE;

create table if not exists `mytable2`
(
    `id`              int unsigned auto_increment primary key,
    `submission_date` date default current_timestamp
        on update current_timestamp,
    `mytable_id`      int unsigned,
    constraint mytable2_mi foreign key (`mytable_id`) references mytable (id)
) engine = innodb
  default charset = utf8;


/*备份数据
mysqldump -uroot -p123456 mydatabase > /database/mydatabase.sql

mysqldump -uroot -p123456 -d mydatabase > /database/mydatabase.sql

*/
