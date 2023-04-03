#mysql 杂
##  左右链接 join
查询分析器中执行：
```
--建表table1,table2：
create table table1(id int,name varchar(10))
create table table2(id int,score int)
insert into table1 select 1,'lee'
insert into table1 select 2,'zhang'
insert into table1 select 4,'wang'
insert into table2 select 1,90
insert into table2 select 2,100
insert into table2 select 3,70

```
如表:
```
-------------------------------------------------
table1 | table2 |
-------------------------------------------------
id name |id score |
1 lee |1 90|
2 zhang| 2 100|
4 wang| 3 70|
-------------------------------------------------
复制代码
以下均在查询分析器中执行 一、外连接 1.概念：包括左向外联接、右向外联接或完整外部联接

2.左连接：left join 或 left outer join (1)左向外联接的结果集包括 LEFT OUTER 子句中指定的左表的所有行，而不仅仅是联接列所匹配的行。如果左表的某行在右表中没有匹配行，则在相关联的结果集行中右表的所有选择列表列均为空值(null)。 (2)sql 语句

select * from table1 left join table2 on table1.id=table2.id
-------------结果-------------
idnameidscore
------------------------------
1lee190
2zhang2100
4wangNULLNULL
------------------------------
```
注释：包含table1的所有子句，根据指定条件返回table2相应的字段，不符合的以null显示

3.右连接：right join 或 right outer join (1)右向外联接是左向外联接的反向联接。将返回右表的所有行。如果右表的某行在左表中没有匹配行，则将为左表返回空值。 (2)sql 语句
```
select * from table1 right join table2 on table1.id=table2.id
-------------结果-------------
idnameidscore
------------------------------
1lee190
2zhang2100
NULLNULL370
------------------------------
```
注释：包含table2的所有子句，根据指定条件返回table1相应的字段，不符合的以null显示

4.完整外部联接:full join 或 full outer join (1)完整外部联接返回左表和右表中的所有行。当某行在另一个表中没有匹配行时，则另一个表的选择列表列包含空值。如果表之间有匹配行，则整个结果集行包含基表的数据值。 (2)sql 语句
```
select * from table1 full join table2 on table1.id=table2.id
-------------结果-------------
idnameidscore
------------------------------
1lee190
2zhang2100
4wangNULLNULL
NULLNULL370
------------------------------
```
注释：返回左右连接的和（见上左、右连接）

二、内连接 1.概念：内联接是用比较运算符比较要联接列的值的联接

2.内连接：join 或 inner join

3.sql 语句
```
select * from table1 join table2 on table1.id=table2.id
-------------结果-------------
idnameidscore
------------------------------
1lee190
2zhang2100
------------------------------
```
注释：只返回符合条件的table1和table2的列

4.等价（与下列执行效果相同）
```
A:select a.*,b.* from table1 a,table2 b where a.id=b.id
B:select * from table1 cross join table2 where table1.id=table2.id (注：cross join后加条件只能用where,不能用on)
```
三、交叉连接(完全)

1.概念：没有 WHERE 子句的交叉联接将产生联接所涉及的表的笛卡尔积。第一个表的行数乘以第二个表的行数等于笛卡尔积结果集的大小。（table1和table2交叉连接产生3*3=9条记录）

2.交叉连接：cross join (不带条件where...)

3.sql语句
```
select * from table1 cross join table2
-------------结果-------------
idnameidscore
------------------------------
1lee190
2zhang190
4wang190
1lee2100
2zhang2100
4wang2100
1lee370
2zhang370
4wang370
------------------------------
```
注释：返回3*3=9条记录，即笛卡尔积

4.等价（与下列执行效果相同）
```
A:select * from table1,table2
```



###    解释 SQL 的 left join 和 right join
```
left join 和 right join 都是两个表进行 merge 的操作，left join 是将右边的表 merge 到左边，right join 是将左边的表 merge 到右边，通常我们会指定按照哪几列进行 merge

举个例子：

left table

姓名	学号
小红	SZ1716029
小明	SZ1716030
小王	SZ1716031
right table

学号	排名
SZ1716029	1
SZ1716030	2
left table left join right table on 学号

学号	姓名	排名
SZ1716029	小红	1
SZ1716030	小明	2
SZ1716031	小王	NULL
left table right join right table on 学号

学号	姓名	排名
SZ1716029	小红	1
SZ1716030	小明	2

```


###    事务的隔离级别
``` 
Mysql有四种事务隔离级别,默认的是可重复读.

事务隔离级别	脏读	不可重复读	幻读
读未提交	     是	是	        是
读已提交	     否	是	        是
可重复读	     否	否	        是
串行      	 否	否	        否
读未提交(Read uncommitted)
一个事务可以读取另一个未提交事务的数据，最低级别，任何情况都无法保证。

(1)所有事务都可以看到其他未提交事务的执行结果 
(2)本隔离级别很少用于实际应用，因为它的性能也不比其他级别好多少 (3)该级别引发的问题是——脏读(Dirty Read)：读取到了未提交的数据

读已提交(Read committed)
一个事务要等另一个事务提交后才能读取数据，可避免脏读的发生。

(1)这是大多数数据库系统的默认隔离级别（但不是MySQL默认的） 
(2)它满足了隔离的简单定义：一个事务只能看见已经提交事务所做的改变 (3)这种隔离级别出现的问题是——不可重复读(Nonrepeatable Read),不可重复读意味着我们在同一个事务中执行完全相同的select语句时可能看到不一样的结果。

导致这种情况的原因可能有：

(1)有一个交叉的事务有新的commit，导致了数据的改变; 
(2)一个数据库被多个实例操作时,同一事务的其他实例在该实例处理其间可能会有新的commit.

可重复读(Repeatable read)
就是在开始读取数据（事务开启）时，不再允许修改操作，可避免脏读、不可重复读的发生。

(1)这是MySQL的默认事务隔离级别 
(2)它确保同一事务的多个实例在并发读取数据时，会看到同样的数据行 
(3)此级别可能出现的问题——幻读(Phantom Read)：当用户读取某一范围的数据行时，另一个事务又在该范围内插入了新行，当用户再读取该范围的数据行时，会发现有新的“幻影” 行 
(4)InnoDB和Falcon存储引擎通过多版本并发控制(MVCC，Multiversion Concurrency Control)机制解决了该问题.InnoDB采用MVCC来支持高并发，实现了四个标准隔离级别。默认基本是可重复读，并且提供间隙锁（next-key locks）策略防止幻读出现。

串行(Serializable)
串行(Serializable)，是最高的事务隔离级别，在该级别下，事务串行化顺序执行，可以避免脏读、不可重复读与幻读。 但是这种事务隔离级别效率低下，比较耗数据库性能，一般不使用。Mysql的默认隔离级别是Repeatable read。

(1)这是最高的隔离级别. 
(2)它通过强制事务排序，使之不可能相互冲突，从而解决幻读问题。简言之,它是在每个读的数据行上加上共享锁。 
(3)在这个级别，可能导致大量的超时现象和锁竞争.
```