-- 服务启动
-- pg_ctl start -D /usr/local/var/postgres

-- 暂停
-- pg_ctl stop -D /usr/local/var/postgres

-- 确认服务状态
-- ps aux | grep postgres

-- 查看版本
-- postgres --version


create database mydb;

-- 删除数据库
drop database mydb;

-- 显示数据库列表
\l

-- 选择数据库

\c mydb

-- 用户

-- 创建用户
create user myuser;

-- 给用户赋予权限：示例赋予最大权限
grant all privileges on database mydb to myuser;

-- 用户列表
\du

-- 赋予指定权限
grant select, insert, update, delete on mytable to myuser;

-- 删除指定权限
revoke select, insert, update, delete on mytable from myuser;


-- Schema

-- 创建Schema
create schema myschema;

-- 确认当前的Schema
select current_schema;

-- Schema列表
\dn


-- 表

-- 创建表
create table mytable
(
    id      int primary key not null,
    name    text            not null,
    age     int             not null,
    address char(50),
    salary  real
);

-- 显示表的列表
\dt

-- 显示指定表
\d mytable

-- 显示指定表的数据
select *
from mytable;

select *
from mytable
order by id desc;

-- 修改表的owner
alter table mytable
    owner to owner名;


-- 表结构修改

-- 添加列
alter table mytable
    add sex char(6);

-- 删除列
alter table mytable
    drop sex;

-- 列名变更
alter table mytable
    rename address to addr;

-- 改变列数据类型
alter table mytable
    alter salary type int;


-- 索引

-- 创建索引
create index myindex on mytable (name);

-- 删除索引
drop index myindex;


-- 视图

-- 创建视图
create view myview as
(
select *
from mytable);

-- 显示视图列表
\dv

-- 使用视图
select *
from myview;

-- 删除视图
drop view myview;

-- 读入外部SQL文件
\i 文件名
