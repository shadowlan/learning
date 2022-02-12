# awk基础语法与示例

首先建立一个示例文件employee.txt，内容如下
$cat >employee.txt<<EOF
101,John Doe,CEO
102,Jason Smith,IT Manager
103,Raj Reddy,Sysadmin
104,Anand Ram,Developer
105,Jane Miller,Sales Manager
EOF

awk语法
Awk –Fs 'BEGIN { awk-commands } /pattern/ {action} END { awk-commands }' input-file
BEGIN 区域的命令只最开始、在 awk 执行 body 区域命令之前执行一次。
END 区域在 awk 执行完所有操作后执行，并且只执行一次。
BEGIN和END区域都是可选的
如果命令很长，即可以放到单行执行，也可以用\折成多行执行,例如：

```bash
awk 'BEGIN{FS=",";print "---header---"} \
/Thomas/ {print $1} \
END{print "---footer---"}' employee.txt
```


## awk内置变量

* FS - 输入字段分隔符
```bash
$cat > employee-multiple-fs.txt <<EOF
101,John Doe:CEO%10000
102,Jason Smith:IT Manager%5000
103,Raj Reddy:Sysadmin%4500
104,Anand Ram:Developer%4500
105,Jane Miller:Sales Manager%3000
EOF
```

如employee-multiple-fs.txt所示有多个不同的分割符时，可以使用正则表达式来指定多个分隔符：
`awk 'BEGIN{FS="[,:%]"}{print $2,$3}' employee-multiple-fs.txt`


* OFS - 输出字段分隔符

awk默认OFS为空格符，如果修改输出字段之间的分隔符，可以修改OFS. 

请注意在 print 语句中使用和不使用逗号的细微差别(打印多个变量时).当在print语句中指定了逗号，awk会使用OFS。如下面的例子所示，指定的OFS会被使用，所以你会看到输出值之间的":"。不使用逗号是，awk将不会使用OFS，其输出变量之间没有任何空格或者指定的OFS。
`awk 'BEGIN{FS="[,:%]";OFS=":"}{print $2,$3}' employee-multiple-fs.txt`

* RS – 记录分隔符

```bash
cat > employee-one-line.txt <<EOF
101,John Doe:102,Jason Smith:103,Raj Reddy:104,Anand Ram:105,Jane, Miller
EOF
```

默认awk会将文件中的内容作为单独一行来处理，如果要把文件内容作为5行记录来处理(而不是单独的一行)，并且打印每条记录中雇员的姓名，就必须把记录分隔符指定为分号，如下所示：

$awk -F, 'BEGIN { RS=":" } {print $2}' employee-one-line.txt
John Doe
Jason Smith
Raj Reddy
Anand Ram
Jane Miller


* ORS – 输出记录分隔符

awk默认使用换行符"\n"作为ORS,如果想修改打印输出结果的行分隔符，可以修改ORS: 
为每个输出行后面追加"---": `awk 'BEGIN{ORS="\n---\n"}{print $1}' employee.txt`

* NR – 记录序号

在循环内部标识记录序号。用于 END 区域时，代表输入文件的总记录数. `awk 'BEGIN{FS=","}{print "id of line" ,NR, "is",$1}' employee.txt`

* FILENAME – 当前处理的文件名

当时用awk批处理多个文件时，FILENAME能够显示当前处理文件的名称： `awk '{print FILENAME}' employee.txt employee-multiple-fs.txt`
如果awk从标准输入获取内容或者通过echo,FILENAME名显示`-`(mac下为空白): `echo 'john doe' |awk '{print $2;print FILENAME}'`

* FNR – 文件中的 NR

当awk处理多个文件时，NR会在多个文件中持续增加，此时NR并不是当前处理行在文件中的实际行数，如果要显示实际行数，就需要用到FNR。

* NF - 字段数量

当条记录中的字段个数： `awk -F, '{print NF}' employee.txt`

## awk变量操作符

```bash
$cat >employee-sal.txt<<EOF
101,John Doe,CEO,10000
102,Jason Smith,IT Manager,5000
103,Raj Reddy,Sysadmin,4500
104,Anand Ram,Developer,4500
105,Jane Miller,Sales Manager,3000
EOF
```

`awk 'BEGIN{FS=",";total=0;}{print $2 "'\''s salary is:" $4;total=total+$4}END{print "total salary is",total}'  employee-sal.txt`
在awk中打印单引号需要比较麻烦： `'\''`,可以将其定义为变量后在awk中使用： `awk -v sq=\' 'BEGIN{FS=",";total=0;}{print $2 sq"s salary is:" $4;total=total+$4}END{print "total salary is",total}'  employee-sal.txt`,定义变量的方式在需要多次使用单引号时很有帮助。

## awk操作符

* 一元操作符

| 擦作符 |      描述        |
|-------|-----------------|
|  +    | 取正（返回数字本身) |
|  -    |  取反    |
| ++    | 自增     |
| --    | 自减     |

自增/自减操作符可在变量前或者变量后。使用改变后的变量值(pre)即是在变量前面加上++(或--)，使用改变前的变量值(post)即是在变量后面加上++(或--)，

* awk算数操作符: [+-*/%](加减乘除取模)

* 比较操作符: ｜ > ｜ >= ｜ < ｜ <= ｜ == ｜ != ｜ && ｜ || ｜
打印薪水少于4500的职工信息: `awk -F, '$4<4500' employee-sal.txt`
打印非CEO职工信息: `awk -F, '$3!="CEO"' employee-sal.txt`

* 正则表达式操作符: ~:匹配, !~:不匹配

## awk分支和循环
...

