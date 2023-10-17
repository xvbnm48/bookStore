package mysql

func (r *dbReadWriter) GetAllBook() ([]string, error) {
	rows, err := r.db.Query(`SELECT * FROM library `)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []string
	for rows.Next() {
		var book string
		if err := rows.Scan(&book); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
