package sqlx

// 涉及锁的维度看条件语句的扫描范围，不带索引的那么锁的就是全表
// sqlx两个事务，涉及到查询同一条数据，当两条都使用for update指令，会阻塞宁外一条
// sqlx只开启一个事务，当事务使用for update指令，会阻塞宁外一条不带事务的for update
// sqlx只开启一个事务，当事务使用for update指令，不会阻塞宁外一条（不带事务、不带for update）的语句
