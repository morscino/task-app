package helpers

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"unicode"

	"strings"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
	validator "gopkg.in/go-playground/validator.v9"
)

const (
	alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

type Map map[string]interface{}

// Enum enum interface for validation
type Enum interface {
	IsValid() bool
}

func ValidateInput(input interface{}) []string {
	var errors []string
	v := validator.New()
	v.RegisterValidation("is_enum", ValidateEnum)
	v.RegisterValidation("is_uuid", ValidateUUID)
	v.RegisterValidation("is_password", ValidatePassword)

	err := v.Struct(input)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			switch e.ActualTag() {
			case "required":
				errors = append(errors, fmt.Sprintf("%s field is required", e.Field()))
			case "email":
				errors = append(errors, fmt.Sprintf("%s must be a valid email", e.Field()))
			case "url":
				errors = append(errors, fmt.Sprintf("%s must be a valid url", e.Field()))
			case "gt":
				errors = append(errors, fmt.Sprintf("%s array cannot be empty", e.Field()))
			case "is_enum":
				errors = append(errors, fmt.Sprintf("%s is not a valid %v", e.Value(), e.Type()))
			case "is_uuid":
				errors = append(errors, fmt.Sprintf("%s is not a valid uuid", e.Value()))
			case "is_password":
				errors = append(errors, fmt.Sprintf("%s is not a valid password", e.Value()))
			case "min":
				errors = append(errors, fmt.Sprintf("%s must be at least %s letters", e.Field(), e.Param()))
			case "max":
				errors = append(errors, fmt.Sprintf("%s cannot be more than %s letters", e.Field(), e.Param()))
			case "len":
				errors = append(errors, fmt.Sprintf("%s must be %s in length", e.Field(), e.Param()))
			default:
				errors = append(errors, "an error occurred")
			}
		}
	}

	return errors
}

// ValidateEnum validates if a value is a member of an enum
func ValidateEnum(field validator.FieldLevel) bool {
	value := field.Field().Interface().(Enum)
	return value.IsValid()
}

// ValidateUUID validates if a value is a valid uuid
func ValidateUUID(field validator.FieldLevel) bool {
	value := field.Field().Interface().(string)
	_, err := uuid.Parse(value)
	return err == nil
}

// ValidateUUID validates if a value is a secured password
func ValidatePassword(field validator.FieldLevel) bool {
	var number, upper, special, lenghtOrMore bool
	password := field.Field().Interface().(string)
	letters := 0
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			letters++
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			letters++
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
		}
	}
	lenghtOrMore = letters >= 8
	if !number || !upper || !special || !lenghtOrMore {
		return false
	}
	return true
}

// ToSlug converts a string to a lower case slug
func ToSlug(value string) string {
	var slug string

	splitted := strings.Split(strings.ToLower(value), " ")
	for i := 0; i < len(splitted); i++ {
		slug += splitted[i]
		if i < len(splitted)-1 {
			slug += "-"
		}
	}
	return slug
}

func generateRandom(length, randomRange int) string {
	var random string
	for i := 0; i < length; i++ {
		random += fmt.Sprintf("%s", string(alphabet[rand.Intn(randomRange)]))
	}
	return random
}

func Hash(s string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 14)
	fmt.Println(err)

	return string(bytes)
}

// HashString hashes a string
func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// CompareHashString compares a clear string with it's hash value
func CompareHashString(s, hash string) bool {
	return HashString(s) == hash
}

// CompareHash compares a clear string with it's hash value
func CompareHash(hashed, s string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(s))
	return err == nil
}

// Getenv get envitoment variables
func Getenv(variable string, defaultValue ...string) string {
	env := os.Getenv(variable)
	if env == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}
	return env
}

// StructToMap converts a struct to a map
func StructToMap(s interface{}) map[string]interface{} {
	var mapInterface map[string]interface{}

	marshaled, _ := json.Marshal(s)
	json.Unmarshal(marshaled, &mapInterface)

	return mapInterface
}
