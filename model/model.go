package model

import (
	"fmt"
)

// DataBaseType : model type
type DataBaseType string

// MeetAndGreet : method on DataBaseType
func (d *DataBaseType) MeetAndGreet() string {

	return fmt.Sprintf("Hi i am a %s database", string(*d))
}
