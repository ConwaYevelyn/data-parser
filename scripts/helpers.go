package data_parser

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

//go:embed static
var staticFS embed.FS

func copyFile(src, dst string) error {
	in, err := staticFS.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		return err
	}

	if err := out.Close(); err != nil {
		return err
	}

	return nil
}

func validateError(err error) error {
	if err != nil {
		if strings.Contains(err.Error(), "Key not found") || strings.Contains(err.Error(), "Unknown") {
			return errors.New("validation failed")
		}
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func parseStringSlice(s string, sep []rune) ([]string, error) {
	pattern := regexp.QuoteMeta(string(sep))
	re := regexp.MustCompile(pattern)
	return re.Split(s, -1), nil
}

func intSliceToString(sl []int) (string, error) {
	var sb strings.Builder
	for _, s := range sl {
		if err := sb.WriteByte(strconv.Itoa(s)[0]); err != nil {
			return "", err
		}
	}
	return sb.String(), nil
}

func validateString(str string, pattern *regexp.Regexp) error {
	if pattern.MatchString(str) {
		return nil
	}
	return errors.New("invalid string")
}

func validateInt(v int) (int, error) {
	if v < 0 {
		return 0, errors.New("invalid int")
	}
	return v, nil
}

func isDirEmpty(p string) bool {
	files, err := os.ReadDir(p)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Println(err)
		}
		return true
	}
	return len(files) == 0
}

func getUInt64(b []byte) (uint64, error) {
	if len(b) > 8 {
		b = b[:8]
	}
	var e error
	if b[0] == 0 {
		e = errors.New("zero value")
	} else {
		var uint64 int64
		uint64, e = strconv.ParseInt(string(b), 10, 64)
	}
	return uint64(uint64), e
}

func validateEmail(email string) bool {
	regex := regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")
	return regex.MatchString(email)
}

func validatePassword(password string) bool {
	regex := regexp.MustCompile("^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]{8,}$")
	return regex.MatchString(password)
}

func validateStringSlice(sl []string, pattern *regexp.Regexp) ([]string, error) {
	var err error
	for _, s := range sl {
		if err := validateError(pattern.MatchString(s)); err != nil {
			return nil, err
		}
	}
	return sl, nil
}

func validateIntSlice(sl []int, min, max int) ([]int, error) {
	var err error
	for _, v := range sl {
		if err := validateError(validateInt(v)); err != nil {
			return nil, err
		}
		if v < min || v > max {
			return nil, errors.New("out of range")
		}
	}
	return sl, nil
}

func validateStringMap(sl map[string]string, pattern *regexp.Regexp) (map[string]string, error) {
	var err error
	for k, v := range sl {
		if err := validateError(validateString(v, pattern)); err != nil {
			return nil, err
		}
	}
	return sl, nil
}

func getValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("email", func(fl validator.FieldLevel) bool {
		return validateEmail(fl.Field().String())
	})
	v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		return validatePassword(fl.Field().String())
	})
	return v
}