package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func main() {

	runMarsServer()

}

type MarsServer struct {
	server        http.Server
	db            *sql.DB
	sc            stan.Conn
	reportIdQueue chan bool
	// qouta         int
}

func (srv *MarsServer) newReport(w http.ResponseWriter, r *http.Request) {
	if len(srv.reportIdQueue) == 0 {
		srv.reportIdQueue <- true
	}
}

type RequestFromNats struct {
	ClusterCounter string `json:"cluster_counter`
	Report         Report `json:"report"`
}

type Report struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Body     string `json:"body"`
	FileName string `json:"file_name"`
}

func (srv *MarsServer) sendData() {
	idStart, err := srv.db.Query("select id from report where status = 1 limit 1")
	var lastSent int
	if err != nil {
		log.Fatal(err)
	}
	defer idStart.Close()
	for idStart.Next() {
		if err := idStart.Scan(&lastSent); err != nil {
			log.Fatal(err)
		}
	}

	idEnd, err := srv.db.Query("select id from report where status = 1 limit 1")
	var lastReceived int
	if err != nil {
		log.Fatal(err)
	}
	defer idEnd.Close()
	for idEnd.Next() {
		if err := idEnd.Scan(&lastReceived); err != nil {
			log.Fatal(err)
		}
	}
	if lastSent+1 > lastReceived {
		<-srv.reportIdQueue
		return
	}

	rows, err := srv.db.Query("select * from report where id = " + strconv.Itoa(lastSent+1))
	if err != nil {
		log.Fatal(err) //TODO переделять в что-то менее фатальное
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var reportText string
		var name string
		var path string
		var status int

		if err := rows.Scan(&id, &reportText, &name, &path, &status); err != nil {
			log.Fatal(err)
		}

		file, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		fileInfo, _ := file.Stat()

		var fileSize int64 = fileInfo.Size()

		const fileChunk = 1 * (1 << 20) // 1 MB, change this to your requirement

		totalPartsNum := int(math.Ceil(float64(fileSize) / float64(fileChunk)))

		reportRequest := RequestFromNats{
			ClusterCounter: strconv.Itoa(totalPartsNum),
			Report: Report{
				Id:       id,
				Name:     name,
				Body:     reportText,
				FileName: path,
			},
		}
		reportRequestJson, err := json.Marshal(reportRequest)
		if err != nil {
			panic(err)
		}

		err = srv.sc.Publish("foo", reportRequestJson)
		if err != nil {
			log.Fatalf("Error during publish: %v\n", err)
		}

		for i := 0; i < totalPartsNum; i++ {
			partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
			partBuffer := make([]byte, partSize)

			file.Read(partBuffer)

			err = srv.sc.Publish("foo", partBuffer)
			if err != nil {
				log.Fatalf("Error during publish: %v\n", err)
			}
		}

		// TODO: апдейтнуть статус
	}
}

func runMarsServer() {
	db, err := sql.Open("postgres", "user=tm_admin password=admin dbname=mars sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	nc, err := nats.Connect("0.0.0.0:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	sc, err := stan.Connect("test-cluster", "mars", stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, "0.0.0.0:4222")
	}
	defer sc.Close()

	marsSrv := &MarsServer{
		db:            db,
		sc:            sc,
		reportIdQueue: make(chan bool),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/inform", marsSrv.newReport)
	marsSrv.server = http.Server{
		Handler: mux,
		Addr:    ":4000",
	}

	signalChan := make(chan os.Signal, 1)
	done := make(chan bool)
	go func() {

		go func() {
			for {
				marsSrv.sendData()
			}
		}()
		log.Println("Server run on: http:localhost:4000")
		err := marsSrv.server.ListenAndServe()
		fmt.Println(err)

		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
		}
		done <- true
	}()
	<-done
}
