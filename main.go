package main

import "esmartcare/handler"

func main() {
	handler.StartApp()

}

// package main

// // Import Cloudinary and other necessary libraries
// //===================
// import (
// 	"context"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"

// 	"github.com/cloudinary/cloudinary-go/v2"
// 	"github.com/cloudinary/cloudinary-go/v2/api"
// 	"github.com/cloudinary/cloudinary-go/v2/api/admin"
// 	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
// )

// func credentials() (*cloudinary.Cloudinary, context.Context) {
// 	// Add your Cloudinary credentials, set configuration parameter
// 	// Secure=true to return "https" URLs, and create a context
// 	//===================
// 	cld, _ := cloudinary.NewFromParams("dciv82xna", "843286689852972", "rBKo_sWQWaVO61GZfBfzutYl-zY")
// 	cld.Config.URL.Secure = true
// 	ctx := context.Background()
// 	return cld, ctx
// }

// func uploadImage(cld *cloudinary.Cloudinary, ctx context.Context) {

// 	// Upload the image.
// 	// Set the asset's public ID and allow overwriting the asset with new versions
// 	resp, err := cld.Upload.Upload(ctx, "uploads/adminku@gmail.com.jpeg", uploader.UploadParams{
// 		PublicID:       "adminku@gmail.com",
// 		UniqueFilename: api.Bool(false),
// 		Overwrite:      api.Bool(true)})
// 	if err != nil {
// 		fmt.Println("error")
// 	}

// 	// Log the delivery URL
// 	fmt.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL, "\n")

// }

// func getImageInfo(cld *cloudinary.Cloudinary, ctx context.Context, publicID string) (string, error) {
// 	// Mendapatkan informasi tentang asset yang sudah di-upload
// 	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: publicID})
// 	if err != nil {
// 		return "", err
// 	}
// 	return resp.SecureURL, nil
// }

// // func getAssetInfo(cld *cloudinary.Cloudinary, ctx context.Context) {
// // 	// Get and use details of the image
// // 	// ==============================
// // 	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: "quickstart_butterfly"})
// // 	if err != nil {
// // 		fmt.Println("error")
// // 	}
// // 	fmt.Println("****3. Get and use details of the image****\nDetailed response:\n", resp, "\n")

// // 	// Assign tags to the uploaded image based on its width. Save the response to the update in the variable 'update_resp'.
// // 	if resp.Width > 900 {
// // 		update_resp, err := cld.Admin.UpdateAsset(ctx, admin.UpdateAssetParams{
// // 			PublicID: "quickstart_butterfly",
// // 			Tags:     []string{"large"}})
// // 		if err != nil {
// // 			fmt.Println("error")
// // 		} else {
// // 			// Log the new tag to the console.
// // 			fmt.Println("New tag: ", update_resp.Tags, "\n")
// // 		}
// // 	} else {
// // 		update_resp, err := cld.Admin.UpdateAsset(ctx, admin.UpdateAssetParams{
// // 			PublicID: "quickstart_butterfly",
// // 			Tags:     []string{"small"}})
// // 		if err != nil {
// // 			fmt.Println("error")
// // 		} else {
// // 			// Log the new tag to the console.
// // 			fmt.Println("New tag: ", update_resp.Tags, "\n")
// // 		}
// // 	}

// // }

// // func transformImage(cld *cloudinary.Cloudinary, ctx context.Context) {
// // 	// Instantiate an object for the asset with public ID "my_image"
// // 	qs_img, err := cld.Image("quickstart_butterfly")
// // 	if err != nil {
// // 		fmt.Println("error")
// // 	}

// // 	// Add the transformation
// // 	qs_img.Transformation = "r_max/e_sepia"

// // 	// Generate and log the delivery URL
// // 	new_url, err := qs_img.String()
// // 	if err != nil {
// // 		fmt.Println("error")
// // 	} else {
// // 		print("****4. Transform the image****\nTransfrmation URL: ", new_url, "\n")
// // 	}
// // }

// func downloadImage(url string, filePath string) error {
// 	// Membuka HTTP request untuk mengunduh gambar
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Membuat file lokal untuk menyimpan gambar
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Menulis konten dari respon HTTP ke file lokal
// 	_, err = io.Copy(file, resp.Body)
// 	return err
// }

// func main() {
// 	cld, ctx := credentials()
// 	//uploadImage(cld, ctx)

// 	imageURL, err := getImageInfo(cld, ctx, "quickstart_butterfly")

// 	fmt.Println(imageURL)
// 	//getAssetInfo(cld, ctx)
// 	// transformImage(cld, ctx)

// 	if err != nil {
// 		fmt.Println("error getting image info:", err)
// 		return
// 	}

// 	// Menyimpan gambar ke local storage
// 	filePath := "butterfly.png"
// 	err = downloadImage(imageURL, filePath)
// 	if err != nil {
// 		fmt.Println("error downloading image:", err)
// 		return
// 	}

// 	fmt.Println("Image downloaded successfully to", filePath)
// }
