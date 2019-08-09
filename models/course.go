package models

import (
	"fmt"
	"github.com/google/uuid"
)

type Course struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Par  int    `json:"par"`
}

func GetCourses() ([]Course, error) {

	results, err := getDatabase().Query("SELECT * FROM courses")
	if err != nil {
		return []Course{}, err
	}

	var courses []Course

	for results.Next() {
		var course Course

		err := results.Scan(&course.Id, &course.Name, &course.Par)
		if err != nil {
			fmt.Println(err)

			return []Course{}, err
		}

		courses = append(courses, course)

	}

	return courses, nil

}

func NewCourse(course Course) error {

	uid, _ := uuid.NewUUID()

	_, err := getDatabase().Query("INSERT INTO courses VALUES (?, ?, ?)", uid.String(), course.Name, course.Par)
	if err != nil {

		fmt.Println(err)

		return err
	}

	return nil

}

func UpdateCourse(course Course) error {

	_, err := getDatabase().Query("UPDATE courses SET name = ?, par = ? WHERE id = ?", course.Name, course.Par, course.Id)
	if err != nil {
		return err
	}

	return nil

}

func DeleteCourse(id string) error {

	_, err := getDatabase().Query("DELETE FROM courses WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil

}
