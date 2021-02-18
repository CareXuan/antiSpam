package workers

import (
	"antispam/base"
	"go.mongodb.org/mongo-driver/bson"
)

func AllRequestWorker() {
	go contentRequest()
	go contentResponse()
}

func contentRequest() {
	CReqDealCount := 10
	CReqDealChan := make(chan bool, CReqDealCount)
	for {
		select {
		case CReqDealChan <- true:
			go func() {
				defer func() { <-CReqDealChan }()
				requestJson := <-base.Conf.ContentRequestChan
				base.AddMongoOne("carexuan", "test", bson.M{"request": requestJson})
			}()
		}
	}
}

func contentResponse() {
	CRepDealCount := 10
	CRepDealChan := make(chan bool, CRepDealCount)
	for {
		select {
		case CRepDealChan <- true:
			go func() {
				defer func() { <-CRepDealChan }()
				responseJson := <-base.Conf.ContentResponseChan
				base.AddMongoOne("carexuan", "test", bson.M{"response": responseJson})
			}()
		}
	}
}
