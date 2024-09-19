package mocks

type MockLogger struct {
	FatalCalled bool
	Args        []interface{}
}

func (m *MockLogger) Fatal(v ...interface{}) {
	m.FatalCalled = true
	m.Args = v
	panic("Fatal called.")
}

func (m *MockLogger) Println(v ...interface{}) {
	m.Args = v
}