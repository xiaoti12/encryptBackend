package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/fsnotify/fsnotify"
)

var watcher, _ = fsnotify.NewWatcher()
var done = make(chan bool)

func main() {
	// watcher, err := fsnotify.NewWatcher()

	defer watcher.Close()

	var taskId string
	flag.StringVar(&taskId, "taskid", "", "task ID")
	flag.Parse()
	if taskId == "" {
		log.Println("no task ID argument!")
		flag.Usage()
		os.Exit(1)
	}

	go watchLogFile(taskId)
	go watchDBStatus(taskId)

	workingPath, _ := os.Executable()
	workingPath = filepath.Dir(workingPath)
	watchingPath := filepath.Join(workingPath, "watch")
	log.Println("watching path:", watchingPath)
	err := watcher.Add(watchingPath)
	if err != nil {
		log.Fatal(err)
	}

	<-done
	log.Println("stop watching...")
	os.Exit(0)
}

func watchLogFile(taskId string) {
	log.Println("start watching...")
	tlsPattern := `tls\.(\d+-)+\d+\.log$`
	removePattern := `\.(\d+-)+\d+\.log$`
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				absolutePath, err := filepath.Abs(event.Name)
				matched, err := regexp.MatchString(tlsPattern, event.Name)
				if err != nil {
					log.Printf("regex match error:%s\n", err)
				}
				if !matched {
					// log.Println("not matched!", event.Name)
					removed, _ := regexp.MatchString(removePattern, event.Name)
					if removed {
						os.Remove(absolutePath)
					}
				} else {
					start := time.Now()
					if err != nil {
						log.Printf("cannot get file absolute path:%s\n", err)
						continue
					}
					log.Println("process:", absolutePath)

					// cmd := exec.Command("python", "log_predict.py", "--log", absolutePath, "--taskid", taskId)
					// output, err := cmd.Output()
					// log.Printf("%s", string(output))
					if err != nil {
						log.Printf("error when processing %s: %s\n", absolutePath, err)
					}
					log.Println("Duration:", time.Since(start))
				}
			}
			// if event.Op&fsnotify.Rename == fsnotify.Rename {
			// 	log.Println("rename:", event.Name)
			// }

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}

func watchDBStatus(taskId string) {
	// time.Sleep(10 * time.Second)
	// done <- true
	conn, err := sql.Open("clickhouse", "tcp://10.3.242.84:9000?&compress=true&debug=false&password=password&database=TLS")
	if err != nil {
		log.Printf("failed to connect to clickhouse:%s\n", err)
		return
	}
	defer conn.Close()

	query := "SELECT status FROM Task WHERE taskId = ?"
	var row *sql.Row
	var status int
	ticker := time.NewTicker(5 * time.Second)

	defer ticker.Stop()
	for range ticker.C {
		row = conn.QueryRow(query, taskId)
		err = row.Scan(&status)
		if err != nil {
			log.Println("failed to get task status:", err)
			continue
		}
		if status == 6 {
			done <- true
		}
		// log.Println("status: ", status)
	}

}
