package store

import (
	"context"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"testing"
	"github.com/stretchr/testify/assert"
	"cloud.google.com/go/firestore"
	"github.com/dchest/uniuri"
)

func TestMain(m *testing.M) {
	// command to start firestore emulator
	cmd := exec.Command("gcloud", "beta", "emulators", "firestore", "start", "--host-port=localhost")

	// this makes it killable
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// we need to capture it's output to know when it's started
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stderr.Close()

	// start her up!
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// ensure the process is killed when we're finished, even if an error occurs
        // (thanks to Brian Moran for suggestion)
	var result int
	defer func() {
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		os.Exit(result)
	}()

	// we're going to wait until it's running to start
	var wg sync.WaitGroup
	wg.Add(1)

	// by starting a separate go routine
	go func() {
		// reading it's output
		buf := make([]byte, 256, 256)
		for {
			n, err := stderr.Read(buf[:])
			if err != nil {
				// until it ends
				if err == io.EOF {
					break
				}
				log.Fatalf("reading stderr %v", err)
			}

			if n > 0 {
				d := string(buf[:n])

				// only required if we want to see the emulator output
				log.Printf("%s", d)

				// checking for the message that it's started
				if strings.Contains(d, "Dev App Server is now running") {
					wg.Done()
				}

				// and capturing the FIRESTORE_EMULATOR_HOST value to set
				pos := strings.Index(d, FirestoreEmulatorHost+"=")
				if pos > 0 {
					host := d[pos+len(FirestoreEmulatorHost)+1 : len(d)-1]
					os.Setenv(FirestoreEmulatorHost, host)
				}
			}
		}
	}()

	// wait until the running message has been received
	wg.Wait()

	// now it's running, we can run our unit tests
	result = m.Run()
}

const FirestoreEmulatorHost = "FIRESTORE_EMULATOR_HOST"

func TestStoreNew(t *testing.T) {
	s := New()
	assert.NotEqual(t, "", s.Id, "New stores should have generated ID")
	assert.Equal(t, "", s.Name, "New stores should not have name set")
	assert.Equal(t, "", s.City, "New stores should not have city set")
	assert.Equal(t, "", s.Zipcode, "New stores should not have zipcode set")
}

func TestStoreLoad(t *testing.T) {
	// new client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "price-comparator-dev")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	defer client.Close()
	// new store
	newStoreId := uniuri.New()
	newStoreName := uniuri.New()
	newStoreCity := uniuri.New()
	newStoreCode := uniuri.New()
	store := client.Collection("store").Doc(newStoreId)
	_, err = store.Create(ctx, map[string]interface{}{
			"name":  newStoreName,
			"city": newStoreCity,
			"zipcode": newStoreCode,
		})
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	// load
	s, err := Load(store.ID)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	// checks
	assert.Equal(t, newStoreId, s.Id)
	assert.Equal(t, newStoreName, s.Name)
	assert.Equal(t, newStoreCity, s.City)
	assert.Equal(t, newStoreCode, s.Zipcode)
	// cleanup
	_, err = store.Delete(ctx)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
}
