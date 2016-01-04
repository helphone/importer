package transaction

import (
	"testing"
)

func TestCreateConnection(t *testing.T) {
	conn, err := CreateConnection()
	if err != nil {
		t.Errorf("The connection creation has an error %s", err)
	}

	conn.Finish(err)
}
