package base



import (
	"errors"
)

func IsTrue(v bool, msg string) error {
	if !v {
		return errors.New(msg)
	}
	return nil
}

func IsNotEmpty(v string, param string) error {
	if v == "" {
		return errors.New(param + " is empty!")
	}
	return nil

}

func IsEmpty(v string, param string) error {
	var err = IsNotEmpty(v, param)
	if err == nil {
		return errors.New(param + "is not empty!")
	}
	return nil
}

func HasAnyString(v []string, param string) error {
	if v == nil {
		return errors.New(param + " is null!")
	}
	if len(v) == 0 {
		return errors.New(param + " count must > 0!")
	}
	return nil
}

func HasAnyInt(v []int, param string) error {
	if v == nil {
		return errors.New(param + " is null!")
	}
	if len(v) == 0 {
		return errors.New(param + " count must > 0!")
	}
	return nil
}
