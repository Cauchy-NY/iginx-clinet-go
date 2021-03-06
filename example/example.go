package main

import (
	"fmt"
	"log"
	"math"

	"github.com/Cauchy-NY/iginx-clinet-go/client"
	"github.com/Cauchy-NY/iginx-clinet-go/rpc"
)

var (
	session *client.Session

	s1 = "test.go.a"
	s2 = "test.go.b"
	s3 = "test.go.c"
	s4 = "test.go.d"
	s5 = "test.go.e"
	s6 = "test.go.f"
)

func main() {
	session = client.NewSession("127.0.0.1", "6888", "root", "root")

	if err := session.Open(); err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	showReplicaNum()
	showClusterInfo()

	insertRowData()
	insertNonAlignedRowRecords()
	insertColumnData()
	insertNonAlignedColumnRecords()

	showTimeSeries()

	queryAllData()
	valueFilterQuery()
	downSampleQuery()
	aggregateQuery()
	lastQuery()
}

func showReplicaNum() {
	num, err := session.GetReplicaNum()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("replica num: %d\n", num)
}

func showClusterInfo() {
	info, err := session.GetClusterInfo()
	if err != nil {
		log.Fatal(err)
	}
	info.PrintInfo()
}

func insertRowData() {
	path := []string{s1, s2, s3, s4, s5, s6}
	timestamps := []int64{1, 2, 3, 4, 5, 6, 7}
	values := [][]interface{}{
		{"one", int32(1), int64(1), float32(1.1), float64(1.1), true},
		{"two", int32(2), int64(2), float32(2.1), float64(2.1), false},
		{"three", nil, int64(3), float32(3.1), float64(3.1), true},
		{"four", int32(4), nil, float32(4.1), float64(4.1), false},
		{"five", int32(5), int64(5), nil, float64(5.1), true},
		{"six", int32(6), int64(6), float32(6.1), nil, false},
		{"seven", int32(7), int64(7), float32(7.1), float64(7.1), nil},
	}
	types := []rpc.DataType{rpc.DataType_BINARY, rpc.DataType_INTEGER, rpc.DataType_LONG, rpc.DataType_FLOAT, rpc.DataType_DOUBLE, rpc.DataType_BOOLEAN}
	err := session.InsertRowRecords(path, timestamps, values, types)
	if err != nil {
		log.Fatal(err)
	}
}

func insertNonAlignedRowRecords() {
	path := []string{s1, s2, s3, s4, s5, s6}
	timestamps := []int64{8, 9}
	values := [][]interface{}{
		{"one", int32(8), int64(8), float32(8.1), float64(8.1), false},
		{"two", int32(9), int64(9), float32(9.1), float64(9.1), true},
	}
	types := []rpc.DataType{rpc.DataType_BINARY, rpc.DataType_INTEGER, rpc.DataType_LONG, rpc.DataType_FLOAT, rpc.DataType_DOUBLE, rpc.DataType_BOOLEAN}
	err := session.InsertNonAlignedRowRecords(path, timestamps, values, types)
	if err != nil {
		log.Fatal(err)
	}
}

func insertColumnData() {
	path := []string{s1, s2, s3, s4, s5, s6}
	timestamps := []int64{10, 11}
	values := [][]interface{}{
		{"ten", "eleven"},
		{int32(10), int32(11)},
		{int64(10), int64(11)},
		{float32(10.1), float32(11.1)},
		{float64(10.1), float64(11.1)},
		{false, true},
	}
	types := []rpc.DataType{rpc.DataType_BINARY, rpc.DataType_INTEGER, rpc.DataType_LONG, rpc.DataType_FLOAT, rpc.DataType_DOUBLE, rpc.DataType_BOOLEAN}
	err := session.InsertColumnRecords(path, timestamps, values, types)
	if err != nil {
		log.Fatal(err)
	}
}

func insertNonAlignedColumnRecords() {
	paths := []string{s1, s2, s3, s4, s5, s6}
	timestamps := []int64{12, 13}
	values := [][]interface{}{
		{"twelve", "thirteen"},
		{int32(12), int32(13)},
		{int64(12), int64(13)},
		{float32(12.1), float32(13.1)},
		{float64(12.1), float64(13.1)},
		{false, true},
	}
	types := []rpc.DataType{rpc.DataType_BINARY, rpc.DataType_INTEGER, rpc.DataType_LONG, rpc.DataType_FLOAT, rpc.DataType_DOUBLE, rpc.DataType_BOOLEAN}
	err := session.InsertNonAlignedColumnRecords(paths, timestamps, values, types)
	if err != nil {
		log.Fatal(err)
	}
}

func showTimeSeries() {
	fmt.Println("show time series:")
	tsList, err := session.ListTimeSeries()
	if err != nil {
		log.Fatal(err)
	}
	for _, ts := range tsList {
		fmt.Println(ts.ToString())
	}
}

func queryAllData() {
	fmt.Println("query all data:")
	paths := []string{s1, s2, s3, s4, s5, s6}
	dataSet, err := session.Query(paths, 0, math.MaxInt64)
	if err != nil {
		log.Fatal(err)
	}
	dataSet.PrintDataSet()
}

func valueFilterQuery() {
	fmt.Println("value filter query:")
	paths := []string{s1, s2, s3, s4, s5, s6}
	expression := s2 + " > 6" + " && " + s3 + " < 9"
	dataSet, err := session.ValueFilterQuery(paths, 0, math.MaxInt64, expression)
	if err != nil {
		log.Fatal(err)
	}
	dataSet.PrintDataSet()
}

func downSampleQuery() {
	fmt.Println("downSample query:")
	paths := []string{s2, s5}
	dataSet, err := session.DownSampleQuery(paths, 0, 10, rpc.AggregateType_MAX, 5)
	if err != nil {
		log.Fatal(err)
	}
	dataSet.PrintDataSet()
}

func aggregateQuery() {
	fmt.Println("aggregate query:")
	paths := []string{s1, s2}
	dataSet, err := session.AggregateQuery(paths, 0, 10, rpc.AggregateType_MAX)
	if err != nil {
		log.Fatal(err)
	}
	dataSet.PrintDataSet()
}

func lastQuery() {
	fmt.Println("last query:")
	paths := []string{s1, s2, s3}
	dataSet, err := session.LastQuery(paths, 5)
	if err != nil {
		log.Fatal(err)
	}
	dataSet.PrintDataSet()
}
