# make image: pack build --builder=gcr.io/buildpacks/builder prancing-walrus
# run image: docker run -e APIKey=${APIKey} -e BOT_TOKEN=${BOT_TOKEN} -e CONFIG=${CONFIG} prancing-walrus
# debug code: cd walrusbot;go run .
# interactive debug: F5

#https://stackoverflow.com/questions/69447497/is-it-possible-to-customize-docker-image-generated-with-spring-native-with-buil
# publish image: pack build --builder=gcr.io/buildpacks/builder prancing-walrus --publish <repo details here>
