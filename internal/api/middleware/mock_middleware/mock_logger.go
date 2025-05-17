package mock_middleware

import context "context"

type MockLogger struct {
	Infos []string
}

func (m *MockLogger) Debug(ctx context.Context, msg string)   {}
func (m *MockLogger) Info(ctx context.Context, msg string)    { m.Infos = append(m.Infos, msg) }
func (m *MockLogger) Warning(ctx context.Context, msg string) {}
func (m *MockLogger) Error(ctx context.Context, msg string)   {}
func (m *MockLogger) Fatal(ctx context.Context, msg string)   {}
func (m *MockLogger) Close(ctx context.Context) error         { return nil }
