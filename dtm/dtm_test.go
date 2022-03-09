package dtm

import (
	"fmt"
	"github.com/dtm-labs/dtmcli"
	"github.com/gin-gonic/gin"
	"log"
	"testing"
	"time"
)

// 解决跨数据库、跨服务、跨语言栈更新数据的一致性问题。
// 通俗一点说，DTM提供跨服务事务能力，一组服务要么全部成功，要么全部回滚，避免只更新了一部分数据产生的一致性问题。

const qsBusiAPI = "/api/busi_start"
const qsBusiPort = 8082
const dtmServer = "http://localhost:36789/api/dtmsvr"

var qsBusi = fmt.Sprintf("http://localhost:%d%s", qsBusiPort, qsBusiAPI)

func TestDtm(t *testing.T) {
	QsStartSvr()
	gid := SagaDtm()
	log.Printf("transaction: %s submitted", gid)
	select {}
}

func QsStartSvr() {
	app := gin.New()
	app.POST(qsBusiAPI+"/TransIn", func(c *gin.Context) {
		log.Printf("TransIn")
		//c.JSON(200, "")
		c.JSON(409, "") // Status 409 for Failure. Won't be retried
	})
	app.POST(qsBusiAPI+"/TransInCompensate", func(c *gin.Context) {
		log.Printf("TransInCompensate")
		c.JSON(200, "")
	})
	app.POST(qsBusiAPI+"/TransOut", func(c *gin.Context) {
		log.Printf("TransOut")
		c.JSON(200, "")
	})
	app.POST(qsBusiAPI+"/TransOutCompensate", func(c *gin.Context) {
		log.Printf("TransOutCompensate")
		c.JSON(200, "")
	})
	go func() {
		_ = app.Run(fmt.Sprintf(":%d", qsBusiPort))
	}()
	time.Sleep(100 * time.Millisecond)
}

func SagaDtm() string {
	req := &gin.H{"amount": 30} // 微服务的载荷
	//整个SAGA事务的逻辑是：
	//执行转出成功=>执行转入成功=>全局事务完成
	//如果在中间发生错误，例如转入B发生错误，则会调用已执行分支的补偿操作，即：
	//执行转出成功=>执行转入失败=>执行转入补偿成功=>执行转出补偿成功=>全局事务回滚完成
	saga := dtmcli.NewSaga(dtmServer, dtmcli.MustGenGid(dtmServer)).
		// 添加一个TransOut的子事务，正向操作为url: qsBusi+"/TransOut"， 逆向操作为url: qsBusi+"/TransOutCompensate"
		Add(qsBusi+"/TransOut", qsBusi+"/TransOutCompensate", req).
		// 添加一个TransIn的子事务，正向操作为url: qsBusi+"/TransOut"， 逆向操作为url: qsBusi+"/TransInCompensate"
		Add(qsBusi+"/TransIn", qsBusi+"/TransInCompensate", req)
	saga.TimeoutToFail = 1800

	// 个事务中的操作分为可回滚的操作，以及不可回滚的操作。那么把可回滚的操作放到前面，把不可回滚的操作放在后面执行，
	//saga := dtmcli.NewSaga(DtmServer, dtmcli.MustGenGid(DtmServer)).
	//	Add(Busi+"/CanRollback1", Busi+"/CanRollback1Revert", req).
	//	Add(Busi+"/CanRollback2", Busi+"/CanRollback2Revert", req).
	//	Add(Busi+"/UnRollback1", "", req).
	//	Add(Busi+"/UnRollback2", "", req).
	//	EnableConcurrent().
	//	AddBranchOrder(2, []int{0, 1}). // 指定step 2，需要在0，1完成后执行
	//	AddBranchOrder(3, []int{0, 1}) // 指定step 3，需要在0，1完成后执行

	err := saga.Submit() // 提交saga事务，dtm会完成所有的子事务/回滚所有的子事务

	if err != nil {
		panic(err)
	}
	return saga.Gid
}

// msg的提交是按照两个阶段发起的
// 第一阶段调用Prepare，
// 第二阶段调用Commit，DTM收到Prepare调用后，不会调用分支事务，而是等待后续的Submit。只有收到了Submit，开始分支调用，最终完成全局事务。
func TwoSubmlitDtm() string {
	return ""
}

// TCC分为3个阶段
// Try 阶段：尝试执行，完成所有业务检查（一致性）, 预留必须业务资源（准隔离性）
// Confirm 阶段：如果所有分支的Try都成功了，则走到Confirm阶段。Confirm真正执行业务，不作任何业务检查，只使用 Try 阶段预留的业务资源
// Cancel 阶段：如果所有分支的Try有一个失败了，则走到Cancel阶段。Cancel释放 Try 阶段预留的业务资源。
func TccDtm() string {
	return ""
}

func XaDtm() string {
	return ""
}
