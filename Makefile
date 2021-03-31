OLD_RUNNER_VERSION:=v0.0.4
NEW_RUNNER_VERSION:=v0.0.5

comby-dry-run-bump-version:
	comby -d lang $(OLD_RUNNER_VERSION) $(NEW_RUNNER_VERSION)

comby-in-place-bump-version:
	comby -d lang -in-place $(OLD_RUNNER_VERSION) $(NEW_RUNNER_VERSION)
