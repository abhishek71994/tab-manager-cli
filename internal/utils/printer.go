package utils

import (
	"fmt"
	"reflect"
)

// PrintListStyle outputs data in a readable, labeled format without truncating any fields.
func PrintListStyle(data interface{}) error {
	// Validate input is a slice
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		return fmt.Errorf("expected a slice, but got %s", val.Kind())
	}

	// Check for empty data
	if val.Len() == 0 {
		return fmt.Errorf("no data to display")
	}

	// Get the type of the slice's element
	elemType := val.Index(0).Type()

	// Loop through each element
	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)

		fmt.Printf("Record %d:\n", i+1)
		for j := 0; j < elem.NumField(); j++ {
			fieldName := elemType.Field(j).Name
			fieldValue := fmt.Sprintf("%v", elem.Field(j).Interface())

			fmt.Printf("  %s: %s\n", fieldName, fieldValue)
		}
		fmt.Println() // Add a blank line between records
	}

	return nil
}
