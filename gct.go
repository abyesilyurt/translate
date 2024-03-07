package main

import (
	"context"
	"fmt"
	"os"

	translate "cloud.google.com/go/translate/apiv3"
	"cloud.google.com/go/translate/apiv3/translatepb"
)

type TranslationRequest struct {
	TargetLang string
	SourceLang string
	Filename   string
}

func TranslateDocument(request TranslationRequest) ([]byte, error) {
	ctx := context.Background()
	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewTranslationClient: %w", err)
	}
	defer client.Close()

	data, err := os.ReadFile(request.Filename)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: %w", err)
	}

	projectID := os.Getenv("TRANSLATE_PROJECT_ID")
	if projectID == "" {
		return nil, fmt.Errorf("TRANSLATE_PROJECT_ID not set")
	}
	req := &translatepb.TranslateDocumentRequest{
		Parent:             fmt.Sprintf("projects/%s/locations/global", projectID),
		TargetLanguageCode: request.TargetLang,
		SourceLanguageCode: request.SourceLang,
		DocumentInputConfig: &translatepb.DocumentInputConfig{
			MimeType: "application/pdf",
			Source: &translatepb.DocumentInputConfig_Content{
				Content: data,
			},
		},
	}
	// _ = req
	// return nil, nil
	resp, err := client.TranslateDocument(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("TranslateDocument: %w", err)
	}
	return resp.DocumentTranslation.ByteStreamOutputs[0], nil

}
