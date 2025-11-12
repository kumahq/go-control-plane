#!/bin/bash

set -e

# Get the root module tag (main version) from each branch
current_root_tag=$(git tag --list 'v[0-9].[0-9]*.[0-9]*' --merged origin/release --sort=-creatordate | grep -v '+kong-' | head -n 1)
main_root_tag=$(git tag --list 'v[0-9].[0-9]*.[0-9]*' --merged origin/main --sort=-creatordate | head -n 1)

# Remove the prefix `v` and any suffix for comparison
released_tag="${current_root_tag%+kong-*}"
released_tag="${released_tag%-kong-*}"
currentVersion="${released_tag#v}"

# The main branch tag
newVersion="${main_root_tag#v}"

echo "Current release root tag: $released_tag"
echo "Upstream root tag: $main_root_tag"

# Get all tags pointing to the same commit as the main root tag
main_commit=$(git rev-list -n 1 "refs/tags/${main_root_tag}")
echo "Upstream commit: $main_commit"

# Get all upstream tags for this commit (excluding any kong-tagged versions)
upstream_tags=()
while IFS= read -r tag; do
  upstream_tags+=("$tag")
done < <(git tag --points-at "$main_commit" | grep -v '\-kong-' | grep -v '+kong-' | sort)

echo "Found ${#upstream_tags[@]} upstream tags for this release:"
for tag in "${upstream_tags[@]}"; do
  echo "  - $tag"
done

# Check if we need to create a new release
new_root_tag="${main_root_tag}+kong-1"

# Get the current release commit's tags to find the highest kong version
if [[ -n "$current_root_tag" ]]; then
  release_commit=$(git rev-list -n 1 "refs/tags/${current_root_tag}" 2>/dev/null || echo "")
  if [[ -n "$release_commit" ]]; then
    # Find the highest kong version for the current release
    highest_kong_version=0
    while IFS= read -r tag; do
      if [[ $tag =~ \+kong-([0-9]+)$ ]]; then
        kong_num="${BASH_REMATCH[1]}"
        if [[ $kong_num -gt $highest_kong_version ]]; then
          highest_kong_version=$kong_num
        fi
      fi
    done < <(git tag --points-at "$release_commit" | grep '+kong-')

    echo "Current highest kong version: $highest_kong_version"
  fi
fi

# Compare versions
IFS='.' read -r -a currentParts <<< "$currentVersion"
IFS='.' read -r -a newParts <<< "$newVersion"

should_release=false

# Compare each part
for i in 0 1 2; do
  if [[ ${newParts[i]:-0} -gt ${currentParts[i]:-0} ]]; then
    echo "The upstream version is higher - new release needed."
    should_release=true
    break
  elif [[ ${newParts[i]:-0} -lt ${currentParts[i]:-0} ]]; then
    echo "The current tag is higher. That shouldn't be the case, please fix tagging."
    exit 1
  fi
done

if [[ "$should_release" == "false" ]]; then
  echo "Tags are equal, no need to release"
  exit 0
fi

# Build the list of new tags with +kong-1 suffix
new_tags=()
for tag in "${upstream_tags[@]}"; do
  new_tags+=("${tag}+kong-1")
done

echo ""
echo "New tags to be created:"
for tag in "${new_tags[@]}"; do
  echo "  - $tag"
done

# Output for GitHub Actions
echo "released_tag=$released_tag" >> $GITHUB_OUTPUT
echo "main_tag=$main_root_tag" >> $GITHUB_OUTPUT
echo "main_commit=$main_commit" >> $GITHUB_OUTPUT

# Export all new tags as a JSON array for easy parsing in the workflow
new_tags_json=$(printf '%s\n' "${new_tags[@]}" | jq -R . | jq -s .)
echo "new_tags=$new_tags_json" >> $GITHUB_OUTPUT

# Also export the root tag for backward compatibility
echo "new_tag=$new_root_tag" >> $GITHUB_OUTPUT
