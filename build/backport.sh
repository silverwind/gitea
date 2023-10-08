#!/bin/bash
set -euo pipefail

backport() {
  if [ $# -eq 0 ]; then; echo "backport PR VER"; exit 1; fi
  PR="$1"
  VER="$2"

  git switch "release/v$VER"
  git branch -D "backport-$PR-$VER"
  git checkout -b "backport-$PR-$VER"
  HASH="$(curl -sH "X-GitHub-Api-Version: 2022-11-28" "https://api.github.com/repos/go-gitea/gitea/pulls/$PR" | jq -r .merge_commit_sha)"
  git cherry-pick "$HASH"
}
