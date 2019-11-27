package sorry

import "testing"

func Unexpected(t *testing.T, result, expected string) {
	t.Logf("unexpected result")
	t.Logf("expected |%s|", expected)
	t.Logf("result   |%s|", result)
}

func UnexpectedNil(t *testing.T) {
	t.Logf("unexpected nil")
}

func UnexpectedBlank(t *testing.T) {
	t.Logf("unexpected blank string")
}

func UnexpectedError(t *testing.T, err error) {
	t.Logf("unexpected err: %s", err.Error())
}

func ExpectedError(t *testing.T, err error) {
	t.Logf("expected err: %s", err.Error())
}
