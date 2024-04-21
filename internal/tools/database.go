package tools

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/4shb0rne/goapi-basic/internal/models"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coins    int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func SetupDatabase() (*sql.DB, error) {
	// MySQL connection parameters
	username := "root"
	password := ""
	dbname := "go_learn"

	// Connect to the MySQL database
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", username, password, dbname))
	if err != nil {
		return nil, err
	}

	// Attempt to ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Info("Connected to the MySQL database")
	return db, nil
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}

type DatabaseHandler interface {
	GetDatabase() (*sql.DB, error)
	Insert(data interface{}) error
	ViewAll() (interface{}, error)
	Update(id int, data interface{}) error
	Delete(id int) error
}

// MySQLHandler implements DatabaseHandler for MySQL database
type MySQLHandler struct {
	Username string
	Password string
	DBName   string
}

func NewMySQLHandler(username, password, dbname string) *MySQLHandler {
	return &MySQLHandler{
		Username: username,
		Password: password,
		DBName:   dbname,
	}
}

func (mh *MySQLHandler) GetDatabase() (*sql.DB, error) {
	// Connect to the MySQL database
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", mh.Username, mh.Password, mh.DBName))
	if err != nil {
		return nil, err
	}

	// Attempt to ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to the MySQL database")
	return db, nil
}

func (mh *MySQLHandler) ViewAll(db *sql.DB) (interface{}, error) {
	// Prepare the SQL query
	query := "SELECT * FROM Product"

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store the products
	var products []models.Product

	// Iterate over the rows
	for rows.Next() {
		// Create a new Product struct to store the data for each row
		var product models.Product

		// Scan the row data into the Product struct fields
		err := rows.Scan(&product.ProductID, &product.ProductName, &product.ProductPrice, &product.ProductStock)
		if err != nil {
			return nil, err
		}

		// Append the Product struct to the products slice
		products = append(products, product)
	}

	// Check for any errors that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// Create inserts data into the database
func (mh *MySQLHandler) Insert(data interface{}) error {
	// Type assertion to convert data to []models.Product
	products, ok := data.([]models.Product)
	if !ok {
		return errors.New("data is not a []models.Product")
	}

	// Connect to the MySQL database
	db, err := mh.GetDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	// Prepare the SQL statement for inserting products
	query := "INSERT INTO Product (productName, productPrice, productStock) VALUES (?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Iterate over the products and execute the INSERT statement for each
	for _, product := range products {
		_, err := stmt.Exec(product.ProductName, product.ProductPrice, product.ProductStock)
		if err != nil {
			return err
		}
	}

	return nil
}

// Read retrieves data from the database based on the provided ID

// Update updates data in the database based on the provided ID
func (mh *MySQLHandler) Update(id int, data interface{}) error {

	return nil
}

// Delete deletes data from the database based on the provided ID
func (mh *MySQLHandler) Delete(id int) error {

	return nil
}
