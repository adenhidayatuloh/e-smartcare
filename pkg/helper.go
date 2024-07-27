package pkg

import (
	"esmartcare/pkg/errs"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func ValidateStruct(payload interface{}) errs.MessageErr {
	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}
	return nil
}

func ValidateDate(i interface{}, context interface{}) bool {

	Date, ok := i.(string)
	if !ok {
		return false
	}
	// Define the expected date format
	dateFormat := "2006-01-02"

	// Try parsing the date string using the specified format
	_, err := time.Parse(dateFormat, Date)

	print(err.Error())

	// Check if parsing was successful
	// Check if parsing was successful
	if err == nil {
		return true
	} else {
		return false
	}

}

func ValidateJenisAkun(i interface{}, context interface{}) bool {
	jenisAkun, ok := i.(string)
	if !ok {
		return false
	}
	return jenisAkun == "admin" || jenisAkun == "siswa" || jenisAkun == "pakar"
}

func ValidateStatusAlarm(i interface{}, context interface{}) bool {
	statusAlarm, ok := i.(string)
	if !ok {
		return false
	}
	return statusAlarm == "1" || statusAlarm == "0"
}

func init() {
	// Register the custom validation function
	govalidator.CustomTypeTagMap.Set("jenisAkunValidator", govalidator.CustomTypeValidator(ValidateJenisAkun))
	govalidator.CustomTypeTagMap.Set("date", govalidator.CustomTypeValidator(ValidateDate))
	govalidator.CustomTypeTagMap.Set("statusAlarm", govalidator.CustomTypeValidator(ValidateStatusAlarm))
}

func DeleteImage(filename string) errs.MessageErr {

	filepath := "." + filename

	if err := os.Remove(filepath); err != nil {
		if os.IsNotExist(err) {
			return errs.NewNotFound("Error : file not found ")
		} else {
			return errs.NewInternalServerError("Error : file cannot delete ")
		}

	}

	return nil

}

func UploadImage(formFile string, editFileName string, ctx *gin.Context) (*string, errs.MessageErr) {
	// Handle file upload

	file, err := ctx.FormFile(formFile)

	if err != nil && err != http.ErrMissingFile {

		return nil, errs.NewInternalServerError(err.Error())
	}

	if file == nil {

		kosong := ""
		return &kosong, nil
	}

	// Generate a filename using the user's email and keep the original file extension
	extension := filepath.Ext(file.Filename)

	if !(extension == ".png" || extension == ".jpg" || extension == ".jpeg" || extension == ".webp" || extension == ".JPG") {

		return nil, errs.NewBadRequest("Format file not suported, please upload only image type")

	}

	newFilename := fmt.Sprintf("%s%s", editFileName, extension)
	newFilename = strings.ReplaceAll(newFilename, " ", "")

	// Save the file to the server with the new filename
	if err := ctx.SaveUploadedFile(file, "./uploads/"+newFilename+"-temp"); err != nil {

		return nil, errs.NewInternalServerError(err.Error())
	}

	finalRoutes := "/uploads/" + newFilename + "-temp"

	return &finalRoutes, nil
}

// RenameImage renames the temporary uploaded image to the desired filename.
func RenameImage(tempFilename, newFilename string) errs.MessageErr {

	tempFilename = "." + tempFilename
	newFilename = "." + newFilename
	if err := os.Rename(tempFilename, newFilename); err != nil {
		return errs.NewInternalServerError("Cannot rename file")
	}
	return nil
}

func CreateIndex(indexPath string, indexMapping mapping.IndexMapping) (bleve.Index, error) {
	index, err := bleve.New(indexPath, indexMapping)
	if err != nil {
		return nil, err
	}
	return index, nil
}

// deleteAllDocuments menghapus semua dokumen dari indeks
func DeleteAllDocuments(index bleve.Index) error {
	// Retrieve all documents in the index
	batch := index.NewBatch()
	searchRequest := bleve.NewSearchRequest(bleve.NewMatchAllQuery())
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		return err
	}

	// Add all documents to the batch for deletion
	for _, hit := range searchResult.Hits {
		batch.Delete(hit.ID)
	}

	// Execute the batch delete operation
	err = index.Batch(batch)
	if err != nil {
		return err
	}

	return nil
}

func UploadImagePemeriksaan(formFile string, editFileName string, ctx *gin.Context) (*string, errs.MessageErr) {
	// Handle file upload

	file, err := ctx.FormFile(formFile)

	if err != nil && err != http.ErrMissingFile {

		return nil, errs.NewInternalServerError(err.Error())
	}

	if file == nil {

		kosong := ""
		return &kosong, nil
	}

	// Generate a filename using the user's email and keep the original file extension
	extension := filepath.Ext(file.Filename)

	if !(extension == ".png" || extension == ".jpg" || extension == ".jpeg" || extension == ".webp" || extension == ".JPG") {

		return nil, errs.NewBadRequest("Format file not suported, please upload only image type")

	}

	newFilename := fmt.Sprintf("%s%s", editFileName, extension)
	newFilename = strings.ReplaceAll(newFilename, " ", "")

	// Save the file to the server with the new filename
	if err := ctx.SaveUploadedFile(file, "./uploads/pemeriksaan/"+newFilename+"-temp"); err != nil {

		return nil, errs.NewInternalServerError(err.Error())
	}

	finalRoutes := "/uploads/pemeriksaan/" + newFilename + "-temp"

	return &finalRoutes, nil
}
