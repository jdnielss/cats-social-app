package repository

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"

	"jdnielss.dev/cats-social-app/model"
)

type CatRepository interface {
	// Get(race, sex *string) ([]model.Cat, error)
	Get(q []string) ([]model.Cat, error)
}

type catRepository struct {
	db *sql.DB
}

func NewCatRepository(db *sql.DB) CatRepository {
	return &catRepository{db: db}
}

func (c *catRepository) Get(q []string) ([]model.Cat, error) {
	sqlQuery := "SELECT * FROM cats"
	var args []any

	for _, queryString := range q {
		var conditions []string
		sqlQuery += " WHERE "
		splittedQuery := strings.Split(queryString, "&")

		for idx, pair := range splittedQuery {
			splittedPair := strings.Split(pair, "=")
			key := splittedPair[0]
			value := splittedPair[1]

			if key == "ageInMonth" {
				decoded, err := url.QueryUnescape(value)

				if err != nil {
					continue
				}

				if decoded == "<4" {
					conditions = append(conditions, fmt.Sprintf("\"%v\" < 4", key))
				} else if decoded == ">4" {
					conditions = append(conditions, fmt.Sprintf("\"%v\" > 4", key))
				} else if decoded == "4" {
					conditions = append(conditions, fmt.Sprintf("\"%v\" = 4", key))
				}
				continue
			}

			args = append(args, value)
			conditions = append(conditions, fmt.Sprintf("\"%v\" = $%d", key, idx+1))
		}
		sqlQuery += strings.Join(conditions, " AND ")
	}

	sqlQuery += " ORDER BY \"createdAt\" DESC"

	fmt.Println(sqlQuery)

	rows, err := c.db.Query(sqlQuery, args...)
	if err != nil {
		fmt.Printf(`Error Repo %s`, err)
		return nil, err
	}
	defer rows.Close()

	var cats []model.Cat
	for rows.Next() {
		var cat model.Cat
		var imageUrlsJSON string // Untuk menyimpan data JSON dalam bentuk string dari kolom imageUrl

		err := rows.Scan(
			&cat.ID,
			&cat.Name,
			&cat.Race,
			&cat.Sex,
			&cat.AgeInMonth,
			&cat.Description,
			&imageUrlsJSON, // Pindai nilai JSON dari kolom imageUrl ke dalam string
			&cat.HasMatched,
			&cat.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Memisahkan string JSON menjadi slice string berdasarkan tanda koma
		imageUrls := strings.Split(imageUrlsJSON, ",")
		// Membersihkan setiap URL dari spasi kosong di awal dan akhir
		for i := range imageUrls {
			imageUrls[i] = strings.TrimSpace(imageUrls[i])
		}
		// Set slice string hasil pemisahan sebagai nilai untuk atribut ImageUrls dari struktur Cat
		cat.ImageUrls = imageUrls

		cats = append(cats, cat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cats, nil
}
