package main

func main() {
	load_from_db()
	done := make(chan bool)
	go runSubscriber(done)
	go runHttpServer(done)
	go CLReader(done)
	<-done
	<-done

}
