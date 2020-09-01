package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
	endpoints      = []string{"localhost:2379" /*, "localhost:22379", "localhost:32379"*/}
)

func main() {
	log.Println("Hi, Ed")

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:          endpoints,
		DialTimeout:        dialTimeout,
		MaxCallRecvMsgSize: 4096,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	_, err = cli.Put(context.TODO(), "foo", "bar")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, "foo")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range resp.Kvs {
		log.Printf("%s : %s\n", ev.Key, ev.Value)
	}

	// resp, err = cli.Get(ctx, "/Common/Log")
	log.Println("Starting another one...")

	ctx, cancel = context.WithTimeout(context.Background(), requestTimeout)

	resp, err = cli.Get(ctx, "Common/Log", clientv3.WithLimit(0))
	// resp, err = cli.Get(ctx, "Operations/Agent")
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Resp.Count: %d\n", resp.Count)
	log.Printf("Resp.More: %v\n", resp.More)

	var commonLog CommonLog

	for _, ev := range resp.Kvs {
		log.Printf("ev.Size : %d\n", ev.Size())

		log.Printf("%s : %s\n", ev.Key, string(ev.Value))

		err2 := json.Unmarshal(ev.Value, &commonLog)

		if err2 != nil {
			log.Fatalf("JSON parsing error: %s\n", err2)
		} else {
			// log.Printf("\n------\nObj: %+v\n", commonLog)

			g := commonLog.Writers.(map[string]interface{})

			for k, v := range g {

				innerMap := v.(map[string]interface{})
				log.Printf("\tKey: %s\n", k)

				for kk, vv := range innerMap {
					log.Printf("\t\tKeyKey: %s. Value: %+v\n", kk, vv)
				}

				log.Println("--------")
				// log.Printf("\tKey: %s. Type: %s. Value: %+v\n-----\n", k, reflect.TypeOf(v).String(), v)
			}
		}
	}

	log.Print("\n\n\n---- Going for all keys ----\n\n\n")

	ctx, cancel = context.WithTimeout(context.Background(), requestTimeout)
	resp, err = cli.Get(ctx, " ", clientv3.WithLimit(0))
	// resp, err = cli.Get(ctx, "Operations/Agent")
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Resp.Count: %d\n", resp.Count)
	log.Printf("Resp.More: %v\n", resp.More)
	log.Printf("Header: %s\n", resp.Header.String())

}
