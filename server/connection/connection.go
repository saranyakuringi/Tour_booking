package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// connection parameters for postgres
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Saranya@426"
	dbname   = "postgres"
)

var db *sql.DB

type Booking struct {
	Custid         int     `json:"custid"`
	Firstname      string  `json:"firstname"`
	Lastname       string  `json:"lastname"`
	Tourid         int     `json:"tourid"`
	Tourname       string  `json:"tourname"`
	Tourdate       string  `json:"tourdate"`
	Payment_amount float32 `json:"payment_amount"`
}

func connection() {
	//converting the connections parameters to string and saving the values
	//using sprinf
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//to connect to database using sql.open
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Error in db", err)
		return
	}
	//defer db.Close()

	//checking the connection using ping command
	err = db.Ping()
	if err != nil {
		log.Println("Error in ping", err)
		return
	}
	fmt.Println("Sucessfully connected")

	//write the query

}

func Query(Querytype int) {
	connection()
	var query string
	switch Querytype {
	case 1:
		query = "SELECT * FROM BOOKING where tourid=1"
	case 2:
		query = "SELECT * FROM BOOKING where tourid=2"
	case 3:
		query = "SELECT * FROM BOOKING where tourid=3"
	default:
		log.Println("Error in query type")
		return
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error in rows", err)
		return
	}
	defer rows.Close()

	//looping through data
	var data Booking
	var count int
	for rows.Next() {

		err := rows.Scan(&data.Custid, &data.Firstname, &data.Lastname, &data.Tourid, &data.Tourname, &data.Tourdate, &data.Payment_amount)
		if err != nil {
			log.Println("Error in err", err)
			return
		}
		fmt.Printf("custid:%d,firstname:%s,lastname:%s,tourid:%d,tourname:%s,tourdate:%s,payment_amount:%f\n", data.Custid, data.Firstname, data.Lastname, data.Tourid, data.Tourname, data.Tourdate, data.Payment_amount)
		count++
	}
	fmt.Println("Number of entries that matched the query:", count)
	err = rows.Err()
	if err != nil {
		log.Println("Error in err", err)
		return
	}
	defer db.Close()
}
