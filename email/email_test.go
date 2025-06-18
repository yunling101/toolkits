package email

import "testing"

func TestName(t *testing.T) {
	err := New(
		"",
		"",
		"",
		465,
	).
		Tls(true).
		To("").
		Subject("test1").
		//Body("text/plain", "a test email").
		Body("text/html", "<h1>hello world</h1>").
		Deliver()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success")
}
