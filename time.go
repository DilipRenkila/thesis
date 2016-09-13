package main
import "time"
import "fmt"
import "log"
import "os"
import (
	"strconv"
	"reflect"
)

type Timestamp time.Time

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t).Unix()
	stamp := fmt.Sprint(ts)

	return []byte(stamp), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	*t = Timestamp(time.Unix(int64(ts), 0).UTC())
	fmt.Println(&t)
	return nil
}



func main() {
	t := Timestamp(time.Now().UTC())
	fmt.Println(reflect.TypeOf(t),t)
	x,err := t.MarshalJSON()
		if err != nil {
		log.Println("error: %s", err)
		os.Exit(1)
	}
	err = t.UnmarshalJSON(x)
			if err != nil {
		log.Println("error: %s", err)
		os.Exit(1)
	}
}