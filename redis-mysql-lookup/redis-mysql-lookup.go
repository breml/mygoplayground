// Queries Redis for a Key, if not present in Redis, fallback to query MySQL DB

// In redis-cli use "FLUSHALL" to remove all saved keys
// Performance-Tests:
// MySQL only is half as fast as redis, if key is found in redis
// if the key is not found in redis, time is near factor 2
// Results:
// MySQL only: ~	320ms
// Warming Redis by MySQL: ~520ms
// Query Redis: ~125ms
package main

import (
	"flag"
	//"fmt"
	// "os"
	"database/sql"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var id int
	var use_redis *bool = flag.Bool("r", false, "use redis cache")

	flag.Parse()

	// Redis Connection Part
	var err error
	var c redis.Conn
	if *use_redis {
		c, err = redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			panic(err)
		}
		defer c.Close()
	}

	// MySQL Connection Part
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	for id = 0; id < 500; id++ {
		var command string

		if *use_redis {
			command, err = redis.String(c.Do("GET", id))
		}
		if err != nil || !*use_redis {
			// Open doesn't open a connection. Validate DSN data:
			err = db.Ping()
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}

			err = db.QueryRow("SELECT name FROM help_topic WHERE help_topic_id = ?", id).Scan(&command)
			switch {
			case err == sql.ErrNoRows:
				panic("No command with that ID.")
			case err != nil:
				panic(err)
			default:
				//fmt.Println("Found in MySQL")
				if *use_redis {
					c.Do("SET", id, command)
				}
			}
		}

		//fmt.Printf("Command is %s\n", command)
	}
}
