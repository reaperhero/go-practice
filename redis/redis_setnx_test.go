package redis

import "testing"

// SETNX 实现分布式锁
// 只有在 key 不存在时设置 key 的值。

// 方案1：SETNX + Delete
// 此实现方式的问题在于：一旦服务获取锁之后，因某种原因挂掉，则锁一直无法自动释放。从而导致死锁。
func Test_setNx_01(t *testing.T) {
	//setnx lock_a random_value
	//do sth
	//delete lock_a

}

// 方案2：SETNX + SETEX
// 按需设置超时时间。此方案解决了方案 1 死锁的问题，但同时引入了新的死锁问题：如果 SETNX 之后、SETEX 之前服务挂掉，会陷入死锁。根本原因为 SETNX/SETEX 分为了两个步骤，非原子操作
func Test_setNx_02(t *testing.T) {
	//setnx lock_a random_value
	//setex lock_a 10 random_value // 10s超时
	//do sth
	//delete lock_a
}

// 方案3：SET NX PX
// 此方案通过 SET 的 NX/PX 选项，将加锁、设置超时两个步骤合并为一个原子操作，从而解决方案 1、2 的问题。( PX 与 EX 选项的语义相同，差异仅在单位。)
// 如果锁被错误的释放（如超时），或被错误的抢占，或因redis问题等导致锁丢失，无法很快的感知到。
func Test_setNx_03(t *testing.T) {
	//SET lock_a random_value NX PX 10000 // 10s超时
	//do sth
	//delete lock_a
}

// 方案4：SET Key RandomValue NX PX
// 方案 4 在 3 的基础上，增加对 value 的检查，只解除自己加的锁。类似于 CAS，不过是 compare-and-delete。此方案 Redis 原生命令不支持，为保证原子性，需要通过 Lua 脚本实现。
// 此方案更严谨：即使因为某些异常导致锁被错误的抢占，也能部分保证锁的正确释放。并且在释放锁时能检测到锁是否被错误抢占、错误释放，从而进行特殊处理
func Test_setNx_04(t *testing.T) {
	//SET lock_a random_value NX PX 10000
	//do sth
	//eval "if redis.call('get',KEYS[1]) == ARGV[1] then return redis.call('del',KEYS[1]) else return 0 end" 1 lock_a random_value
}
