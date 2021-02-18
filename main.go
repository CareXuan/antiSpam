package main

import (
	"antispam/base"
	"antispam/http"
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	key string
	//连接信息
	Mqurl string
}

func main() {
	err := base.Init("./conf/local.yaml")
	if err != nil {
		fmt.Print(err)
	}
	//rabbitmq := &RabbitMQ{QueueName: "carexuan_test", Exchange: "carexuan_exchange", key: "testkey", Mqurl: "amqp://guest:guest@localhost:5672/"}
	////创建RabbitMQ连接
	//rabbitmq.conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	//if err != nil {
	//	fmt.Print(err)
	//}
	//channel, _ := rabbitmq.conn.Channel()
	//_, err = channel.QueueDeclare(
	//	"carexuan_test",
	//	false,
	//	false,
	//	false,
	//	false,
	//	nil,
	//)
	//if err != nil {
	//	fmt.Print(err)
	//}
	//err = channel.Publish(
	//	"",
	//	"carexuan_test",
	//	false,
	//	false,
	//	amqp.Publishing{
	//		ContentType: "text/plain",
	//		Body:        []byte("哈哈哈哈"),
	//	},
	//)
	//if err != nil {
	//	fmt.Print(err)
	//}
	////msg, _, err := channel.Get("carexuan_test", true)
	////if err != nil {
	////	fmt.Print(err)
	////}
	//////fmt.Print(ok)
	//////fmt.Print("\n")
	////fmt.Print(string(msg.Body))
	//forever := make(chan bool)
	//msgs, err := channel.Consume("carexuan_test", "", false, false, false, false, nil)
	//go func() {
	//	for i := range msgs {
	//		fmt.Print(string(i.Body))
	//		i.Ack(true)
	//	}
	//	<-forever
	//}()
	//client := base.Conf.MongoDB
	//collection := client.Database("carexuan").Collection("test")
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//res, err := collection.InsertOne(ctx, bson.M{"aaa": "bbb"})
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Print(res)
	//client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://root:123456@localhost:27017"))
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//err = client.Connect(ctx)
	//if err != nil {
	//	fmt.Print(err)
	//}
	//collect := base.GetMongoCollection("carexuan", "test")
	//res, err := collect.InsertOne(base.Conf.Ctx, bson.M{"name": "carexuan759", "value": "7595991314"})
	//base.UpdateOne("carexuan", "test", bson.M{"unique_id": "759599131415"}, bson.M{"aaa": "ccc"})
	//res, err := collect.UpdateOne(base.Conf.Ctx, bson.M{"unique_id": "759599131415"}, bson.M{"$set": bson.M{"aaa": "bbb"}})
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Print(res)
	//result, err := src.ContentCheckSecondStep("裸聊", "759599", "PORN_AD")
	//re, err := videoSpam.VideoContentResult("7595991314")
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Print(re)
	r := http.InitGin()
	r.Run(":1234")
}

func mysqlTest() {
	//err := base.Init("./conf/local.yaml")
	//if err != nil {
	//	fmt.Print(err)
	//}
	//mysqlConn := conf.Mysql
	//defer mysqlConn.Close()
	//sql := "select * from t1"
	//rows, err := mysqlConn.Query(sql)
	//if err != nil {
	//	fmt.Print(err)
	//}
	//var T1s []T1
	//for rows.Next() {
	//	var id int64
	//	var val int64
	//	var t1 T1
	//	rows.Scan(&id, &val)
	//	t1.Id = id
	//	t1.Val = val
	//	T1s = append(T1s, t1)
	//}
	//fmt.Print(T1s)
}
