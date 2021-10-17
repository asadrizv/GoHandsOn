package main

import (
	"fmt"

	"github.com/couchbase/gocb"
)

func main() {
	cluster, _ := gocb.Connect("couchbase://localhost")

	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "shikra",
	})

	bucket_, _ := cluster.OpenBucket("beer-sample", "")

	query := gocb.NewN1qlQuery("SELECT * FROM `beer-sample`")

	results, err := bucket_.ExecuteN1qlQuery(query, nil)

	if err != nil {
		fmt.Println(err.Error())

	}
	results.Close()
	fmt.Println(results.Metrics().ExecutionTime)

	var value interface{}
	cas, err1 := bucket_.Get("21st_amendment_brewery_cafe-amendment_pale_ale", &value)

	if err1 != nil {
		fmt.Println(cas)
	}

	for key, val := range value.(map[string]interface{}) {
		switch val.(type) {
		case string:
			fmt.Println(key, val)
		}
	}

}
