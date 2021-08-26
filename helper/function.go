package helper

import (
	"context"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials"
)

func init() {
	Wg = &sync.WaitGroup{}
	Ctx, CloseCtx = context.WithCancel(context.Background())
	C = make(chan os.Signal, 3)
	signal.Notify(C, os.Interrupt)
}

func CreateConn(addr string) (conn *grpc.ClientConn, ctx context.Context, closeCtx context.CancelFunc, err error) {
	// cred, err := CredentialsClient()
	// if err != nil {
	// 	return nil, nil, nil, err
	// }

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
		// grpc.WithTransportCredentials(cred),
		grpc.WithInsecure(),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.DefaultConfig,
		}),
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

func CredentialsServer() (credentials.TransportCredentials, error) {
	ca, err := ioutil.ReadFile(CA)
	if err != nil {
		return nil, fmt.Errorf(KeyErrorRead + " ca")
	}

	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(ca) {
		return nil, fmt.Errorf("cannot append CA")
	}

	cert, err := tls.LoadX509KeyPair(Server, ServerKey)
	if err != nil {
		return nil, fmt.Errorf(KeyErrorRead + " server cert")
	}

	conf := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    pool,
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(conf), nil
}

func CredentialsClient() (credentials.TransportCredentials, error) {
	ca, err := ioutil.ReadFile(CA)
	if err != nil {
		return nil, fmt.Errorf(KeyErrorRead + " ca")
	}

	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(ca) {
		return nil, fmt.Errorf("cannot append CA")
	}

	cert, err := tls.LoadX509KeyPair(Client, ClientKey)
	if err != nil {
		return nil, fmt.Errorf(KeyErrorRead + " server cert")
	}

	conf := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "127.0.0.1",
		RootCAs:      pool,
	}

	return credentials.NewTLS(conf), nil
}
