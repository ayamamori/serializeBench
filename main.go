package main

import (
	"bytes"
	"testing"

	"log"

	"github.com/ayamamori/serializeBench/messagePackData"
	"github.com/ayamamori/serializeBench/protoData"
	"github.com/golang/protobuf/proto"
	"github.com/vmihailenco/msgpack"
)

const NUM_OF_CARDS = 100
const NUM_OF_REPEAT = 1000
const NS_TO_MS = 1000

func main() {
	protoBench()
	messagePackBench()
}

func protoBench() {
	userCardList := make([]*protoData.UserCard, NUM_OF_CARDS)
	for i := 0; i < NUM_OF_CARDS; i++ {
		userCardList[i] = &protoData.UserCard{CardId: int32(i * 33), Level: int32(i * 3333)}
	}
	responseTop := &protoData.ResponseTop{
		Ts:  111111111,
		Pid: 111111111,
		Rev: 111111111,
		Login: &protoData.Login{
			UserStatus: &protoData.UserStatus{
				UserId:   22222,
				UserName: "Nishimura",
				Exp:      22222222,
			},
			UserCardList: userCardList,
		},
	}
	testResultMarshal := testing.Benchmark(func(b *testing.B) { ProtoBenchMarshal(responseTop) })
	log.Printf("protobuf Marshal Time:%d[ms]", testResultMarshal.T.Nanoseconds()/NUM_OF_REPEAT/NS_TO_MS) //[ns] -> [ms]
	marshalData, _ := proto.Marshal(responseTop)
	testResultUnmarshal := testing.Benchmark(func(b *testing.B) { ProtoBenchUnmarshal(marshalData) })
	log.Printf("protobuf Unmarshal Time:%d[ms]", testResultUnmarshal.T.Nanoseconds()/NUM_OF_REPEAT/NS_TO_MS)
}

func ProtoBenchMarshal(responseTop *protoData.ResponseTop) {
	for i := 0; i < NUM_OF_REPEAT; i++ {
		proto.Marshal(responseTop)
	}
}

func ProtoBenchUnmarshal(marshalData []byte) {
	responseTop2 := &protoData.ResponseTop{}
	for i := 0; i < NUM_OF_REPEAT; i++ {
		proto.Unmarshal(marshalData, responseTop2)
	}

}

func messagePackBench() {
	userCardList := make([]messagePackData.UserCard, NUM_OF_CARDS)
	for i := 0; i < NUM_OF_CARDS; i++ {
		userCardList[i] = messagePackData.UserCard{CardId: i * 33, Level: i * 3333}
	}
	responseTop := &messagePackData.ResponseTop{
		Ts:  111111111,
		Pid: 111111111,
		Rev: 111111111,
		Login: messagePackData.Login{
			UserStatus: messagePackData.UserStatus{
				UserId:   22222,
				UserName: "Nishimura",
				Exp:      22222222,
			},
			UserCardList: userCardList,
		},
	}

	testResultMarshal := testing.Benchmark(func(b *testing.B) { MessagePackBenchMarshal(responseTop) })
	log.Printf("msgpack Marshal Time:%d[ms]", testResultMarshal.T.Nanoseconds()/NUM_OF_REPEAT/NS_TO_MS)
	marshalData, _ := msgpack.Marshal(responseTop)
	testResultUnmarshal := testing.Benchmark(func(b *testing.B) { MessagePackBenchUnmarshal(marshalData) })
	log.Printf("msgpack Unmarshal Time:%d[ms]", testResultUnmarshal.T.Nanoseconds()/NUM_OF_REPEAT/NS_TO_MS)
}

func MessagePackBenchMarshal(responseTop *messagePackData.ResponseTop) {
	var buf bytes.Buffer
	encoder := msgpack.NewEncoder(&buf)
	for i := 0; i < NUM_OF_REPEAT; i++ {
		encoder.Encode(responseTop)
	}
}

func MessagePackBenchUnmarshal(marshalData []byte) {
	responseTop2 := &messagePackData.ResponseTop{}
	for i := 0; i < NUM_OF_REPEAT; i++ {
		msgpack.Unmarshal(marshalData, responseTop2)
	}

}
