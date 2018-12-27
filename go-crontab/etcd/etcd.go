package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
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
	kv := clientv3.NewKV(client)
	// 新建一个租约
	lease := clientv3.NewLease(client)

	// op etcd
	// putResp(kv)
	// getResp(kv)
	// getFixResp(kv)
	// deleteResp(kv)
	leaseResp(kv, lease)
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
