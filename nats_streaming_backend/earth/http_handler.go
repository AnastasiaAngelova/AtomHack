package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func showOrder(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Print(cache)
	if _, ok := cache.Reports[id]; ok {
		fmt.Fprint(w, cache.to_json(id))

		// if val == nil

		// fmt.Print(val)
	} else {
		err := fmt.Errorf("Error: try to get order by ID = %s", id)
		fmt.Errorf("\n%e\n", err)
		fmt.Fprintf(w, "Error: try to get order by ID = %s(there is no such order)", id)
	}

}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return

	}

}

// func chanHandler(w http.ResponseWriter, r *http.Request) {
// 	conn, err := stan.Connect(
// 		"test-cluster",
// 		"test-client",
// 		stan.NatsURL("nats://localhost:4222"),
// 	)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// wg := &sync.WaitGroup{}

// 	sub, err := conn.Subscribe("counter", func(msg *stan.Msg) {
// 		// Print the value and whether it was redelivered.
// 		fmt.Printf("seq = %d [redelivered = %v]\n", msg.Sequence, msg.Redelivered)

// 		// Add jitter.
// 		// time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
// 		log.Print(msg)
// 		// Mark it is done.
// 		// wg.Done()
// 	})
// 	if err != nil {
// 		fmt.Errorf(err.Error())
// 		log.Panic(sub)
// 	}

// }

func runHttpServer(done chan bool) {
	mux := http.NewServeMux()

	// mux.HandleFunc("", chanHandler)
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/report", showOrder)
	signalChan := make(chan os.Signal, 1)
	go func() {
		log.Println("Server run on: http:localhost:4000")
		err := http.ListenAndServe(":4001", mux)
		log.Fatal(err)
		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			// Do not unsubscribe a durable on exit, except if asked to.
			// if durable == "" || unsubscribe {
			// 	sub.Unsubscribe()
			// }
			// sc.Close()

			done <- true
		}

	}()
	done <- true
	// <-done

	// done <- true
}
