#!/bin/bash

set -e

echo "=== Detecting new upstream tags ==="

# Function to get latest Kong tag for a module
get_latest_kong_tag() {
  local module="$1"
  local pattern

  if [[ "$module" == "root" ]]; then
    pattern="v[0-9]*"
    git tag --list --merged origin/test-rel "$pattern" | grep -E '(\+kong-|-kong-)' | grep -v '/' | sort -V | tail -n 1
  else
    pattern="${module}/v[0-9]*"
    git tag --list --merged origin/test-rel "$pattern" | grep -E '(\+kong-|-kong-)' | sort -V | tail -n 1
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

# Find latest Kong tags per module
echo "Finding latest Kong tags per module..."
latest_root=$(get_latest_kong_tag "root")
latest_envoy=$(get_latest_kong_tag "envoy")
latest_contrib=$(get_latest_kong_tag "contrib")
latest_ratelimit=$(get_latest_kong_tag "ratelimit")
latest_xdsmatcher=$(get_latest_kong_tag "xdsmatcher")

# Strip Kong suffix to get base versions
if [[ -n "$latest_root" ]]; then
  latest_root="${latest_root%+kong-*}"
  latest_root="${latest_root%-kong-*}"
  echo "  root: $latest_root"
else
  echo "  root: none"
fi

if [[ -n "$latest_envoy" ]]; then
  latest_envoy="${latest_envoy%+kong-*}"
  latest_envoy="${latest_envoy%-kong-*}"
  echo "  envoy: $latest_envoy"
else
  echo "  envoy: none"
fi

if [[ -n "$latest_contrib" ]]; then
  latest_contrib="${latest_contrib%+kong-*}"
  latest_contrib="${latest_contrib%-kong-*}"
  echo "  contrib: $latest_contrib"
else
  echo "  contrib: none"
fi

if [[ -n "$latest_ratelimit" ]]; then
  latest_ratelimit="${latest_ratelimit%+kong-*}"
  latest_ratelimit="${latest_ratelimit%-kong-*}"
  echo "  ratelimit: $latest_ratelimit"
else
  echo "  ratelimit: none"
fi

if [[ -n "$latest_xdsmatcher" ]]; then
  latest_xdsmatcher="${latest_xdsmatcher%+kong-*}"
  latest_xdsmatcher="${latest_xdsmatcher%-kong-*}"
  echo "  xdsmatcher: $latest_xdsmatcher"
else
  echo "  xdsmatcher: none"
fi

# Find new upstream tags for each module
new_upstream_tags=()

echo ""
echo "Checking for new upstream tags..."

# Process root module
if [[ -n "$latest_root" ]]; then
  found_latest=false
  while IFS= read -r tag; do
    if [[ "$found_latest" == "true" ]]; then
      new_upstream_tags+=("$tag")
      echo "  New tag in root: $tag"
    elif [[ "$tag" == "$latest_root" ]]; then
      found_latest=true
    fi
  done < <(get_upstream_tags "root")
else
  latest_upstream=$(get_upstream_tags "root" | tail -n 1)
  if [[ -n "$latest_upstream" ]]; then
    new_upstream_tags+=("$latest_upstream")
    echo "  First tag for root: $latest_upstream"
  fi
fi

# Process envoy module
if [[ -n "$latest_envoy" ]]; then
  found_latest=false
  while IFS= read -r tag; do
    if [[ "$found_latest" == "true" ]]; then
      new_upstream_tags+=("$tag")
      echo "  New tag in envoy: $tag"
    elif [[ "$tag" == "$latest_envoy" ]]; then
      found_latest=true
    fi
  done < <(get_upstream_tags "envoy")
else
  latest_upstream=$(get_upstream_tags "envoy" | tail -n 1)
  if [[ -n "$latest_upstream" ]]; then
    new_upstream_tags+=("$latest_upstream")
    echo "  First tag for envoy: $latest_upstream"
  fi
fi

# Process contrib module
if [[ -n "$latest_contrib" ]]; then
  found_latest=false
  while IFS= read -r tag; do
    if [[ "$found_latest" == "true" ]]; then
      new_upstream_tags+=("$tag")
      echo "  New tag in contrib: $tag"
    elif [[ "$tag" == "$latest_contrib" ]]; then
      found_latest=true
    fi
  done < <(get_upstream_tags "contrib")
else
  latest_upstream=$(get_upstream_tags "contrib" | tail -n 1)
  if [[ -n "$latest_upstream" ]]; then
    new_upstream_tags+=("$latest_upstream")
    echo "  First tag for contrib: $latest_upstream"
  fi
fi

# Process ratelimit module
if [[ -n "$latest_ratelimit" ]]; then
  found_latest=false
  while IFS= read -r tag; do
    if [[ "$found_latest" == "true" ]]; then
      new_upstream_tags+=("$tag")
      echo "  New tag in ratelimit: $tag"
    elif [[ "$tag" == "$latest_ratelimit" ]]; then
      found_latest=true
    fi
  done < <(get_upstream_tags "ratelimit")
else
  latest_upstream=$(get_upstream_tags "ratelimit" | tail -n 1)
  if [[ -n "$latest_upstream" ]]; then
    new_upstream_tags+=("$latest_upstream")
    echo "  First tag for ratelimit: $latest_upstream"
  fi
fi

# Process xdsmatcher module
if [[ -n "$latest_xdsmatcher" ]]; then
  found_latest=false
  while IFS= read -r tag; do
    if [[ "$found_latest" == "true" ]]; then
      new_upstream_tags+=("$tag")
      echo "  New tag in xdsmatcher: $tag"
    elif [[ "$tag" == "$latest_xdsmatcher" ]]; then
      found_latest=true
    fi
  done < <(get_upstream_tags "xdsmatcher")
else
  latest_upstream=$(get_upstream_tags "xdsmatcher" | tail -n 1)
  if [[ -n "$latest_upstream" ]]; then
    new_upstream_tags+=("$latest_upstream")
    echo "  First tag for xdsmatcher: $latest_upstream"
  fi
fi

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
current_root_tag="$latest_root"
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

  # Export all new tags as a JSON array
  new_tags_json=$(printf '%s\n' "${new_tags[@]}" | jq -R . | jq -s .)
  echo "new_tags=$new_tags_json" >> "$GITHUB_OUTPUT"
fi

echo ""
echo "=== Detection complete ==="
