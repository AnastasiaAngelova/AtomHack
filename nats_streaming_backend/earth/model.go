package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

//	type delivery struct {
//		Name    string `json:"name"`
//		Phone   string `json:"phone"`
//		Zip     string `json:"zip"`
//		City    string `json:"city"`
//		Address string `json:"address"`
//		Region  string `json:"region"`
//		Email   string `json:"email"`
//	}
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

type Cache struct {
	Reports map[int]Report
}

var body []byte

var Request RequestFromNats
var cache = Cache{
	Reports: make(map[int]Report),
}

func (req RequestFromNats) createFile() *os.File {
	file, err := os.Create(req.Report.FileName)
	if err != nil {
		fmt.Print(err)
	}
	return file
}

func (req RequestFromNats) writeFile(data []byte) {
	file, err := os.OpenFile(req.Report.FileName, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()
	file.Write(data)
}

func (c Cache) from_json(json_str string) error {
	request := RequestFromNats{}

	if err := json.Unmarshal([]byte(json_str), &request); err != nil {
		// log.Panic()
		// fmt.Errorf()
		// fmt.Errorf("Error: %e", err)
		return err
	} // Десериализация JSON в структуру
	Request = request
	report := request.Report
	if _, ok := cache.Reports[report.Id]; ok {
		return nil

		// if val == nil

		// fmt.Print(val)
	}
	c.Reports[report.Id] = report

	// log.Print(cache.Orders[len(cache.Orders)-1])
	err1 := saveInDB(report.Id, json_str)
	if err1 != nil {
		return err1
	}
	return nil

}

// func (c Cache) from_db(id string, json_str string) {
// 	order := Order{}

// 	if err := json.Unmarshal([]byte(json_str), &order); err != nil {
// 		// log.Panic()
// 		// fmt.Errorf()
// 		panic(err)
// 	} // Десериализация JSON в структуру

// 	c.Orders[id] = order
// 	// log.Print(cache.Orders["b563feb7b2b84b6test"])
// }

func (c Cache) by_id(id int) Report {
	// if c.Orders[id] != nil {

	// }
	return c.Reports[id]
}

func (c Cache) to_json(id int) string {
	bytes, err := json.MarshalIndent(c.by_id(id), "", "\t")
	if err != nil {
		// log.Panic()
		// fmt.Errorf()
		panic(err)
	} // Десериализация JSON в структуру
	// order := ""
	// for i := 0; i < len(order_bytes); i++ {
	// 	order += order_bytes[i]
	// }
	return string(bytes)
}

func (r Report) to_json() string {
	bytes, err := json.Marshal(r)
	if err != nil {
		// log.Panic()
		// fmt.Errorf()
		panic(err)
	} // Десериализация JSON в структуру
	// order := ""
	// for i := 0; i < len(order_bytes); i++ {
	// 	order += order_bytes[i]
	// }
	return string(bytes)
}

func (r Report) Value() (driver.Value, error) {
	return json.Marshal(r)
	// if err != nil {
	// 	// log.Panic()
	// 	// fmt.Errorf()
	// 	panic(err)
	// } // Десериализация JSON в структуру
	// return bytes
}

// func start_program() {

// }

func (o *Report) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &o)
}

func saveInDB(id int, json_str string) error {

	db, err := sql.Open("postgres", "user=tm_admin password=admin dbname=nats_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	order := cache.Reports[id]
	// fmt.Print("\n")
	// fmt.Print(order.to_json())
	sqlStatement := `
	INSERT INTO message VALUES 
	($1, 
	$2, $3)`
	// fmt.Print(cache.Orders[id])
	_, err = db.Exec(sqlStatement, id, order, order.FileName)
	if err != nil {
		return err
	}
	return nil

}
func load_from_db() {
	db, err := sql.Open("postgres", "user=tm_admin password=admin dbname=earth sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from reports")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var data string

		if err := rows.Scan(&id, &data); err != nil {
			log.Fatal(err)
		}

		// fmt.Print(reflect.TypeOf(data))
		// fmt.Print(data)
		// cache.from_db(id, data)

		// cache[id] = data
	}
	// log.Print(cache)
}