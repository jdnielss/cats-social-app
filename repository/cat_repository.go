package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"jdnielss.dev/cats-social-app/model"
	"jdnielss.dev/cats-social-app/model/dto"
)

type CatRepository interface {
	// Get(race, sex *string) ([]model.Cat, error)
	Get(q []string) ([]model.Cat, error)
	Create(payload dto.CatRequestDTO) (dto.CreateCatResponseDTO, error)
}

type catRepository struct {
	db *sql.DB
}

func NewCatRepository(db *sql.DB) CatRepository {
	return &catRepository{db: db}
}

func (c *catRepository) Get(q []string) ([]model.Cat, error) {
	// fmt.Printf(`Q %v`, q)
	sqlQuery := "SELECT * FROM cats WHERE "
	var args []interface{}
	for key, values := range q {
		fmt.Printf(`Q %v`, values)

		if len(values) == 0 {
			continue
		}

		var conditions []string

		for _, value := range values {
			conditions = append(conditions, fmt.Sprintf("%s = ?", key))
			args = append(args, value)
		}

		sqlQuery += strings.Join(conditions, " AND")
	}

	// fmt.Printf(`Query %v`, sqlQuery)

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

func (c *catRepository) Create(payload dto.CatRequestDTO) (dto.CreateCatResponseDTO, error) {
	// Start a transaction
	tx, err := c.db.Begin()
	if err != nil {
		return dto.CreateCatResponseDTO{}, err
	}

	// Perform database operation within the transaction
	var cat model.Cat
	var imagesUrl string

	err = tx.QueryRow(`
		INSERT INTO cats (name, race, sex, ageinmonth, description, imageurls, hasmatched, createdat, userid) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, name, race, sex, ageinmonth, description, imageurls, hasmatched, createdat
	`, payload.Name, payload.Race, payload.Sex, payload.AgeInMonth, payload.Description, strings.Join(payload.ImageUrls, ","), false, time.Now(), 1).Scan(
		&cat.ID,
		&cat.Name,
		&cat.Race,
		&cat.Sex,
		&cat.AgeInMonth,
		&cat.Description,
		&imagesUrl,
		&cat.HasMatched,
		&cat.CreatedAt,
		&cat.UserID,
	)

	// Split the imagesUrl string into a slice of strings
	cat.ImageUrls = strings.Split(imagesUrl, ",")

	// Commit the transaction if there are no errors, otherwise rollback
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			// Rollback failed, return both errors
			return dto.CreateCatResponseDTO{}, fmt.Errorf("transaction rollback error: %v, query error: %v", rollbackErr, err)
		}
		// Return only the query error
		return dto.CreateCatResponseDTO{}, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return dto.CreateCatResponseDTO{}, err
	}

	res := dto.CreateCatResponseDTO{
		ID:        cat.ID,
		CreatedAt: cat.CreatedAt,
	}
	// Return the created cat
	return res, nil
}
