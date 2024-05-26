containerVersion = v0.21
debugVersion = test

MODELS_DIR := ./walrusbot/sheetDAO
MODELS_GO := $(MODELS_DIR)/models.go $(MODELS_DIR)/enums.go
MODEL_FILES := $(MODELS_DIR)/model_snail.go $(MODELS_DIR)/model_snailstat.go

.PHONY: image debug runImage publish

run: /tmp/gcloud.inited $(MODEL_FILES)
	cd walrusbot && go run .

image: $(MODEL_FILES) /tmp/image.built
/tmp/image.built:
	cd walrusbot && pack build prancing-walrus --builder=gcr.io/buildpacks/builder
	touch /tmp/image.built

inspectImage: image
	pack inspect-image prancing-walrus

exploreImage: image
	docker run --rm --entrypoint launcher -it prancing-walrus bash

runImage: image
	docker run -e APIKey=${APIKey} -e BOT_TOKEN=${BOT_TOKEN} -e CONFIG=${CONFIG} prancing-walrus

publish: gcloudinit image
	docker tag prancing-walrus us-central1-docker.pkg.dev/prancingwalrus/prancing-walrus/prancing-walrus:$(containerVersion)
	docker push us-central1-docker.pkg.dev/prancingwalrus/prancing-walrus/prancing-walrus:$(containerVersion)

publishDebug: gcloudinit image
	docker tag prancing-walrus us-central1-docker.pkg.dev/prancingwalrus/prancing-walrus/prancing-walrus:$(debugVersion)
	docker push us-central1-docker.pkg.dev/prancingwalrus/prancing-walrus/prancing-walrus:$(debugVersion)

# see https://cloud.google.com/docs/authentication/provide-credentials-adc#sa-impersonation regarding service account creds
gcloudinit: /tmp/gcloud.inited
/tmp/gcloud.inited:
	gcloud auth login
	gcloud auth application-default login --impersonate-service-account walrus-sheet-access@prancingwalrus.iam.gserviceaccount.com
	docker login us-central1-docker.pkg.dev
	touch /tmp/gcloud.inited

cleanAll: cleanGcloud clean cleandbImportInit

cleanGcloud:
	rm /tmp/gcloud.inited

cleandbImportInit:
	rm /tmp/sa_private_key.pem

clean:
	rm /tmp/image.built

dbImportInit: 
	rm -rf ./importWalrusDb/sheetDAO && cp -r ./walrusbot/sheetDAO ./importWalrusDb
	rm -rf ./importWalrusDb/utility && cp -r ./walrusbot/utility ./importWalrusDb
	rm -f ./importWalrusDb/devconfig.json && cp ./walrusbot/devconfig.json ./importWalrusDb/devconfig.json
	rm -f ./importWalrusDb/config.json && cp ./walrusbot/config.json ./importWalrusDb/config.json

importDb: gcloudinit dbImportInit $(MODEL_FILES)
	cd importWalrusDb && go run .

$(MODEL_FILES): $(MODELS_GO)
	@echo "Regenerating $@ from $<"
	cd walrusbot && go generate ./sheetDAO

$(MODELS_GO): /workspace/go/bin/sheetdb-modeler

/workspace/go/bin/sheetdb-modeler: # or @v0.2.0
	go install github.com/delve/sheetdb/tools/sheetdb-modeler@latest

updateSheetdbVersion:
	echo OH SHIT. You better write the code to do this.
	echo update in all go.mod files, go.work file, and gitpodconfig command.sh file
updateGsheetsVersion:
	echo OH SHIT. You better write the code to do this.
	echo update in all go.mod files, go.work file, and gitpod.yml file
