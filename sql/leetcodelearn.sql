# CONCAT 用来拼接字符串 ● LEFT 从左边截取字符 ● RIGHT 从右边截取字符 ● UPPER 变为大写 ● LOWER 变为小写 ● LENGTH 获取字符串长度

select user_id, CONCAT(UPPER(left(name, 1)), LOWER(SUBSTRING(name, 2))) as name
select user_id, CONCAT(UPPER(left(name, 1)), LOWER(RIGHT(name, length(name) - 1))) as name
from Users
order by user_id


# 2.1 CONCAT() 函数
# CONCAT 可以将多个字符串拼接在一起。

# 2.2 LEFT(str, length) 函数
# 从左开始截取字符串，length 是截取的长度。

# 2.3 UPPER(str) 与 LOWER(str)
# UPPER(str) 将字符串中所有字符转为大写
# LOWER(str) 将字符串中所有字符转为小写

# 2.4 SUBSTRING(str, begin, end)
# 截取字符串，end 不写默认为空。
# SUBSTRING(name, 2) 从第二个截取到末尾，注意并不是下标，就是第二个。


###########################
# 按日期分组销售产品


select sell_date,count(distinct product) as num_sold ,group_concat(distinct product) as products
from Activities
group by sell_date
order by sell_date asc