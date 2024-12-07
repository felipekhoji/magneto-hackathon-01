package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Conecta ao banco de dados SQLite e inicializa a tabela de taxas de câmbio, se necessário.
func InitDB(dbLocation string) error {
	var err error
	db, err = sql.Open("sqlite3", dbLocation)
	if err != nil {
		return err
	}

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS exchange_rates (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        from_currency TEXT NOT NULL,
        to_currency TEXT NOT NULL,
        rate REAL NOT NULL,
        creation_date DATETIME DEFAULT CURRENT_TIMESTAMP
    );`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	fmt.Println("Banco de dados SQLite conectado e tabela de taxas de câmbio inicializada.")
	return nil
}

// Função para adicionar uma nova taxa de câmbio
func AddExchangeRate(fromCurrency, toCurrency string, rate float64) error {
	insertQuery := `INSERT INTO exchange_rates (from_currency, to_currency, rate) VALUES (?, ?, ?)`
	_, err := db.Exec(insertQuery, fromCurrency, toCurrency, rate)
	if err != nil {
		return err
	}
	fmt.Printf("Taxa de câmbio adicionada: %s -> %s = %.2f\n", fromCurrency, toCurrency, rate)
	return nil
}

// Função para consultar a taxa de câmbio entre duas moedas
func GetExchangeRate(fromCurrency, toCurrency string) (float64, error) {
	var rate float64
	query := `SELECT rate FROM exchange_rates WHERE from_currency = ? AND to_currency = ? ORDER BY creation_date DESC LIMIT 1`
	err := db.QueryRow(query, fromCurrency, toCurrency).Scan(&rate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("taxa de câmbio não encontrada")
		}
		return 0, err
	}
	return rate, nil
}
