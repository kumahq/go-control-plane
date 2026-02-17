#!/bin/bash

set -e

echo "=== Detecting new upstream tags ==="

# Define all modules to process
MODULES=("root" "envoy" "contrib" "ratelimit" "xdsmatcher")

# Function to get latest Kong tag for a module
get_latest_kong_tag() {
  local module="$1"
  local pattern

  if [[ "$module" == "root" ]]; then
    pattern="v[0-9]*"
    git tag --list "$pattern" | grep -E -- '[-+]kong-' | grep -v '/' | sort -V | tail -n 1
  else
    pattern="${module}/v[0-9]*"
    git tag --list "$pattern" | grep -E -- '[-+]kong-' | sort -V | tail -n 1
  fi
}

# Function to get upstream tags for a module
get_upstream_tags() {
  local module="$1"
  local pattern

  if [[ "$module" == "root" ]]; then
    pattern="v[0-9]*"
    git tag --list --merged origin/main "$pattern" | grep -Ev -- '[-+]kong-' | grep -v '/' | sort -V
  else
    pattern="${module}/v[0-9]*"
    git tag --list --merged origin/main "$pattern" | grep -Ev -- '[-+]kong-' | sort -V
  fi
}

# Function to strip Kong suffix from version tag
strip_kong_suffix() {
  local tag="$1"
  tag="${tag%-kong-*}"
  echo "$tag"
}

# Function to process new upstream tags for a module
process_module_tags() {
  local module="$1"
  local latest_version="$2"

  if [[ -n "$latest_version" ]]; then
    # Find tags newer than the latest Kong version
    local found_latest=false
    while IFS= read -r tag; do
      if [[ "$found_latest" == "true" ]]; then
        new_upstream_tags+=("$tag")
        echo "  New tag in $module: $tag"
      elif [[ "$tag" == "$latest_version" ]]; then
        found_latest=true
      fi
    done < <(get_upstream_tags "$module")
  else
    # No existing Kong tag, get the latest upstream tag
    local latest_upstream=$(get_upstream_tags "$module" | tail -n 1)
    if [[ -n "$latest_upstream" ]]; then
      new_upstream_tags+=("$latest_upstream")
      echo "  First tag for $module: $latest_upstream"
    fi
  fi
}

# Find latest Kong tags per module
echo "Finding latest Kong tags per module..."
latest_versions=()
for module in "${MODULES[@]}"; do
  kong_tag=$(get_latest_kong_tag "$module")
  if [[ -n "$kong_tag" ]]; then
    stripped_version=$(strip_kong_suffix "$kong_tag")
    latest_versions+=("$stripped_version")
    echo "  $module: $stripped_version"
  else
    latest_versions+=("")
    echo "  $module: none"
  fi
done

# Find new upstream tags for each module
new_upstream_tags=()

echo ""
echo "Checking for new upstream tags..."

for i in "${!MODULES[@]}"; do
  process_module_tags "${MODULES[$i]}" "${latest_versions[$i]}"
done

if [[ ${#new_upstream_tags[@]} -eq 0 ]]; then
  echo ""
  echo "No new upstream tags found. No action needed."
  if [[ -n "${GITHUB_OUTPUT:-}" ]]; then
    echo "needs_release=false" >> "$GITHUB_OUTPUT"
  fi
  exit 0
fi

echo ""
echo "Found ${#new_upstream_tags[@]} new upstream tags"

# Find the rebase target: the latest commit among all new upstream tags
# This ensures release includes upstream commits up to the newest tag, not beyond
echo ""
echo "Finding rebase target from new upstream tags..."
rebase_target=""
for tag in "${new_upstream_tags[@]}"; do
  tag_commit=$(git rev-parse "$tag")
  if [[ -z "$rebase_target" ]]; then
    rebase_target="$tag_commit"
  else
    # Pick the commit that is ahead (closer to main HEAD)
    if git merge-base --is-ancestor "$rebase_target" "$tag_commit"; then
      rebase_target="$tag_commit"
    fi
  fi
done
echo "Rebase target: $(git rev-parse --short "$rebase_target") (from tag: $(git describe --tags --exact-match "$rebase_target" 2>/dev/null || echo 'multiple tags'))"

# Find the base commit where release diverges from main (last shared upstream commit)
base_commit=$(git merge-base origin/main origin/release)
echo "Base commit: $(git rev-parse --short "$base_commit")"

custom_commits=$(git rev-list "${base_commit}..origin/release" --count)
echo "Found $custom_commits custom Kong commit(s) on release branch"

# Build list of new Kong tags
new_tags=()
for tag in "${new_upstream_tags[@]}"; do
  new_tags+=("${tag}-kong-1")
done

echo ""
echo "New Kong tags to be created:"
for tag in "${new_tags[@]}"; do
  echo "  - $tag"
done

# Output for GitHub Actions
if [[ -n "${GITHUB_OUTPUT:-}" ]]; then
  echo "needs_release=true" >> "$GITHUB_OUTPUT"
  echo "rebase_target=$rebase_target" >> "$GITHUB_OUTPUT"
  echo "base_commit=$base_commit" >> "$GITHUB_OUTPUT"

  # Export all new tags as a JSON array (compact format for GitHub Actions)
  new_tags_json=$(printf '%s\n' "${new_tags[@]}" | jq -R . | jq -sc .)
  echo "new_tags=$new_tags_json" >> "$GITHUB_OUTPUT"
fi

echo ""
echo "=== Detection complete ==="
