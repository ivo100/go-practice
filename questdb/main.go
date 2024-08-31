package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	qdb "github.com/questdb/go-questdb-client/v3"
	"log"
	"time"
)

/*

https://questdb.io/docs/reference/api/postgres/

https://github.com/questdb/go-questdb-client/blob/v3.2.0/README.md

https://questdb.io/docs/guides/enterprise-quick-start/#0-secure-the-built-in-admin

/opt/homebrew/var/questdb
less  /opt/homebrew/var/questdb/conf/server.conf
grep admin  /opt/homebrew/var/questdb/conf/server.conf
#pg.user=admin
#acl.admin.user.enabled=true
#acl.admin.user=admin
#acl.admin.password=quest


# the built-in admin's user name and password
acl.admin.user=myadmin
acl.admin.password=my_very_secure_pwd

	--drop table 'trades_go'

	  CREATE TABLE 'trades_go' (
	    pair SYMBOL capacity 256 CACHE,
	    type SYMBOL capacity 256 CACHE,
	    traded_price DOUBLE,
	    limit_price DOUBLE,
	    qty LONG,
	    timestamp TIMESTAMP
	  ) timestamp (timestamp) PARTITION BY DAY WAL;

-- select * from 'trades_go'
--drop table 'trades_go'
CREATE TABLE 'trades' (ts TIMESTAMP, date DATE, name STRING, value INT) timestamp (ts);
*/

var conn *pgx.Conn
var err error

func main() {
	pg()
	//write()
}

func pg() {
	ctx := context.Background()
	//conn, _ = pgx.Connect(ctx, "postgresql://admin:quest@localhost:8812/qdb")
	conn, _ = pgx.Connect(ctx, "postgresql://admin:quest@localhost:8812/qdb")
	defer conn.Close(ctx)

	// text-based query
	_, err := conn.Exec(ctx,
		"CREATE TABLE IF NOT EXISTS trades ("+
			"    ts TIMESTAMP, date DATE, name STRING, value INT"+
			") timestamp(ts);")
	if err != nil {
		log.Fatalln(err)
	}

	// Prepared statement given the name 'ps1'
	_, err = conn.Prepare(ctx, "ps1", "INSERT INTO trades VALUES($1,$2,$3,$4)")
	if err != nil {
		log.Fatalln(err)
	}

	// Insert all rows in a single commit
	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 10; i++ {
		// Execute 'ps1' statement with a string and the loop iterator value
		_, err = conn.Exec(
			ctx,
			"ps1",
			time.Now().UTC(),
			time.Now().Round(time.Millisecond),
			"go prepared statement",
			i+1)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Read all rows from table
	rows, err := conn.Query(ctx, "SELECT * FROM trades")
	fmt.Println("Reading from trades table:")
	for rows.Next() {
		var name string
		var value int64
		var ts time.Time
		var date time.Time
		err = rows.Scan(&ts, &date, &name, &value)
		fmt.Println(ts, date, name, value)
	}
	//err = conn.Close(ctx)
}
func write() {
	ctx := context.TODO()
	// Connect to QuestDB running locally.
	/*
		name: "auto flush",
		config: fmt.Sprintf("http::addr=%s;auto_flush_rows=100;auto_flush_interval=1000;",
				addr),
	*/
	sender, err := qdb.LineSenderFromConf(ctx, "http::addr=localhost:9000;")
	if err != nil {
		log.Fatal(err)
	}
	// Make sure to close the sender on exit to release resources.
	defer sender.Close(ctx)

	// Send a few ILP messages.
	tradedTs, err := time.Parse(time.RFC3339, "2024-08-06T15:04:05.123456Z")
	if err != nil {
		log.Fatal(err)
	}
	err = sender.
		Table("trades_go").
		Symbol("pair", "USDGBP").
		Symbol("type", "buy").
		Float64Column("traded_price", 0.83).
		Float64Column("limit_price", 0.84).
		Int64Column("qty", 100).
		At(ctx, tradedTs)
	if err != nil {
		log.Fatal(err)
	}

	tradedTs, err = time.Parse(time.RFC3339, "2024-08-06T15:04:06.987654Z")
	if err != nil {
		log.Fatal(err)
	}
	err = sender.
		Table("trades_go").
		Symbol("pair", "GBPJPY").
		Symbol("type", "sell").
		Float64Column("traded_price", 135.97).
		Float64Column("limit_price", 0.84).
		Int64Column("qty", 400).
		At(ctx, tradedTs)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure that the messages are sent over the network.
	err = sender.Flush(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
