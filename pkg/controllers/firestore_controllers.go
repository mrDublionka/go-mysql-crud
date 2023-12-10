package controllers

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go"
	"github.com/mrDublionka/go-mysql-crud/pkg/models"
	"google.golang.org/api/option"
	"io"
	"log"
	"net/http"
)

type ImageControllers struct {
	ctx     context.Context
	storage *storage.Client
	client  *firestore.Client
}

func NewImageController() *ImageControllers {
	ic := &ImageControllers{}
	ic.InitFirebaseStorage()
	return ic
}

func (ic *ImageControllers) InitFirebaseStorage() {
	ic.ctx = context.Background()

	sa := option.WithCredentialsFile("serviceAccountKey.json")

	var err error
	ic.storage, err = storage.NewClient(ic.ctx, sa)

	if err != nil {
		log.Fatalf("Error initializing firebase storage client: %v", err)
		return
	}

	app, err := firebase.NewApp(ic.ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	ic.client, err = app.Firestore(ic.ctx)
	if err != nil {
		log.Fatalln(err)
	}

	println("Storage init success!")
}

func (ic *ImageControllers) UploadImage(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("image")
	r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer file.Close()

	bucketName := "blog-next-php.appspot.com"
	objectName := handler.Filename

	wc := ic.storage.Bucket(bucketName).Object(objectName).NewWriter(ic.ctx)
	_, err = io.Copy(wc, file)

	if err != nil {
		return
	}

	if err := wc.Close(); err != nil {
		return
	}

	err = CreateImageUrl(objectName, bucketName, ic.ctx, ic.client)
	if _, err := io.Copy(wc, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func CreateImageUrl(imagePath string, bucket string, ctx context.Context, client *firestore.Client) error {
	imageStructure := models.ImageStructure{
		ImageName: imagePath,
		URL:       "https://storage.cloud.google.com/" + bucket + "/" + imagePath,
	}

	_, _, err := client.Collection("image").Add(ctx, imageStructure)
	if err != nil {
		return err
	}

	return nil
}
