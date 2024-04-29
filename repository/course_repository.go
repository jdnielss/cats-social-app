package repository

import (
	"database/sql"

	"jdnielss.dev/cats-social-app/model"
)

type CourseRepository interface {
	Get(id string) (model.Course, error)
}

type courseRepository struct {
	db *sql.DB
}

func (c *courseRepository) Get(id string) (model.Course, error) {
	var course model.Course
	err := c.db.QueryRow(`SELECT * FROM courses WHERE id = $1`, id).
		Scan(
			&course.Id,
			&course.CourseFullName,
			&course.CourseShortName,
			&course.Description,
			&course.CourseStartDate,
			&course.CourseEndDate,
			&course.CourseImage,
			&course.CreatedAt,
			&course.UpdatedAt,
		)

	if err != nil {
		return model.Course{}, err
	}

	return course, nil
}

func NewCourseRepository(db *sql.DB) CourseRepository {
	return &courseRepository{db: db}
}
