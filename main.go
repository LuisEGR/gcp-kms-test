// Sample quickstart is a basic program that uses Cloud KMS.
package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"log"

	kms "cloud.google.com/go/kms/apiv1"
	"google.golang.org/api/iterator"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

func main() {
	// GCP project with which to communicate.
	projectID := "svc-project-1fae"

	// Location in which to list key rings.
	locationID := "us-central1"

	// Create the client.
	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}
	defer client.Close()

	// Create the request to list KeyRings.
	listKeyRingsReq := &kmspb.ListKeyRingsRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s", projectID, locationID),
	}

	// List the KeyRings.
	it := client.ListKeyRings(ctx, listKeyRingsReq)

	// Iterate and print the results.
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to list key rings: %v", err)
		}

		fmt.Printf("key ring: %s\n", resp.Name)
	}

	var buf bytes.Buffer
	// fmt.Fprintf(&buf, "Size: %d MB.", 85)
	keyname := "projects/svc-project-1fae/locations/us-central1/keyRings/my-keyring/cryptoKeys/my-generated-key"

	err, encrypted := EncryptSymmetric(&buf, keyname, "hello world")
	fmt.Println("err :", err)
	fmt.Println("encrypted :", encrypted)
	// fmt.Println("buf :", buf)
	// s := buf.
	// fmt.Println("s :", s)

	encryptedEncoded := hex.EncodeToString(encrypted)
	fmt.Println("encryptedEncoded :", encryptedEncoded)

	var buf2 bytes.Buffer
	DecryptSymmetric(&buf2, keyname, encrypted)
	fmt.Println("buf2 :", buf2.String())

	// var buf3 bytes.Buffer
	// url, err := GenerateV4GetObjectSignedURL(&buf3, "test-bucket-mediahub", "fusion708.mov")
	// fmt.Println("err :", err)
	// fmt.Println("url :", url)

	serviceAccount := `{ example service account }`

	var bufSAEnc bytes.Buffer

	err, encriptedSA := EncryptSymmetric(&bufSAEnc, keyname, serviceAccount)
	fmt.Println("err :", err)
	// fmt.Println("bufSAEnc :", bufSAEnc)
	// fmt.Println("encriptedSA :", encriptedSA)
	encryptedEncodedSA := hex.EncodeToString(encriptedSA)
	fmt.Println("encryptedEncodedSA :", encryptedEncodedSA)
}

// Keyring: my-keyring
// key region: us-central1
// key name: my-generated-key
