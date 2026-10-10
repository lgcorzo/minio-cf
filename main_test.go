/*
 * Minio tool to be used with Pivotal "cf" tool to generate the
 * config for create-service command.
 */

package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func testAppWithArgs(args []string) (string, error) {
	app := newApp()
	var buf bytes.Buffer
	app.Writer = &buf

	err := app.Run(args)
	return buf.String(), err
}

func TestAppServerConfig(t *testing.T) {
	args := []string{"minio-cf", "--access-key", "myaccess", "--secret-key", "mysecret"}
	output, err := testAppWithArgs(args)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var m map[string]string
	if err := json.Unmarshal([]byte(output), &m); err != nil {
		t.Fatalf("failed to unmarshal JSON output: %v, raw: %s", err, output)
	}

	if m["accesskey"] != "myaccess" {
		t.Errorf("expected accesskey 'myaccess', got '%s'", m["accesskey"])
	}
	if m["secretkey"] != "mysecret" {
		t.Errorf("expected secretkey 'mysecret', got '%s'", m["secretkey"])
	}
}

func TestAppAzureGateway(t *testing.T) {
	args := []string{"minio-cf", "--access-key", "azureacc", "--secret-key", "azurekey", "--azure"}
	output, err := testAppWithArgs(args)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var m map[string]string
	if err := json.Unmarshal([]byte(output), &m); err != nil {
		t.Fatalf("failed to unmarshal JSON output: %v, raw: %s", err, output)
	}

	if m["gateway"] != "azure" {
		t.Errorf("expected gateway 'azure', got '%s'", m["gateway"])
	}
}

func TestAppS3Gateway(t *testing.T) {
	args := []string{"minio-cf", "--access-key", "s3acc", "--secret-key", "s3key", "--s3", "--subdomain", "mysub"}
	output, err := testAppWithArgs(args)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var m map[string]string
	if err := json.Unmarshal([]byte(output), &m); err != nil {
		t.Fatalf("failed to unmarshal JSON output: %v, raw: %s", err, output)
	}

	if m["gateway"] != "s3" {
		t.Errorf("expected gateway 's3', got '%s'", m["gateway"])
	}
	if m["subdomain"] != "mysub" {
		t.Errorf("expected subdomain 'mysub', got '%s'", m["subdomain"])
	}
}

func TestAppGCSGateway(t *testing.T) {
	tmpDir := t.TempDir()
	credPath := filepath.Join(tmpDir, "gcs-creds.json")
	credContent := []byte(`{"type": "service_account", "project_id": "test"}`)
	if err := os.WriteFile(credPath, credContent, 0600); err != nil {
		t.Fatalf("failed to write temp cred file: %v", err)
	}

	args := []string{"minio-cf", "--access-key", "gcsacc", "--secret-key", "gcskey", "--gcs", credPath}
	output, err := testAppWithArgs(args)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var m map[string]string
	if err := json.Unmarshal([]byte(output), &m); err != nil {
		t.Fatalf("failed to unmarshal JSON output: %v, raw: %s", err, output)
	}

	if m["gateway"] != "gcs" {
		t.Errorf("expected gateway 'gcs', got '%s'", m["gateway"])
	}

	expectedCreds := base64.StdEncoding.EncodeToString(credContent)
	if m["googlecredentials"] != expectedCreds {
		t.Errorf("expected googlecredentials '%s', got '%s'", expectedCreds, m["googlecredentials"])
	}
}
