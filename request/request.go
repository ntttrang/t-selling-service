package request

import (
	"github.com/ntttrang/t-selling-service/messagehandler"
	validator "gopkg.in/go-playground/validator.v9"
)

//ValidateFields validate request
func ValidateFields(model interface{}) *messagehandler.Mess {
	var err *messagehandler.Mess
	validate := validator.New()

	e := validate.Struct(model)
	if nil != e {
		if errs, ok := e.(validator.ValidationErrors); ok {
			for _, er := range errs {
				switch er.Tag() {
				case "required":
					err = messagehandler.W000007.GetMessage(er.StructField())
					return err
				}
				// TODO: Define more case
			}
		}
	}
	return err
}
