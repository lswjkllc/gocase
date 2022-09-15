package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	STABLE_OID = primitive.NewObjectID().Hex()
	TARGETS    = []string{"smart watch", "apple watch", "huawei watch", "xiaomi watch", "google watch"}
)

type Mention struct {
	Dt        string   `json:"dt"`
	Typ       uint8    `json:"typ"`
	Topicid   string   `json:"topicid"`
	Target    string   `json:"target"`
	Dataid    string   `json:"dataid"`
	Tokens    []string `json:"tokens"`
	Indexes   []string `json:"indexes"`
	Strategy  uint8    `json:"strategy"`
	Platform  uint8    `json:"platform"`
	GmtCreate int64    `json:"gmtCreate"`
}

func randTarget() string {
	return TARGETS[rand.Intn(5)]
}

func generateRandMention(typ uint8, topicid string, subn int) []*Mention {
	target := randTarget()

	data := make([]*Mention, 0, subn)
	for i := 0; i < subn; i++ {
		data = append(data, &Mention{
			Dt:        "2021-01-02 01:02:03",
			Typ:       typ,
			Topicid:   topicid,
			Target:    target,
			Dataid:    primitive.NewObjectID().Hex(),
			Tokens:    []string{"smartwatch", "smart watch", "smart@watch"},
			Indexes:   []string{"1,11", "13,24", "25,36"},
			Strategy:  uint8(rand.Intn(6)),
			Platform:  uint8(rand.Intn(7)),
			GmtCreate: time.Now().UnixMilli(),
		})
	}
	return data
}

func generateRandMentionList(n, subn int, typ uint8, topicid string) []*Mention {
	data := make([]*Mention, 0, n*subn)
	data = append(data, generateRandMention(typ, topicid, subn)...)
	data = append(data, generateRandMention(typ, topicid, subn)...)
	data = append(data, generateRandMention(typ, topicid, subn)...)
	return data
}

func generateStableMention(typ uint8, topicid string, subn int) []*Mention {
	target := randTarget()

	data := make([]*Mention, 0, subn)
	for i := 0; i < subn; i++ {
		data = append(data, &Mention{
			Dt:        "2021-01-02 01:02:03",
			Typ:       typ,
			Topicid:   topicid,
			Target:    target,
			Dataid:    "6302fb43c8511eb28c62c704",
			Tokens:    []string{"smartwatch", "smart watch", "smart@watch"},
			Indexes:   []string{"1,11", "13,24", "25,36"},
			Strategy:  1,
			Platform:  1,
			GmtCreate: time.Now().UnixMilli(),
		})
	}

	return data
}

func generateStableMentionList(n, subn int, typ uint8, topicid string) []*Mention {
	data := make([]*Mention, 0, n*subn)
	data = append(data, generateStableMention(typ, topicid, subn)...)
	data = append(data, generateStableMention(typ, topicid, subn)...)
	data = append(data, generateStableMention(typ, topicid, subn)...)
	return data
}

func generateData() []byte {
	n, subn := 3, 3                       // 总条数: n * subn
	var typ uint8 = 1                     // source/mention
	topicid := "6302fb43c8511eb28c62c704" // 项目ID

	data := generateRandMentionList(n, subn, typ, topicid)
	// data := generateStableMentionList(n, subn, typ, topicid)

	bs, _ := json.Marshal(data)
	return bs
}

func main() {
	topic := "mention"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "ai.wgine-dev.com:32623", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: generateData()},
		kafka.Message{Value: generateData()},
		kafka.Message{Value: generateData()},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
