package scylladb

import (
	"fmt"
	"os"
	"strings"

	utils "lss_go/lib/utils"
	// _ "github.com/joho/godotenv/autoload"
	"github.com/gocql/gocql"
)

var session *gocql.Session

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Println("%s: %s", msg, err)
		// log.Panicf("%s: %s", msg, err)
	}
}

func SetKeySpace(keyspace string) {

	use_keyspace := fmt.Sprintf("USE %s", keyspace)
	err := session.Query(use_keyspace).Exec()
	utils.FailOnError(err, "Failed to switch Keyspace")

}

func Connect() {
	nodesStr := os.Getenv("SCYLLA_NODES")
	nodes := strings.Split(nodesStr, ",")

	fmt.Println(nodesStr)

	// 初始化連線1
	cluster := gocql.NewCluster(nodes...)
	cluster.Keyspace = os.Getenv("SCYLLA_DEFAULT_KEYSPACE")

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Getenv("SCYLLA_USERNAME"), // 替換為你的使用者名稱
		Password: os.Getenv("SCYLLA_PASSWORD"), // 替換為你的密碼
	}

	var err error
	session, err = cluster.CreateSession()
	utils.FailOnError(err, "Failed to connect to ScyllaDB")

	fmt.Println("ok")
}
func Close() {
	session.Close()
}
