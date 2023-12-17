.PHONY: image debug testimage runimage publish

image:
	cd walrusbot && pack build --builder=gcr.io/buildpacks/builder prancing-walrus

run:
	cd walrusbot && go run .

runimage:
	docker run -e APIKey=${APIKey} -e BOT_TOKEN=${BOT_TOKEN} -e CONFIG=${CONFIG} prancing-walrus

publish:
	pack build --builder=gcr.io/buildpacks/builder prancing-walrus --publish <repo details here>