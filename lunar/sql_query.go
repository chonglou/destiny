package lunar

import "database/sql"

//PgSQLQuery query by gorm
type PgSQLQuery struct {
	Db *sql.DB
}

//Clear clear table items
func (p *PgSQLQuery) Clear() error {
	_, err := p.Db.Exec("DELETE FROM lunar_items")
	return err
}

//Prepare create table and indexes
func (p *PgSQLQuery) Prepare() error {
	for _, s := range []string{
		`CREATE TABLE IF NOT EXISTS lunar_items(
			id SERIAL PRIMARY KEY,

			l_year SMALLINT NOT NULL,
			l_month SMALLINT NOT NULL,
			l_day SMALLINT NOT NULL,
			s_year SMALLINT NOT NULL,
			s_month SMALLINT NOT NULL,
			s_day SMALLINT NOT NULL,

			leap BOOL NOT NULL,
			term varchar(12)
		);`,
		"CREATE UNIQUE INDEX IF NOT EXISTS idx_lunar_items_s ON lunar_items(s_year, s_month, s_day);",
		"CREATE UNIQUE INDEX IF NOT EXISTS idx_lunar_items_l ON lunar_items(l_year, l_month, l_day, leap);",
		"CREATE INDEX IF NOT EXISTS idx_lunar_items_term ON lunar_items(term);",
	} {
		if _, err := p.Db.Exec(s); err != nil {
			return err
		}
	}
	return nil
}

//Write create a new record
func (p *PgSQLQuery) Write(item *Date) error {
	_, err := p.Db.Exec(
		"INSERT INTO lunar_items(s_year, s_month, s_day, l_year, l_month, l_day, leap, term) VALUES($1, $2, $3, $4, $5, $6, $7, $8)",
		item.SYear,
		item.SMonth,
		item.SDay,
		item.LYear,
		item.LMonth,
		item.LDay,
		item.Leap,
		item.Term,
	)
	return err
}

//FromSolar 从公历查询
func (p *PgSQLQuery) FromSolar(y, m, d int) (Date, error) {
	var item Date
	err := p.Db.
		QueryRow(
			"SELECT id, s_year, s_month, s_day, l_year, l_month, l_day, leap, term FROM lunar_items WHERE s_year = $1 AND s_month = $2 AND s_day = $3 LIMIT 1",
			y, m, d,
		).
		Scan(
			&item.ID,
			&item.SYear,
			&item.SMonth,
			&item.SDay,
			&item.LYear,
			&item.LMonth,
			&item.LDay,
			&item.Leap,
			&item.Term,
		)

	return item, err
}

//FromLunar 从农历查询
func (p *PgSQLQuery) FromLunar(y, m, d int) ([]Date, error) {
	var items []Date
	rows, err := p.Db.Query(
		"SELECT id, s_year, s_month, s_day, l_year, l_month, l_day, leap, term FROM lunar_items WHERE l_year = $1 AND l_month = $2 AND l_day = $3 ORDER BY ID ASC",
		y, m, d,
	)
	defer rows.Close()
	for rows.Next() {
		var item Date
		if err = rows.Scan(
			&item.ID,
			&item.SYear,
			&item.SMonth,
			&item.SDay,
			&item.LYear,
			&item.LMonth,
			&item.LDay,
			&item.Leap,
			&item.Term,
		); err != nil {
			break
		}
		items = append(items, item)
	}

	return items, err
}
