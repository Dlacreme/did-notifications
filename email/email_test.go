package email

import (
	"testing"

	"github.com/dlacreme/did-notifications/env"
)

func init() {
	env.LoadAndCheck()
}

func TestSendEmail(t *testing.T) {
	to := []string{"test+in@test.com"}
	if err := SendEmail("test+out@test.com", to, "Test email", "This is a test content"); err != nil {
		t.Fatalf("Send email fail. Stacktrace:\n%v", err)
	}
}
