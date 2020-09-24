package driver

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	// "github.com/lib/pq"
	_ "github.com/denisenkom/go-mssqldb"
)

var (
	debug         = flag.Bool("debug", true, "enable debugging")
	password      = flag.String("password", os.Getenv("DB_USER_PASS"), "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", os.Getenv("DB_SERVER"), "the database server")
	database      = flag.String("database", os.Getenv("DB_DATABASE"), "the database name")
	user          = flag.String("user", os.Getenv("DB_USER_ID"), "the database user")
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" database:%s\n", *database)
		fmt.Printf(" user:%s\n", *user)
	}

	*password = os.Getenv("DB_USER_PASS")
	*port = 1433
	*server = os.Getenv("DB_SERVER")
	*database = os.Getenv("DB_DATABASE")
	*user = os.Getenv("DB_USER_ID")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", *server, *user, *password, *port, *database)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	return conn
	// pgURL, err := pq.ParseURL(os.Getenv("SQLSERVER_DSN"))
	// logFatal(err)

	// db, err = sql.Open("postgres", pgURL)
	// logFatal(err)

	// err = db.Ping()
	// logFatal(err)

	// log.Println(pgURL)

	// return db
}
