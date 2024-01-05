package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/EdimarRibeiro/inventory/api/common"
	"github.com/EdimarRibeiro/inventory/api/models"
	entitiesinterface "github.com/EdimarRibeiro/inventory/internal/interfaces/entities"
	"github.com/EdimarRibeiro/inventory/internal/internalfunc"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

type fileUploadController struct {
	tenant entitiesinterface.TenantRepositoryInterface
}

func CreateFileUploadController(tenantRep entitiesinterface.TenantRepositoryInterface) *fileUploadController {
	return &fileUploadController{tenant: tenantRep}
}

func uploadURL(sess *session.Session, file []byte, bucket string, filename string, duration time.Duration) (string, error) {
	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   bytes.NewReader(file),
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return "", err
	}
	return result.Location, nil
}

func (repo *fileUploadController) HandleFileUpload(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	tenants, err := repo.tenant.Search("Id = " + strconv.FormatUint(tenantId, 10))
	if err != nil {
		http.Error(w, "Error retrieving tenant "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(tenants) == 0 {
		http.Error(w, "Notfound tenant", http.StatusMethodNotAllowed)
		return
	}
	fileName := uuid.New().String() + "_" + r.Header.Get("x-file-name")
	if fileName == "" {
		// If the header is not present, try to get it from the query parameter
		fileName = uuid.New().String() + "_" + r.URL.Query().Get("fileName")
	}
	if fileName == "" {
		fileName = uuid.New().String()
	}
	document := tenants[0].Document

	fileBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusInternalServerError)
		return
	}

	configSpace := internalfunc.InitGent()
	duration, err := time.ParseDuration("9600h") //20 dias
	if err != nil {
		http.Error(w, fmt.Sprintf("Error configure duration: %v", err), http.StatusInternalServerError)
		return
	}
	if duration < 0 {
		http.Error(w, "Duration is negative", http.StatusMethodNotAllowed)
		return
	}

	config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(configSpace.Key, configSpace.Secret, ""),
		Endpoint:    aws.String(configSpace.Host),
		Region:      aws.String(configSpace.Region),
	}

	sess, err := session.NewSession(config)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating AWS session: %v", err), http.StatusInternalServerError)
		return
	}

	url, err := uploadURL(sess, fileBytes, configSpace.Bucket, document+"/"+fileName, duration)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error uploading file to S3: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ResponseUploadFile{Url: url})
}

func (repo *fileUploadController) HandleFileDownload(w http.ResponseWriter, r *http.Request) {
	_, _, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	file := models.ResponseUploadFile{}

	if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
		http.Error(w, "Error decoding JSON "+err.Error(), http.StatusBadRequest)
		return
	}

	configSpace := internalfunc.InitGent()
	config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(configSpace.Key, configSpace.Secret, ""),
		Endpoint:    aws.String(configSpace.Host),
		Region:      aws.String(configSpace.Region),
	}

	sess, err := session.NewSession(config)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating AWS session: %v", err), http.StatusInternalServerError)
		return
	}

	fileContent, err := internalfunc.DownloadURL(sess, configSpace.Bucket, file.Url, 0)

	if err != nil {
		http.Error(w, "download error:"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "file"))
	w.Header().Set("Content-Type", http.DetectContentType(fileContent))
	w.Write(fileContent)
}
