package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// Product model which holds data of each endividual model
type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
	Created     time.Time
}

// Model which wraps the sql.DB connection pool
type ProductModel struct {
	DB *sql.DB
}

// Insert a new product into the database
func (p *ProductModel) Insert(name string, description string, price int) (int, error) {
	// The actual command which is to be executed in the database
	cmd := `insert into products (name, description, price, created) 
        	values(?, ?, ?, datetime());`

	// Execution of the command
	result, err := p.DB.Exec(cmd, name, description, price)
	if err != nil {
		return 0, err
	}
	// Get the ID of the inserted data
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), err
}

// Return a specific product based on its id
func (p *ProductModel) Get(id int) (*Product, error) {
	// Command to get the specific data
	cmd := `select * from products where id = ?;`
	row := p.DB.QueryRow(cmd, id)

	d := &Product{}
	nid, nname, ndescription, nprice, ncreated := 0, "", "", 0, ""
	err := row.Scan(&nid, &nname, &ndescription, &nprice, &ncreated)
	ntime, err := time.Parse("2006-01-02 15:04:05", ncreated)
	if err != nil {
		return nil, err
	}
	// err := row.Scan(&d.ID, &d.Name, &d.Description, &d.Price, &d.Created)

	d.ID = nid
	d.Name = nname
	d.Description = ndescription
	d.Price = nprice
	d.Created = ntime

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return d, nil
}

// Return 10 most recently created snippets
func (p *ProductModel) GetList() ([]*Product, error) {
	const cmd string = `select id, name, price, created from products
											order by id;`
	rows, err := p.DB.Query(cmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*Product{}
	for rows.Next() {

		p := &Product{}
		nid, nname, nprice, ncreated := 0, "", 0, ""
		err := rows.Scan(&nid, &nname, &nprice, &ncreated)
		if err != nil {
			return nil, err
		}
		// err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Created)
		ntime, err := time.Parse("2006-01-02 15:04:05", ncreated)
		if err != nil {
			return nil, err
		}

		p.ID = nid
		p.Name = nname
		p.Price = nprice
		p.Created = ntime

		result = append(result, p)
	}
	// Catch errors encountered during the iteration over the table
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ProductModel) SearchResults(value string, lowerPrice, upperPrice int) ([]*Product, error) {

	var cmd string = `select id, name, price, created from products where price >= ? and price <= ?`
	if value != "" {
		cmd += ` and name like "%` + value + `%"`
	}

	results, err := p.DB.Query(cmd, lowerPrice, upperPrice)
	if err != nil {
		return nil, err
	}

	ret := []*Product{}
	for results.Next() {
		p := &Product{}
		results.Scan(&p.ID, &p.Name, &p.Price, &p.Created)
		ret = append(ret, p)
	}

	if err = results.Err(); err != nil {
		return nil, err
	}

	fmt.Println(cmd)
	return ret, nil

}

// func (p *ProductModel) createProduct(name, description string, price int) error {
// 	const cmd string = `insert into products (name, description, price, created)
// 		values(?, ?, ?, now());`
//
// 	_, err := p.DB.Exec(cmd, name, description, price)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
