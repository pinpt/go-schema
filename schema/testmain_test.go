// GENERATED CODE. DO NOT EDIT

package schema

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jhaynie/go-gator/orm"
)

var (
	database string
	username string
	password string
	hostname string
	port     int
	db       DB
	createdb = true
)

func init() {
	var defuser = "root"
	var defdb = fmt.Sprintf("test_%s", orm.UUID()[0:9])
	flag.StringVar(&username, "username", defuser, "database username")
	flag.StringVar(&password, "password", "", "database password")
	flag.StringVar(&hostname, "hostname", "localhost", "database hostname")
	flag.IntVar(&port, "port", 3306, "database port")
	database = defdb
}

func GetDatabase() DB {
	return db
}

func GetDSN(name string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, hostname, port, name)
}

func openDB(name string) DB {
	dsn := GetDSN(name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return db
}

func dropDB() {
	if createdb {
		_, err := db.ExecContext(context.Background(), fmt.Sprintf("drop database %s", database))
		if err != nil {
			fmt.Printf("error dropping database named %s\n", database)
		}
	}
}

func ToTimestampNow() *timestamp.Timestamp {
	t := time.Now()
	// truncate to 24 hours
	t = t.Truncate(time.Hour * 24)
	ts, _ := ptypes.TimestampProto(t)
	// since mysql is only second precision we truncate
	ts.Nanos = 0
	return ts
}

func TestMain(m *testing.M) {
	flag.Parse()
	if createdb {
		// open without a database so we can create a temp one
		d := openDB("")
		_, err := d.ExecContext(context.Background(), fmt.Sprintf("create database %s", database))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		d.Close()
	}
	// reopen now with the temp database
	db = openDB(database)
	x := m.Run()
	dropDB()
	db.Close()
	os.Exit(x)
}
