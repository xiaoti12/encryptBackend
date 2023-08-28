package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"syscall"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/fsnotify/fsnotify"
)

var watcher, _ = fsnotify.NewWatcher()
var allDone = make(chan bool)
var zeekDone = make(chan bool)

func main() {
	// watcher, err := fsnotify.NewWatcher()

	defer watcher.Close()

	var taskId string
	var netCard string
	flag.StringVar(&taskId, "taskid", "", "task ID")
	flag.StringVar(&netCard, "netcard", "", "network card name")
	flag.Parse()
	if taskId == "" || netCard == "" {
		log.Println("lack of arguments!")
		flag.Usage()
		os.Exit(1)
	}

	go runZeek(netCard)
	go watchLogFile(taskId)
	go watchDBStatus(taskId)

	// 将工作/监视目录设为当前目录下的watch文件夹
	workingPath, _ := os.Executable()
	workingPath = filepath.Dir(workingPath)
	watchingPath := filepath.Join(workingPath, "watch")
	os.Chdir(workingPath)
	fmt.Println("watching path:", watchingPath)
	err := watcher.Add(watchingPath)
	if err != nil {
		log.Fatal(err)
	}

	<-allDone
	fmt.Println("stop watching...")
	os.Exit(0)
}

func runZeek(netcard string) {
	cmd := exec.Command("zeek", "-i", netcard, "../tls.zeek")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("error when processing %s: %s\n", netcard, err)
	}
	<-zeekDone
	cmd.Process.Signal(syscall.SIGTERM)
	fmt.Println("stop zeek...")
	allDone <- true
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
						fmt.Printf("cannot get file absolute path:%s\n", err)
						continue
					}
					fmt.Println("process:", absolutePath)

					cmd := exec.Command("python3", "log_predict.py", "--log", absolutePath, "--taskid", taskId)
					output, err := cmd.Output()
					fmt.Printf("%s", string(output))
					if err != nil {
						fmt.Printf("error when processing %s: %s\n", absolutePath, err)
					}
					fmt.Println("Duration:", time.Since(start))
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
	conn, err := sql.Open("clickhouse", "tcp://localhost:9000?&compress=true&debug=false&password=password&database=TLS")
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
			zeekDone <- true
		}
		// log.Println("status: ", status)
	}

}
