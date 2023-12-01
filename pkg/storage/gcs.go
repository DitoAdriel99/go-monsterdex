// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

// Compile-time check to verify implements interface.
var _ Storage = (*gcs)(nil)

// gcs implements the Blob interface and provides the ability
// write files to Google Cloud Storage.
type gcs struct {
	client *storage.Client
}

// NewGCS creates a Google Cloud Storage Client
func NewGCS(ctx context.Context) (Storage, error) {
	credOpt := option.WithCredentialsFile(os.Getenv("ACCOUNT_PATH"))

	client, err := storage.NewClient(ctx, credOpt)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %w", err)
	}

	// Stat bucket if not exist create a new one.
	// https://cloud.google.com/storage/docs/creating-buckets#storage-create-bucket-go
	_, err = client.Bucket(os.Getenv("GCS_BUCKET")).Attrs(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrBucketNotExist) {
			err = client.Bucket(os.Getenv("GCS_BUCKET")).Create(ctx, os.Getenv("PROJECTID"), nil)
			if err != nil {
				return nil, fmt.Errorf("storage.Bucket.Create: %w", err)
			}
		} else {
			return nil, fmt.Errorf("storage.Bucket.Attrs: %w", err)
		}
	}

	return &gcs{client}, nil
}

// Put creates a new cloud storage object or overwrites an existing one.
func (s *gcs) Put(ctx context.Context, bucket, objectName string, contents []byte, cacheable bool, contentType string) error {
	cacheControl := "public, max-age=86400"
	if !cacheable {
		cacheControl = "no-cache, max-age=0"
	}

	wc := s.client.Bucket(bucket).Object(objectName).NewWriter(ctx)
	wc.CacheControl = cacheControl
	if contentType != "" {
		wc.ContentType = contentType
	}

	if _, err := wc.Write(contents); err != nil {
		return fmt.Errorf("storage.Writer.Write: %w", err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("storage.Writer.Close: %w", err)
	}

	return nil
}

// Delete deletes a cloud storage object, returns nil if the object was
// successfully deleted, or of the object doesn't exist.
func (s *gcs) Delete(ctx context.Context, bucket, objectName string) error {
	if err := s.client.Bucket(bucket).Object(objectName).Delete(ctx); err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			// Object doesn't exist; presumably already deleted.
			return nil
		}
		return fmt.Errorf("storage.DeleteObject: %w", err)
	}
	return nil
}

// Get returns the contents for the given object. If the object does not
// exist, it returns ErrNotFound.
func (s *gcs) Get(ctx context.Context, bucket, object string) ([]byte, error) {
	baseName := filepath.Base(object)
	r, err := s.client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return nil, ErrNotFound
		}
	}
	defer r.Close()

	createFileContent(baseName, r)

	var b bytes.Buffer
	if _, err := io.Copy(&b, r); err != nil {
		return nil, fmt.Errorf("failed to download bytes: %w", err)
	}

	return b.Bytes(), nil
}

// // ResignUrl returns the URL for the given object.
func (s *gcs) ResignUrl(ctx context.Context, bucket, objectName string) (string, error) {
	saKey, err := ioutil.ReadFile(os.Getenv("ACCOUNT_PATH"))
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := google.JWTConfigFromJSON(saKey)
	if err != nil {
		log.Fatalln(err)
	}
	url, err := storage.SignedURL(bucket, objectName, &storage.SignedURLOptions{
		GoogleAccessID: cfg.Email,
		PrivateKey:     []byte(cfg.PrivateKey),
		Method:         "GET",
		Expires:        time.Now().Add(1 * time.Hour),
	})

	if err != nil {
		return "", fmt.Errorf("storage.SignedURL: %w", err)
	}

	if url == "" {
		return "", fmt.Errorf("storage.SignedURL: url is empty")
	}

	return url, nil
}

func createFileContent(baseName string, r *storage.Reader) error {
	if _, err := os.Stat("temporaryDownloadFile/"); os.IsNotExist(err) {
		errDir := os.Mkdir("temporaryDownloadFile/", 0777)
		if errDir != nil {
			log.Fatalf("error while create directory, err %v", errDir)
			return err
		}
	}

	filePath := filepath.Join("temporaryDownloadFile/", baseName)
	// Create a new local file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the GCS object content to the local file
	if _, err := io.Copy(file, r); err != nil {
		return err
	}

	return nil
}
func StoreDataToGCS(ctx context.Context, bucket, objectName string, contents []byte, cacheable bool, contentType string) (*string, error) {
	gcs, err := NewGCS(ctx)
	if err != nil {
		return nil, err
	}

	err = gcs.Put(ctx, bucket, objectName, contents, false, fmt.Sprintf("image/%s", contentType))
	if err != nil {
		return nil, err
	}

	return &objectName, nil
}

func SignedURL(ctx context.Context, bucket, objectName string) (*string, error) {
	gcs, err := NewGCS(ctx)
	if err != nil {
		return nil, err
	}

	url, err := gcs.ResignUrl(ctx, bucket, objectName)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

func GetURL(ctx context.Context, bucket, objectName string) (*string, error) {
	gcs, err := NewGCS(ctx)
	if err != nil {
		return nil, err
	}

	url, err := gcs.Get(ctx, bucket, objectName)
	if err != nil {
		return nil, err
	}

	urlString := string(url)

	return &urlString, nil
}
