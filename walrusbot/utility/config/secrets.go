package config

import (
	"context"
	"encoding/base64"
	"fmt"
	"walrusbot/utility/check"
	"walrusbot/utility/log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"golang.org/x/oauth2"
	"google.golang.org/api/iam/v1"
)

type secrets struct {
	botToken string
	// TODO: not needed?/broken?
	sheetsApiKey    string
	sheetsOauthCert string
	// sheetsOauthToken  *oauth2.Token
	serviceAccountKey *iam.ServiceAccountKey
}

func (s secrets) GetBotToken() string {
	return s.botToken
}

// TODO: Remove if not in use for sheetsdb
func (s secrets) GetSheetsApiKey() string {
	return s.sheetsApiKey
}

// TODO: Remove if not in use for sheetsdb
func (s secrets) GetServiceAccountKeyRaw() *iam.ServiceAccountKey {
	// jsonKeyFile, _ := base64.StdEncoding.DecodeString(key.PrivateKeyData)

	return s.serviceAccountKey
}

// TODO: Remove if not in use for sheetsdb
func (s secrets) GetServiceAccountKey() []byte {
	jsonKeyFile, _ := base64.StdEncoding.DecodeString(s.serviceAccountKey.PrivateKeyData)

	return jsonKeyFile
}

func (s secrets) GetSheetsOauthCert() string {
	return s.sheetsOauthCert
}

// func (s secrets) GetSheetsOauthToken() *oauth2.Token {
// 	return s.sheetsOauthToken
// }

func get_secrets() (err error) {
	//TODO: this is trash, make it nicer
	Values.Secrets = secrets{}
	var secretValue, secretName string
	var secretBytes []byte
	secretValue = ""
	secretName = Values.SMgrSecretList["botToken"]
	log.Infow("retrieving secret from secretmanager", "secretName", secretName)
	secretBytes, err = accessSecretVersion(secretName)
	if err != nil {
		return fmt.Errorf("failed to retrieve secret from secretmanager: %w", err)
	}
	if len(secretBytes) == 0 {
		return fmt.Errorf("retrieved 0 byte secret from secretmanager: %s", secretName)
	}
	secretValue = string(secretBytes)
	Values.Secrets.botToken = secretValue

	secretValue = ""
	secretBytes = []byte{} // deliberately blank the temp value
	secretName = "SheetsAPIKey"
	log.Infow("retrieving secret from secretmanager", "secretName", secretName)
	secretBytes, err = accessSecretVersion(secretName)
	if err != nil {
		return fmt.Errorf("failed to retrieve secret from secretmanager: %w", err)
	}
	if len(secretBytes) == 0 {
		return fmt.Errorf("retrieved 0 byte secret from secretmanager: %s", secretName)
	}
	secretValue = string(secretBytes)
	Values.Secrets.sheetsApiKey = secretValue

	secretValue = ""
	secretBytes = []byte{} // deliberately blank the temp value
	secretName = "SheetsOauth"
	log.Infow("retrieving secret from secretmanager", "secretName", secretName)
	secretBytes, err = accessSecretVersion(secretName)
	if err != nil {
		return fmt.Errorf("failed to retrieve secret from secretmanager: %w", err)
	}
	if len(secretBytes) == 0 {
		return fmt.Errorf("retrieved 0 byte secret from secretmanager: %s", secretName)
	}
	secretValue = string(secretBytes)
	Values.Secrets.sheetsOauthCert = secretValue

	// tidy up any loose SA keys
	keyList, err := listSAKeys()
	check.Err(err)
	for _, key := range keyList {
		deleteSAKey(key)
	}

	Values.Secrets.serviceAccountKey, err = createSAKey()
	check.Err(err)
	return err
}

// accessSecretVersion retrieves the latest version for the named secret from GCP secretmanager
func accessSecretVersion(secretName string) ([]byte, error) {

	version := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", Values.GcpProject, secretName)
	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	check.Err(err, "failed to create secretmanager client")
	defer client.Close()

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: version,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	check.Err(err, "failed to access secret version")

	log.Infow("retrieved payload for secret", "secretName", result.Name, "secretLength", len(result.Payload.Data))
	return result.Payload.Data, nil
}

// listSAKeys returns a slice of SA Key IDs suitable for deleting
func listSAKeys() ([]string, error) {
	service, err := getIamSvc()
	check.Err(err, "could not create IAM service in listSAKeys()")
	keyNames := []string{}
	resource := "projects/" + Values.GcpProject + "/serviceAccounts/" + Values.GcpSAName

	keyList, err := service.Projects.ServiceAccounts.Keys.List(resource).Do()
	if err != nil {
		return nil, fmt.Errorf("Projects.ServiceAccounts.Keys.List: %w", err)
	}

	for _, key := range keyList.Keys {
		if key.KeyType == "USER_MANAGED" {
			keyNames = append(keyNames, key.Name)
		}
	}
	return keyNames, nil
}

// createSAKey creates a service account key.
func createSAKey() (*iam.ServiceAccountKey, error) {
	service, err := getIamSvc()
	check.Err(err, "could not create IAM service in createSAKey()")

	resource := "projects/" + Values.GcpProject + "/serviceAccounts/" + Values.GcpSAName
	request := &iam.CreateServiceAccountKeyRequest{}
	key, err := service.Projects.ServiceAccounts.Keys.Create(resource, request).Do()
	if err != nil {
		return nil, fmt.Errorf("Projects.ServiceAccounts.Keys.Create: %w", err)
	}

	log.Infow("key created successfully", "name", key.Name)
	return key, nil
}

// deleteSAKey deletes a service account key.
func deleteSAKey(fullKeyName string) error {
	service, err := getIamSvc()
	check.Err(err, "could not create IAM service in deleteSAKey()")

	_, err = service.Projects.ServiceAccounts.Keys.Delete(fullKeyName).Do()
	check.Err(err, "Projects.ServiceAccounts.Keys.Delete failed")

	log.Infow("deleted SA key", "name", fullKeyName)
	return nil
}

// getIamSvc returns a GCP IAM service client
func getIamSvc() (*iam.Service, error) {
	ctx := context.Background()
	service, err := iam.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("iam.NewService: %w", err)
	}

	return service, nil
}

// OAUTH2 mess
// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalw("unable to read authorization code", "error", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalw("unable to retrieve token from web", "error", err)
	}
	return tok
}
