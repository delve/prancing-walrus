.PHONY: image debug testimage runimage publish

image:
	cd walrusbot && pack build prancing-walrus --builder=gcr.io/buildpacks/builder 

run:
	cd walrusbot && go run .

runimage:
	docker run -e APIKey=${APIKey} -e BOT_TOKEN=${BOT_TOKEN} -e CONFIG=${CONFIG} prancing-walrus


publish: gcloudinit image
	docker tag prancing-walrus us-central1-docker.pkg.dev/prancingwalrus/prancing-walrus/prancing-walrus:v0.01
	docker push us-central1-docker.pkg.dev/prancingwalrus/prancing-walrus/prancing-walrus:v0.01

gcloudinit: /tmp/gcloud.inited
/tmp/gcloud.inited:
	gcloud auth login
	docker login us-central1-docker.pkg.dev
	touch /tmp/gcloud.inited

clean:
	rm /tmp/gcloud.inited