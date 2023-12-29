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

// defining table Booking as a structure
type Booking struct {
	Custid         int     `json:"custid"`
	Firstname      string  `json:"firstname"`
	Lastname       string  `json:"lastname"`
	Tourid         int     `json:"tourid"`
	Tourname       string  `json:"tourname"`
	Tourdate       string  `json:"tourdate"`
	Payment_amount float32 `json:"payment_amount"`
}

// Function to connect to the database
func connection() (*sql.DB, error) {
	//converting the connections parameters to string and saving the values
	//using sprinf
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//to connect to database using sql.open
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Error in db", err)
		return nil, err
	}
	//defer db.Close()

	//checking the connection using ping command
	err = db.Ping()
	if err != nil {
		log.Println("Error in ping", err)
		return nil, err
	}
	fmt.Println("Sucessfully connected")
	return db, err

}

// Query function : wil return output in the form of booking structure
func Query(Querytype int) ([]Booking, error) {
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
		return nil, fmt.Errorf("invalid query type")
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error in rows", err)
		return nil, err
	}
	defer rows.Close()

	//looping through data
	var data Booking
	var Output []Booking
	var count int
	for rows.Next() {

		err := rows.Scan(&data.Custid, &data.Firstname, &data.Lastname, &data.Tourid, &data.Tourname, &data.Tourdate, &data.Payment_amount)
		if err != nil {
			log.Println("Error in err", err)
			return nil, err
		}
		//fmt.Printf("custid:%d,firstname:%s,lastname:%s,tourid:%d,tourname:%s,tourdate:%s,payment_amount:%f\n", data.Custid, data.Firstname, data.Lastname, data.Tourid, data.Tourname, data.Tourdate, data.Payment_amount)
		count++
		Output = append(Output, data)
	}
	fmt.Println("Number of entries that matched the query:", count)

	//defer db.Close()
	return Output, err
}
