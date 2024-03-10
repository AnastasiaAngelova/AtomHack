package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

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
	qouta         int
	periods       []Period
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
	gotOne := false
	for idStart.Next() {
		if err := idStart.Scan(&lastSent); err != nil {
			log.Fatal(err)
		}
		gotOne = true
		// fmt.Printf("lastSent:%d\n", lastSent)
	}
	if !gotOne {
		<-srv.reportIdQueue
		return
	}

	rows, err := srv.db.Query("select * from report where id = " + strconv.Itoa(lastSent))
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
		// fmt.Printf("id:%d status: %d\n", id, status)

		file, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		fileInfo, _ := file.Stat()

		var fileSize int64 = fileInfo.Size()

		const fileChunk = 1 * (1 << 17) // 1 Mbit, change this to your requirement

		totalPartsNum := int(math.Ceil(float64(fileSize) / float64(fileChunk)))
		println(totalPartsNum)

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

			if partSize < srv.qouta {
				srv.periods = srv.periods[1:]
				time.Sleep(srv.periods[0].From.Sub(time.Now()))
				srv.qouta = int(131072 * srv.periods[0].Speed * float32(srv.periods[0].To.Sub(time.Now()).Seconds()))
			}
			srv.qouta -= partSize

			partBuffer := make([]byte, partSize)

			file.Read(partBuffer)

			err = srv.sc.Publish("foo", partBuffer)
			//println("!!!!!!!!!!!!!")
			//println(partBuffer)
			if err != nil {
				log.Fatalf("Error during publish: %v\n", err)
			}
		}
		//panic("<-------bruh")
		sqlStatement := `UPDATE report
		SET status=?
		WHERE id=?;`
		_, err = srv.db.Exec(sqlStatement, 2, reportRequest.Report.Id)
	}
}

type Period struct {
	Speed float32
	From  time.Time
	To    time.Time
}

func getDatetimeFromStr(s string) time.Time {
	fromY, _ := strconv.Atoi(s[0:4])
	fromM, _ := strconv.Atoi(s[5:7])
	fromD, _ := strconv.Atoi(s[8:10])
	fromh, _ := strconv.Atoi(s[11:13])
	fromm, _ := strconv.Atoi(s[14:16])
	froms, _ := strconv.Atoi(s[17:19])
	tFrom := time.Date(fromY, time.Month(fromM), fromD, fromh, fromm, froms, 0, time.Now().Location())
	return tFrom
}

func parseDates() []Period {
	type PeriodTmp struct {
		Speed float32 `json:"speed"`
		From  string  `json:"from"`
		To    string  `json:"to"`
	}
	periodsTmp := []PeriodTmp{}
	periods := []Period{}

	file, err := os.Open("periods.json")
	if err != nil {
		log.Fatal(fmt.Errorf("unable to open file for reading: %w", err))
	}
	defer func() { _ = file.Close() }()

	fileText, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to read file: %w", err))
	}

	if err = json.Unmarshal(fileText, &periodsTmp); err != nil {
		log.Fatal(fmt.Errorf("unable to unmarshal periods.json: %w", err))
	}
	if err = file.Close(); err != nil {
		log.Fatal(fmt.Errorf("unable to close the periods file: %w", err))
	}

	for _, periodTmp := range periodsTmp {
		tFrom := getDatetimeFromStr(periodTmp.From)
		tTo := getDatetimeFromStr(periodTmp.To)
		fmt.Println(tFrom)
		fmt.Println(tTo)
		period := Period{
			Speed: periodTmp.Speed,
			From:  tFrom,
			To:    tTo,
		}
		periods = append(periods, period)
	}
	return periods
}

func runMarsServer() {
	periods := parseDates()
	timeNow := time.Now()
	curPeriodInd := -1
	for i, period := range periods {
		if (timeNow.After(period.From) || timeNow.Equal(period.From)) && (timeNow.Before(period.To) || timeNow.Equal(period.From)) {
			curPeriodInd = i
		}
	}
	for _, period := range periods {
		if timeNow.Before(period.From) {
			time.Sleep(time.Until(period.From))
			break
		}
	}

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
		qouta:         int(131072 * periods[curPeriodInd].Speed * float32(periods[curPeriodInd].To.Sub(time.Now()).Seconds())),
		periods:       periods[curPeriodInd:],
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
