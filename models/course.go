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

type courseWithResults struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Par     int      `json:"par"`
	Results []result `json:"results"`
}

type result struct {
	Id   string `json:"id"`
	Date string `json:"date"`
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

func GetCourse(id string) (courseWithResults, error) {

	var c courseWithResults
	var rs []result

	// Get Course
	err := getDatabase().QueryRow("SELECT courses.id, courses.name, courses.par FROM courses WHERE courses.id=?;", id).Scan(&c.Id, &c.Name, &c.Par)
	if err != nil {
		return courseWithResults{}, err
	}

	// Get Results Associated With it
	results, err := getDatabase().Query("SELECT results.id, results.date FROM results WHERE results.course=?;", id)
	if err != nil {
		return courseWithResults{}, err
	}

	for results.Next() {
		var r result

		err = results.Scan(&r.Id, &r.Date)
		if err != nil {
			return courseWithResults{}, err
		}

		rs = append(rs, r)
	}

	c.Results = rs

	return c, nil

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
