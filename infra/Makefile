push_image:
	docker pull artdecoction.registry.jetbrains.space/p/wt/services/$(img):$(tag)
	docker tag artdecoction.registry.jetbrains.space/p/wt/services/$(img):$(tag) \
		us-central1-docker.pkg.dev/ad-tower/services/$(img):$(tag)
	docker push us-central1-docker.pkg.dev/ad-tower/services/$(img):$(tag)
	docker rmi \
		$$( \
			docker images 'us-central1-docker.pkg.dev/ad-tower/services/$(img)' \
			--format "{{.Repository}}:{{.Tag}}" \
		) \
		|| [ $$? -eq 1 ]
	docker rmi \
		$$( \
			docker images 'artdecoction.registry.jetbrains.space/p/wt/services/$(img)' \
			--format "{{.Repository}}:{{.Tag}}" \
	 		| sed 1,6d \
		) \
		|| [ $$? -eq 1 ]
	echo "Done"
