package helpers
import (
	"time"
)

func GetCurrentTimeStamp() string{
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetImageDirectory() string{
	return time.Now().Format("2006/01/02/15")

}

func CurrentTimeMillis() int64{
	return time.Now().Unix()*1000
}