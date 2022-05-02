package _defer

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
	"testing"
	"time"
)

// 在使用数据库事务时，我们可以使用上面的代码在创建事务后就立刻调用 Rollback 保证事务一定会回滚。
// 哪怕事务真的执行成功了，那么调用 tx.Commit() 之后再执行 tx.Rollback() 也不会影响已经提交的事务。
func createPost(db *gorm.DB) error {
	tx := db.Begin()
	type User struct {
	}
	defer tx.Rollback()
	if err := tx.Create(&User{}).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

func printSinceTime() {
	startedAt := time.Now()
	defer func() {
		fmt.Println(time.Since(startedAt))
	}()
	time.Sleep(time.Second)
}

type D struct {
	exist   *sync.Cond
	success bool
}

var chanA = make(chan int)

func (d *D) printA() {
	go func() {
		d.exist.L.Lock()
		d.exist.Wait()
		<-chanA
		if d.success {
			fmt.Println("A")
		}
		d.exist.L.Unlock()
	}()
	d.printB(chanA)
	return
}
func (d *D) printB(c chan int) {
	chanB := make(chan int)
	go func() {
		d.exist.L.Lock()
		d.exist.Wait()
		<-chanB
		if d.success {
			fmt.Println("B")
			c <- 1
		}
		d.exist.L.Unlock()
	}()
	d.printC(chanB)
	return
}
func (d *D) printC(c chan int) {

	go func() {
		d.exist.L.Lock()
		d.exist.Wait()
		if d.success {
			fmt.Println("C")
			c <- 1
		}
		d.exist.L.Unlock()
	}()
	return
}

func TestName(t *testing.T) {
	d := D{
		exist:   sync.NewCond(&sync.Mutex{}),
		success: true,
	}
	d.printA()
	time.Sleep(time.Second)
	d.exist.Broadcast()
	time.Sleep(10 * time.Second)
}

