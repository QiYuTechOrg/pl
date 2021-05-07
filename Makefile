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

