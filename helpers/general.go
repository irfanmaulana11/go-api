package helpers

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type ValidationErrors struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func ToSnakeCase(str string) string {
	var (
		matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
		matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
	)
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ValidationError(s interface{}, err error) *ValidationErrors {
	var (
		typ         = reflect.TypeOf(s).Elem()
		errs        = map[string]string{}
		errsByField = govalidator.ErrorsByField(err)
	)

	for i := 0; i < typ.NumField(); i++ {
		var message string
		key := typ.Field(i).Name
		tag := typ.Field(i).Tag

		if tag.Get("json") == "" {
			message = errsByField[key]
		} else {
			message = errsByField[tag.Get("json")]
		}

		if message != "" {
			name := strings.Split(tag.Get("json"), " ")[0]

			if name == "-" {
				continue
			}

			if name != "" {
				errs[name] = message
				continue
			}

			sKey := ToSnakeCase(key)
			errs[sKey] = message
		}
	}

	mapper := &ValidationErrors{Message: "Validation Error", Errors: errs}
	return mapper
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
