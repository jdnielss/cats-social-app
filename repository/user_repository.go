package repository

import (
	"database/sql"

	"jdnielss.dev/cats-social-app/model"
)

// TODO
/*
1. Kita harus siapkan sebuah kontrak (interface)
2. Interface ini yang akan di lempar lempar ke service lainya (injection)
3. Juga berguna untuk kemudahan Unit Testing
4. Biasanya nama interface seperti nama file dan dia ter-Expose(public)
5. Setelah itu kita buatkan sebuah struct untuk di kirim sebagai receiver method
6. Method-method inilah sebagai isian kontrak dari interface
7. Terakhir kita buatkan sebuah function sebagai perantara untuk memanggil interface
   agar method-method yang dibuat bisa di panggil keluar.
8. Function ini biasanya disebut sebagai constructor, diawali dengan kata New...
*/

type UserRepository interface {
	Get(id string) (model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Get(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`
		SELECT 
			id, first_name, last_name, email, username, role, photo, created_at, updated_at 
		FROM 
			users 
		WHERE 
			id = $1`, id).
		Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Username,
			&user.Role,
			&user.Photo,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
