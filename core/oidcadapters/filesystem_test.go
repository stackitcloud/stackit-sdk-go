package oidcadapters

import (
	"context"
	"log"
	"os"
	"testing"
)

func TestReadJWTFromFileSystem(t *testing.T) {
	file, err := os.CreateTemp("", "*.token")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := os.Remove(file.Name())
		if err != nil {
			t.Fatalf("Removing temporary file: %s", err)
		}
	}()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30" // nolint:gosec // This is a fake token for testing purposes only
	err = os.WriteFile(file.Name(), []byte(token), os.ModeAppend)
	if err != nil {
		t.Fatalf("Writing temporary file: %s", err)
	}
	_, err = ReadJWTFromFileSystem(file.Name())(context.Background())
	if err != nil {
		t.Fatalf("Reading JWT from file system: %s", err)
	}
}

func TestReadRandomContentFromFileSystem(t *testing.T) {
	file, err := os.CreateTemp("", "*.token")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := os.Remove(file.Name())
		if err != nil {
			t.Fatalf("Removing temporary file: %s", err)
		}
	}()

	token := "invalid random content"
	err = os.WriteFile(file.Name(), []byte(token), os.ModeAppend)
	if err != nil {
		t.Fatalf("Writing temporary file: %s", err)
	}

	_, err = ReadJWTFromFileSystem(file.Name())(context.Background())
	if err == nil {
		t.Fatalf("Reading JWT from file system must fail")
	}
}

func TestReadMissingFileFromFileSystem(t *testing.T) {
	_, err := ReadJWTFromFileSystem("/path/to/nonexistent/file.token")(context.Background())
	if err == nil {
		t.Fatalf("Reading JWT from file system must fail")
	}
}
