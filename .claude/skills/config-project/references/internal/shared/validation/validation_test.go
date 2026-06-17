package validation_test

import (
	"_/C_/Users/bsgle/OneDrive/Cowork/short-url-by-ssd/.claude/skills/config-project/references/internal/shared/validation"
	"testing"
)

func TestNewMessage(t *testing.T) {
	msg := validation.NewMessage("CODE1", validation.Error, "failed")

	if msg == nil {
		t.Fatal("expected non-nil message")
	}

	if msg.Code != "CODE1" {
		t.Fatalf("expected code %q, got %q", "CODE1", msg.Code)
	}

	if msg.Type != validation.Error {
		t.Fatalf("expected type Error, got %v", msg.Type)
	}

	if msg.Description != "failed" {
		t.Fatalf("expected description %q, got %q", "failed", msg.Description)
	}
}

func TestNewUseCaseResult(t *testing.T) {
	result := validation.NewUseCaseResult()

	if result == nil {
		t.Fatal("expected non-nil result")
	}

	if result.HasError {
		t.Fatal("expected HasError false")
	}

	if result.HasWarning {
		t.Fatal("expected HasWarning false")
	}

	if len(result.Messages) != 0 {
		t.Fatalf("expected 0 messages, got %d", len(result.Messages))
	}

	if result.HasMessage() {
		t.Fatal("expected HasMessage false")
	}
}

func TestNewFailUseCaseResult(t *testing.T) {
	result := validation.NewFailUseCaseResult("VALIDATION", "invalid input")

	if !result.HasError {
		t.Fatal("expected HasError true")
	}

	if result.HasWarning {
		t.Fatal("expected HasWarning false")
	}

	if len(result.Messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(result.Messages))
	}

	msg := result.Messages[0]
	if msg.Code != "VALIDATION" {
		t.Fatalf("expected message code %q, got %q", "VALIDATION", msg.Code)
	}

	if msg.Type != validation.Error {
		t.Fatalf("expected message type Error, got %v", msg.Type)
	}

	if msg.Description != "invalid input" {
		t.Fatalf("expected description %q, got %q", "invalid input", msg.Description)
	}

	if !result.HasMessage() {
		t.Fatal("expected HasMessage true")
	}
}

func TestNewWarninglUseCaseResult(t *testing.T) {
	result := validation.NewWarninglUseCaseResult("WARN", "careful")

	if result.HasError {
		t.Fatal("expected HasError false")
	}

	if !result.HasWarning {
		t.Fatal("expected HasWarning true")
	}

	if len(result.Messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(result.Messages))
	}

	msg := result.Messages[0]
	if msg.Code != "WARN" {
		t.Fatalf("expected message code %q, got %q", "WARN", msg.Code)
	}

	if msg.Type != validation.Warning {
		t.Fatalf("expected message type Warning, got %v", msg.Type)
	}

	if msg.Description != "careful" {
		t.Fatalf("expected description %q, got %q", "careful", msg.Description)
	}
}

func TestUseCaseResult_AddMessage(t *testing.T) {
	result := validation.NewUseCaseResult()
	result.AddMessage(validation.NewMessage("CODE2", validation.Warning, "check this"))

	if len(result.Messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(result.Messages))
	}

	if !result.HasWarning {
		t.Fatal("expected HasWarning true after adding warning message")
	}

	if result.HasError {
		t.Fatal("expected HasError false after adding warning message")
	}

	result.AddMessage(validation.NewMessage("ERR", validation.Error, "bad"))
	if !result.HasError {
		t.Fatal("expected HasError true after adding error message")
	}

	if len(result.Messages) != 2 {
		t.Fatalf("expected 2 messages, got %d", len(result.Messages))
	}
}

func TestUseCaseResult_AddErrorAndAddWarning(t *testing.T) {
	result := validation.NewUseCaseResult()

	result.AddError("ERR1", "error happened")
	if len(result.Messages) != 1 {
		t.Fatalf("expected 1 message after AddError, got %d", len(result.Messages))
	}

	if result.Messages[0].Type != validation.Error {
		t.Fatalf("expected message type Error, got %v", result.Messages[0].Type)
	}

	result.AddWarning("WARN1", "warning issued")
	if len(result.Messages) != 2 {
		t.Fatalf("expected 2 messages after AddWarning, got %d", len(result.Messages))
	}

	if !result.HasMessage() {
		t.Fatal("expected HasMessage true after adding messages")
	}
}
