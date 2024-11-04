// internal/adapters/pg/supabase.go
package pg

import (
    "bytes"
    "fmt"
    "io"
    "mime/multipart"
    "net/http"

    "github.com/kritpi/arom-web-services/configs"
)

func UploadImageToSupabaseV2(file multipart.File, fileName, bucket string, config *configs.Config) (string, error) {
    // Read file contents
    fileBytes, err := io.ReadAll(file)
    if err != nil || len(fileBytes) == 0 {
        return "", fmt.Errorf("file is empty or unreadable")
    }

    // Prepare request to Supabase
    url := fmt.Sprintf("%s/storage/v1/object/%s/%s", config.SUPABASE_URL, bucket, fileName)
    req, err := http.NewRequest("POST", url, bytes.NewReader(fileBytes))
    if err != nil {
        return "", fmt.Errorf("failed to create upload request: %w", err)
    }

    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.SUPABASE_API_KEY))
    req.Header.Set("Content-Type", "application/octet-stream") // Set content type for binary upload

    // Send request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("upload failed: %w", err)
    }
    defer resp.Body.Close()

    // Check for successful status code
    if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
        body, _ := io.ReadAll(resp.Body)
        return "", fmt.Errorf("upload failed with status %d: %s", resp.StatusCode, string(body))
    }

    // Image URL format for Supabase storage
    imageURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", config.SUPABASE_URL, bucket, fileName)
    return imageURL, nil
}
