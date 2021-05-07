OLD_RUNNER_VERSION:=v0.1.0
NEW_RUNNER_VERSION:=v0.1.1

comby-dry-run-runner-version:
	comby -d lang $(OLD_RUNNER_VERSION) $(NEW_RUNNER_VERSION)

comby-in-place-runner-version:
	comby -d lang -in-place $(OLD_RUNNER_VERSION) $(NEW_RUNNER_VERSION)


OLD_PATCH_VERSION:=p7
NEW_PATCH_VERSION:=p8

comby-dry-run-bump-patch:
	comby -d .github/workflows $(OLD_PATCH_VERSION) $(NEW_PATCH_VERSION)

comby-in-place-bump-patch:
	comby -d .github/workflows -in-place $(OLD_PATCH_VERSION) $(NEW_PATCH_VERSION)



local-test:
	docker build -f docker_test/php/Dockerfile . -t php
	cd docker_test/php && docker run -it -v `pwd`/data:/data php runner php

	docker build -f docker_test/node/Dockerfile . -t node
	cd docker_test/node && docker run -it -v `pwd`/data:/data node runner node

	docker build -f docker_test/python/Dockerfile . -t python
	cd docker_test/python && docker run -it -v `pwd`/data:/data python runner python

clean-local-test:
	rm -f docker_test/*/data/out.json
