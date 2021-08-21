package helper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func init() {
	Wg = &sync.WaitGroup{}
	Ctx, CloseCtx = context.WithCancel(context.Background())
	C = make(chan os.Signal, 3)
}

func CreateConn(addr string) (conn *grpc.ClientConn, ctx context.Context, closeCtx context.CancelFunc, err error) {
	timer := time.NewTimer(time.Minute * 5)
	defer timer.Stop()

	ctx, closeCtx = context.WithCancel(Ctx)
	c := make(chan struct{}, 1)

	Wg.Add(1)
	go func() {
		defer Wg.Done()
		for {
			select {
			case <-Ctx.Done():
				closeCtx()
				return
			case <-timer.C:
				closeCtx()
				return
			case <-c:
				close(c)
				return
			}
		}
	}()

	conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	c <- struct{}{}
	return
}

func ParseConfig(path string, out interface{}) error {
	_, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf(KeyErrorNotFound+": file: %v", path)
	}
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf(KeyErrorOpen+": file: %v", err)
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf(KeyErrorRead+": body: %v", err)
	}

	err = json.Unmarshal(buf, out)
	if err != nil {
		return fmt.Errorf(KeyErrorParse+": json: %v", err)
	}

	return nil
}

func Dispatch(f interface{}, args ...interface{}) error {
	ff := reflect.ValueOf(f)
	if ff.Kind() == reflect.Func {
		in := make([]reflect.Value, ff.Type().NumIn())
		for i, arg := range args {
			v := reflect.ValueOf(arg)
			if v.Type().ConvertibleTo(ff.Type().In(i)) {
				in[i] = v.Convert(ff.Type().In(i))
			} else {
				return fmt.Errorf("parameter: %v, expected %v got %v", i+1, ff.Type().In(i), v.Type())
			}
		}

		Wg.Add(1)
		go func() {
			ff.Call(in)
			Wg.Done()
		}()
	}
	return nil
}

func Hash(data, salt string) string {
	hash1 := sha256.New()
	hash2 := sha256.New()

	hash1.Write([]byte(data))
	hash2.Write(hash1.Sum([]byte(salt)))
	hash1.Write(hash2.Sum([]byte(salt)))

	return hex.EncodeToString(hash1.Sum(nil))
}
