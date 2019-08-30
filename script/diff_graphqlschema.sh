#!/bin/bash

branch2="${CI_COMMIT_REF_NAME:-notset}"
case "$branch2" in
	notset)
		echo "Need env var CI_COMMIT_REF_NAME"
		exit 1
		;;
	develop)
		branch1="" # don't diff against master, it gets noisy on Slack#uidev.
		;;
	master)
		branch1=""
		;;
	release/v*)
		branch1=master
		;;
	*)
		branch1=develop
		;;
esac

code=0
if test -n "$branch1" ; then
	token="${GITLAB_API_TOKEN:-}"
	if test -z "$token" ; then
		echo "Need env var GITLAB_API_TOKEN"
		exit 1
	fi
	pushd "$(realpath "$(dirname "$0")/..")" 1>/dev/null || exit 1
	bash script/diff_file.sh \
		"https://gitlab.com/api/v4" \
		"$token" \
		5726034 \
		"$branch1" \
		"$branch2" \
		"internal/gateway/graphql/schema.graphql" \
		|| code=1
	popd 1>/dev/null || exit 1
fi
exit "$code"