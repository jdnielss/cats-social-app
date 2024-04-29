package repository

import (
	"database/sql"
	"time"

	"jdnielss.dev/cats-social-app/model"
)

type EnrollmentRepository interface {
	Create(payload model.Enrollment) (model.Enrollment, error)
	Get(id string) (model.Enrollment, error)
}

type enrollmentRepository struct {
	db *sql.DB
}

func (e *enrollmentRepository) Get(id string) (model.Enrollment, error) {

	var enrollment model.Enrollment
	err := e.db.QueryRow(`
	select
		e.id,
		c.id,
		c.course_full_name,
		c.course_short_name,
		c.description,
		c.course_start_date,
		c.course_end_date,
		c.course_image,
		c.created_at,
		c.updated_at,
		e.status,
		e.created_at,
		e.updated_at
	from
		enrollments e
	join courses c on
		c.id = e.course_id
	where
		e.id = $1;
	`, id).Scan(
		&enrollment.Id,
		&enrollment.Course.Id,
		&enrollment.Course.CourseFullName,
		&enrollment.Course.CourseShortName,
		&enrollment.Course.Description,
		&enrollment.Course.CourseStartDate,
		&enrollment.Course.CourseEndDate,
		&enrollment.Course.CourseImage,
		&enrollment.Course.CreatedAt,
		&enrollment.Course.UpdatedAt,
		&enrollment.Status,
		&enrollment.CreatedAt,
		&enrollment.UpdatedAt,
	)

	if err != nil {
		return model.Enrollment{}, err
	}

	rows, err := e.db.Query(`
	select
		ed.id,
		ed.enrollment_id,
		u.id,
		u.first_name,
		u.last_name,
		u.email,
		u.username,
		u."role" ,
		u.photo ,
		u.created_at ,
		u.updated_at,
		ed.created_at ,
		ed.updated_at
	from
		enrollment_details ed
	join enrollments e on
		e.id = ed.enrollment_id
	join users u on
		u.id = ed.user_id
	where ed.enrollment_id = $1;
	`, enrollment.Id)

	if err != nil {
		return model.Enrollment{}, err
	}

	for rows.Next() {
		var enrollmentDetail model.EnrollmentDetail
		rows.Scan(
			&enrollmentDetail.Id,
			&enrollmentDetail.EnrollmentId,
			&enrollmentDetail.User.Id,
			&enrollmentDetail.User.FirstName,
			&enrollmentDetail.User.LastName,
			&enrollmentDetail.User.Email,
			&enrollmentDetail.User.Username,
			&enrollmentDetail.User.Role,
			&enrollmentDetail.User.Photo,
			&enrollmentDetail.User.CreatedAt,
			&enrollmentDetail.User.UpdatedAt,
			&enrollmentDetail.CreatedAt,
			&enrollmentDetail.UpdatedAt,
		)

		enrollment.EnrollmentDetails = append(enrollment.EnrollmentDetails, enrollmentDetail)
	}

	return enrollment, nil
}

func (e *enrollmentRepository) Create(payload model.Enrollment) (model.Enrollment, error) {
	// Transactional
	tx, err := e.db.Begin()
	if err != nil {
		return model.Enrollment{}, err
	}

	// insert enrollment
	var enrollment model.Enrollment
	var enrollmentDetails []model.EnrollmentDetail
	err = tx.QueryRow(`
	INSERT INTO enrollments (course_id, status, updated_at) VALUES ($1,$2,$3)
	RETURNING id, status, created_at, updated_at`,
		payload.Course.Id,
		"active",
		time.Now(),
	).Scan(
		&enrollment.Id,
		&enrollment.Status,
		&enrollment.CreatedAt,
		&enrollment.UpdatedAt,
	)

	if err != nil {
		return model.Enrollment{}, tx.Rollback()
	}

	for _, v := range payload.EnrollmentDetails {
		var enrollmentDetail model.EnrollmentDetail
		err := tx.QueryRow(`
		INSERT INTO enrollment_details (enrollment_id, user_id, updated_at) VALUES ($1,$2,$3)
		RETURNING id, enrollment_id, created_at, updated_at`,
			enrollment.Id,
			v.User.Id,
			time.Now(),
		).Scan(
			&enrollmentDetail.Id,
			&enrollmentDetail.EnrollmentId,
			&enrollmentDetail.CreatedAt,
			&enrollmentDetail.UpdatedAt,
		)

		if err != nil {
			return model.Enrollment{}, tx.Rollback()
		}
		enrollmentDetail.User = v.User
		enrollmentDetails = append(enrollmentDetails, enrollmentDetail)
	}

	if err := tx.Commit(); err != nil {
		return model.Enrollment{}, err
	}
	enrollment.Course = payload.Course
	enrollment.EnrollmentDetails = enrollmentDetails
	return enrollment, nil
}

func NewEnrollmentRepository(db *sql.DB) EnrollmentRepository {
	return &enrollmentRepository{db: db}
}
