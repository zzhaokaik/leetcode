##  Mysql 知识点整理

###    外键

关系数据库通过外键可以实现一对多、多对多和一对一的关系。外键既可以通过数据库来约束，也可以不设置约束，仅依靠应用程序的逻辑来保证。


当我们用主键唯一标识记录时，我们就可以在students表中确定任意一个学生的记录：

``` 
id	name	other columns…
1	小明	…
2	小红	…
```
我们还可以在classes表中确定任意一个班级记录：
``` 
id	name	other columns…
1	一班	…
2	二班	…
```
但是我们如何确定students表的一条记录，例如，id=1的小明，属于哪个班级呢？

由于一个班级可以有多个学生，在关系模型中，这两个表的关系可以称为“一对多”，即一个classes的记录可以对应多个students表的记录。

为了表达这种一对多的关系，我们需要在students表中加入一列class_id，让它的值与classes表的某条记录相对应：

``` 
id	class_id	name	other columns…
1	1	小明	…
2	1	小红	…
5	2	小白	…
```
这样，我们就可以根据class_id这个列直接定位出一个students表的记录应该对应到classes的哪条记录。

例如：
``` 
小明的class_id是1，因此，对应的classes表的记录是id=1的一班；
小红的class_id是1，因此，对应的classes表的记录是id=1的一班；
小白的class_id是2，因此，对应的classes表的记录是id=2的二班。
```
在students表中，通过class_id的字段，可以把数据与另一张表关联起来，这种列称为外键。

外键并不是通过列名实现的，而是通过定义外键约束实现的：
``` 
ALTER TABLE studentsADD CONSTRAINT fk_class_idFOREIGN KEY (class_id)REFERENCES classes (id);
```
通过定义外键约束，关系数据库可以保证无法插入无效的数据。即如果classes表不存在id=99的记录，students表就无法插入class_id=99的记录。




