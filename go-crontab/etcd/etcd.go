package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

func main() {
	var (
		client *clientv3.Client
		err    error
	)

	config := clientv3.Config{
		Endpoints:   []string{"192.168.50.124:31778"}, //集群列表
		DialTimeout: 5 * time.Second,
	}

	//建立连接，新建一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println("连接成功")
	//用于etcd的键值对
	// kv := clientv3.NewKV(client)
	// // 新建一个租约
	// lease := clientv3.NewLease(client)
	// // 新建一个watcher
	// watcher := clientv3.NewWatcher(client)

	// etcd 常规操作
	// putResp(kv)
	// getResp(kv)
	// getFixResp(kv)
	// deleteResp(kv)

	// 租约，定时续约
	// leaseResp(kv, lease)

	// 服务注册，服务发现，watch监听
	// watcherResp(kv, lease, watcher)

	//etcd 事务处理，任务抢占
	txnResp(client)
}

func putResp(kv clientv3.KV) {
	if putResp, err := kv.Put(context.TODO(), "/corn/jobs/job1", "job1-update", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Revision is: ", putResp.Header.Revision)
		if putResp.PrevKv != nil {
			fmt.Printf("Key: %v, Vaule: %v", string(putResp.PrevKv.Key), string(putResp.PrevKv.Value))
		}
	}
}

func getResp(kv clientv3.KV) {
	if getResp, err := kv.Get(context.TODO(), "/corn/jobs/job1", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Reversion is: ", getResp.Header.Revision)
		// if getResp.Kvs != nil {

		// }
		for k, v := range getResp.Kvs {
			fmt.Println(k, v)

			// 0 key:"/corn/jobs/job1" create_revision:5 mod_revision:12 version:8 value:"job1"
			//create_revision:创建版本
			//mod_revision: 修改版本
			//version:修改了几个版本
		}
	}
}

func getFixResp(kv clientv3.KV) {
	if getResp, err := kv.Get(context.TODO(), "/corn/jobs/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
	} else {
		for k, v := range getResp.Kvs {
			fmt.Println(k, v)
		}
	}
}

func deleteResp(kv clientv3.KV) {
	if deleteResp, err := kv.Delete(context.TODO(), "/corn/jobs/", clientv3.WithFromKey()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(deleteResp.Header.Revision)
		fmt.Println(deleteResp.Deleted)
		for k, v := range deleteResp.PrevKvs {
			fmt.Println(k, v)
		}
	}

}

func leaseResp(kv clientv3.KV, lease clientv3.Lease) {
	var (
		leaseGrantResp *clientv3.LeaseGrantResponse
		err            error
		keepResp       *clientv3.LeaseKeepAliveResponse
		putResp        *clientv3.PutResponse
		getResp        *clientv3.GetResponse
	)
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}
	leaseID := leaseGrantResp.ID

	KeepRespChan, err := lease.KeepAlive(context.TODO(), leaseID)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			select {
			case keepResp = <-KeepRespChan:
				if keepResp == nil {
					fmt.Println("租约已经失效")
					goto END
				} else {
					fmt.Println("收到自动续约应答：", keepResp.ID)
				}
			}
		}
	END:
	}()
	if putResp, err = kv.Put(context.TODO(), "/corn/lease/", "test", clientv3.WithLease(leaseID)); err != nil {
		fmt.Println("err")
		return
	}
	fmt.Println("写入成功:", putResp.Header.Revision)

	// 定时查看key过期没
	for {
		if getResp, err = kv.Get(context.TODO(), "/corn/lease/", clientv3.WithPrevKV()); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		}
		fmt.Println("还没过期", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}

}

func watcherResp(kv clientv3.KV, lease clientv3.Lease, watcher clientv3.Watcher) {
	var (
		getResp   *clientv3.GetResponse
		err       error
		watchResp clientv3.WatchResponse
	)
	//模拟etcd中的kv的变化
	go func() {
		for {
			kv.Put(context.TODO(), "/corn/watch/", "test", clientv3.WithPrevKV())
			kv.Delete(context.TODO(), "/corn/watch/", clientv3.WithPrevKV())
			time.Sleep(1 * time.Second)
		}
	}()

	// 先 get 到当前值，后续监听
	if getResp, err = kv.Get(context.TODO(), "/corn/watch/", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("获取到当前值：", string(getResp.Kvs[0].Value))
	}
	watcherStartReversion := getResp.Header.Revision + 1

	//启动监听
	fmt.Println("从该版本向后监听: ", watcherStartReversion)

	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(15*time.Second, func() {
		cancelFunc()
	})

	// 开始监听
	watchRespChan := watcher.Watch(ctx, "/corn/watch/", clientv3.WithRev(watcherStartReversion))
	//处理kv变化事件
	for watchResp = range watchRespChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为:", string(event.Kv.Value), "Revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了:", "Revision:", event.Kv.ModRevision)
			}
		}
	}
}

func txnResp(client *clientv3.Client) {
	var (
		kv    clientv3.KV
		lease clientv3.Lease
		txn   clientv3.Txn

		// ctx context.Context

		leaseGrantResp     *clientv3.LeaseGrantResponse
		leaseKeepAliveResp <-chan *clientv3.LeaseKeepAliveResponse
		keepResp           *clientv3.LeaseKeepAliveResponse
		txnResp            *clientv3.TxnResponse
		err                error
	)

	kv = clientv3.NewKV(client)

	//1, 上锁 (创建租约，自动续租，拿着租约去抢占一个 key)
	lease = clientv3.NewLease(client)

	// 准备一个用于取消自动续租的 context
	ctx, cancelFunc := context.WithCancel(context.TODO())

	// 申请 1 个 5 秒的租约
	if leaseGrantResp, err = lease.Grant(ctx, 5); err != nil {
		fmt.Println(err)
		return
	}

	// 拿到租约 ID
	leaseID := leaseGrantResp.ID

	//5 秒后会取消租约  续租
	if leaseKeepAliveResp, err = lease.KeepAlive(context.TODO(), leaseID); err != nil {
		fmt.Println(err)
		return
	}

	// 处理续约应答的协程
	go func() {
		for {
			select {
			case keepResp = <-leaseKeepAliveResp:
				if keepResp == nil {
					fmt.Println("租约已经失效")
					goto END
				} else {
					fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " 收到自动续租应答:", keepResp.ID)
				}
			}
		}
	END:
	}()

	// 创建事务
	txn = kv.Txn(context.TODO())

	// 定义事务
	//if 不存在 key, then 设置它， else 抢锁失败
	txn.If(clientv3.Compare(clientv3.CreateRevision("/corn/job/test"), "=", 0)).
		Then(clientv3.OpPut("/corn/job/test", "test", clientv3.WithLease(leaseID))).
		Else(clientv3.OpGet("/corn/job/test", clientv3.WithPrevKV()))

	// 提交事务
	if txnResp, err = txn.Commit(); err != nil {
		fmt.Println(err)
		return
	}

	// 判断是否抢到了锁
	if !txnResp.Succeeded {
		fmt.Println(" 锁被占用:", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	//2, 处理业务
	fmt.Println(" 处理任务 ")
	time.Sleep(50 * time.Second)

	//3, 释放锁（取消自动续租，释放租约
	//defer 会把租约释放掉，关联的 KV 就被删除了
	// 确保函数退出后，自动续租会停止
	defer cancelFunc()
	defer lease.Revoke(context.TODO(), leaseID)
}
