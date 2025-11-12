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
    git tag --list --merged origin/release "$pattern" | grep -E '(\+kong-|-kong-)' | grep -v '/' | sort -V | tail -n 1
  else
    pattern="${module}/v[0-9]*"
    git tag --list --merged origin/release "$pattern" | grep -E '(\+kong-|-kong-)' | sort -V | tail -n 1
  fi
}

# Function to get upstream tags for a module
get_upstream_tags() {
  local module="$1"
  local pattern

  if [[ "$module" == "root" ]]; then
    pattern="v[0-9]*"
    git tag --list --merged origin/main "$pattern" | grep -v '\-kong-' | grep -v '+kong-' | grep -v '/' | sort -V
  else
    pattern="${module}/v[0-9]*"
    git tag --list --merged origin/main "$pattern" | grep -v '\-kong-' | grep -v '+kong-' | sort -V
  fi
}

# Function to strip Kong suffix from version tag
strip_kong_suffix() {
  local tag="$1"
  tag="${tag%+kong-*}"
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

# Get root module versions for rebase decision
# Note: root is the first module in MODULES array (index 0)
current_root_tag="${latest_versions[0]}"
main_root_tag=$(git tag --list --merged origin/main 'v[0-9].[0-9]*.[0-9]*' | grep -v '\-kong-' | grep -v '+kong-' | grep -v '/' | sort -V | tail -n 1)

echo ""
echo "Current release root tag: ${current_root_tag:-none}"
echo "Upstream root tag: $main_root_tag"

# Determine if we need to rebase (root version changed)
needs_rebase=false
if [[ -z "$current_root_tag" ]]; then
  needs_rebase=true
  echo "No existing release - will need to rebase"
elif [[ "$main_root_tag" != "$current_root_tag" ]]; then
  needs_rebase=true
  echo "Root version changed - will rebase preserving custom commits"
else
  echo "Root version unchanged - will only add new tags"
fi

# Build list of new Kong tags
new_tags=()
for tag in "${new_upstream_tags[@]}"; do
  new_tags+=("${tag}+kong-1")
done

echo ""
echo "New Kong tags to be created:"
for tag in "${new_tags[@]}"; do
  echo "  - $tag"
done

# Output for GitHub Actions
if [[ -n "${GITHUB_OUTPUT:-}" ]]; then
  echo "needs_release=true" >> "$GITHUB_OUTPUT"
  echo "needs_rebase=$needs_rebase" >> "$GITHUB_OUTPUT"
  echo "released_tag=$current_root_tag" >> "$GITHUB_OUTPUT"
  echo "main_tag=$main_root_tag" >> "$GITHUB_OUTPUT"

  # Export all new tags as a JSON array (compact format for GitHub Actions)
  new_tags_json=$(printf '%s\n' "${new_tags[@]}" | jq -R . | jq -sc .)
  echo "new_tags=$new_tags_json" >> "$GITHUB_OUTPUT"
fi

echo ""
echo "=== Detection complete ==="
