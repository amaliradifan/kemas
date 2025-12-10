package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const dbuser = "root:12345678"
const dbname = "kemas"

func ListProductHandler(search, limit, offset string) ([]Product, error) {
	db, err := sql.Open("mysql", dbuser+"@tcp(127.0.0.1:33061)/"+dbname)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT id, name, stock FROM products WHERE name LIKE ? LIMIT ? OFFSET ?"
	rows, err := db.Query(query, "%"+search+"%", limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Stock); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func TransferStockHandler(sourceID, destinationID string, quantity int) error {
	db, err := sql.Open("mysql", dbuser+"@tcp(127.0.0.1:33061)/"+dbname)
	if err != nil {
		return err
	}
	defer db.Close()

  tx, err := db.Begin()
  if err != nil {
    return err
  }
  defer tx.Rollback()

  checkSource := "SELECT stock FROM products WHERE id = ?"
  var sourceStock int
  err = tx.QueryRow(checkSource, sourceID).Scan(&sourceStock)
  if err != nil {
    return err
  }

  if sourceStock < quantity {
    tx.Rollback()
    return sql.ErrNoRows
  }

  _, err = tx.Exec("UPDATE products SET stock = stock - ? WHERE id = ?", quantity, sourceID)
  if err != nil {
    return err
  }

  _, err = tx.Exec("UPDATE destinations SET stock = stock + ? WHERE id = ?", quantity, destinationID)
  if err != nil {
    return err
  } 

  return tx.Commit()
}
