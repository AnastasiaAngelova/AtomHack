package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

type delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type payment struct {
	Transaction   string `json:"transaction"`
	Request_id    string `json:"request_id"`
	Currency      string `json:"currency"`
	Provider      string `json:"provider"`
	Amount        int    `json:"amount"`
	Payment_dt    int    `json:"payment_dt"`
	Bank          string `json:"bank"`
	Delivery_cost int    `json:"delivery_cost"`
	Goods_total   int    `json:"goods_total"`
	Custom_fee    int    `json:"custom_fee"`
}

type item struct {
	Chrt_id      int    `json:"chrt_id"`
	Track_number string `json:"track_number"`
	Price        int    `json:"price"`
	Rid          string `json:"rid"`
	Name         string `json:"name"`
	Sale         int    `json:"sale"`
	Size         string `json:"size"`
	Total_price  int    `json:"total_price"`
	Nm_id        int    `json:"nm_id"`
	Brand        string `json:"brand"`
	Status       int    `json:"status"`
}

type Order struct {
	Order_uid    string   `json:"order_uid"`
	Track_number string   `json:"track_number"`
	Entry        string   `json:"entry"`
	Delivery     delivery `json:"delivery"`
	Payment      payment  `json:"payment"`
	Items        []item   `json:"items"`
}

type Cache struct {
	Orders map[string]Order
}

// var body []byte
var cache = Cache{
	Orders: make(map[string]Order),
}

func (c Cache) from_json(json_str string) error {
	order := Order{}

	if err := json.Unmarshal([]byte(json_str), &order); err != nil {
		// log.Panic()
		// fmt.Errorf()
		// fmt.Errorf("Error: %e", err)
		return err
	} // Десериализация JSON в структуру
	if _, ok := cache.Orders[order.Order_uid]; ok {
		return nil

		// if val == nil

		// fmt.Print(val)
	}
	c.Orders[order.Order_uid] = order

	// log.Print(cache.Orders[len(cache.Orders)-1])
	err1 := saveInDB(order.Order_uid, json_str)
	if err1 != nil {
		return err1
	}
	return nil

}

func (c Cache) from_db(id string, json_str string) {
	order := Order{}

	if err := json.Unmarshal([]byte(json_str), &order); err != nil {
		// log.Panic()
		// fmt.Errorf()
		panic(err)
	} // Десериализация JSON в структуру

	c.Orders[id] = order
	// log.Print(cache.Orders["b563feb7b2b84b6test"])
}

func (c Cache) by_id(id string) Order {
	// if c.Orders[id] != nil {

	// }
	return c.Orders[id]
}

func (c Cache) to_json(id string) string {
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

func (o Order) to_json() string {
	bytes, err := json.Marshal(o)
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

func (o Order) Value() (driver.Value, error) {
	return json.Marshal(o)
	// if err != nil {
	// 	// log.Panic()
	// 	// fmt.Errorf()
	// 	panic(err)
	// } // Десериализация JSON в структуру
	// return bytes
}

// func start_program() {

// }

func (o *Order) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &o)
}

func saveInDB(id string, json_str string) error {

	db, err := sql.Open("postgres", "user=tm_admin password=admin dbname=nats_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	order := cache.Orders[id]
	// fmt.Print("\n")
	// fmt.Print(order.to_json())
	sqlStatement := `
	INSERT INTO message VALUES 
	($1, 
	$2)`
	// fmt.Print(cache.Orders[id])
	_, err = db.Exec(sqlStatement, id, order)
	if err != nil {
		return err
	}
	return nil

}
func load_from_db() {
	db, err := sql.Open("postgres", "user=tm_admin password=admin dbname=nats_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from message")
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
		cache.from_db(id, data)

		// cache[id] = data
	}
	// log.Print(cache)
}
