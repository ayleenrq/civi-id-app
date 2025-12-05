package datasources

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UploadResult struct {
	URL string
}

type CloudinaryService interface {
	UploadImage(ctx context.Context, file *multipart.FileHeader, folder, filename string) (*UploadResult, error)
	UploadImageBytes(ctx context.Context, file io.Reader, folder, filename string) (*UploadResult, error)
}

type cloudinaryServiceImpl struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryService() (CloudinaryService, error) {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	if cloudName == "" || apiKey == "" || apiSecret == "" {
		return nil, errors.New("cloudinary credentials are not set")
	}
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cloudinary: %w", err)
	}
	return &cloudinaryServiceImpl{cld: cld}, nil
}

func (c *cloudinaryServiceImpl) UploadImage(ctx context.Context, file *multipart.FileHeader, folder, filename string) (*UploadResult, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	publicID := fmt.Sprintf("%s/%s", folder, filename)

	resp, err := c.cld.Upload.Upload(ctx, f, uploader.UploadParams{
		PublicID:     publicID,
		Folder:       folder,
		Overwrite:    boolPtr(true),
		ResourceType: "image",
	})
	if err != nil {
		return nil, err
	}

	return &UploadResult{URL: resp.SecureURL}, nil
}

func (c *cloudinaryServiceImpl) UploadImageBytes(ctx context.Context, file io.Reader, folder, filename string) (*UploadResult, error) {
	publicID := fmt.Sprintf("%s/%s", folder, filename)

	resp, err := c.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:     publicID,
		Folder:       folder,
		Overwrite:    boolPtr(true),
		ResourceType: "image",
	})
	if err != nil {
		return nil, err
	}

	return &UploadResult{URL: resp.SecureURL}, nil
}

func boolPtr(v bool) *bool {
	return &v
}

// NewCloudinaryClient returns raw *cloudinary.Cloudinary
func NewCloudinaryClient() (*cloudinary.Cloudinary, error) {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	if cloudName == "" || apiKey == "" || apiSecret == "" {
		return nil, errors.New("cloudinary credentials are not set")
	}

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cloudinary client: %w", err)
	}

	return cld, nil
}
