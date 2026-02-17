# go-control-plane Fork

This is a fork of [go-control-plane](https://github.com/envoyproxy/go-control-plane). Before introducing any changes here, please create an issue or submit a change to the upstream repository, as it might impact other users

## How to Contribute

In this fork, we maintain two branches:

- `main`: This branch is synced daily with the upstream repository via an automated action. The only custom addition is the job responsible for synchronization
- `release`: This branch contains all user-introduced changes. It diverges from the stable version to ensure no unreleased commits are included. Any changes you make should target this branch

## Steps to Contribute

1. Check out the `release` branch
2. Make your changes and create a pull request targeting the `release` branch
3. Once your pull request is reviewed, merge it into the `release` branch

## Releasing

For standard releases, the `Sync with Upstream` action handles synchronization and creates custom versions whenever there's a new upstream release. We fetch upstream tags, so, for example, `v0.13.1` in this fork corresponds directly to the same version in the upstream repository

### Multi-Module Tagging

Upstream now uses multi-module tagging, creating separate tags for each Go module in the repository. For a single release commit, you'll see multiple tags like:

- `v0.13.4` (root module)
- `envoy/v1.32.3` (envoy module)
- `contrib/v1.32.3` (contrib module)
- `xdsmatcher/v0.13.4` (xdsmatcher module)
- `ratelimit/v0.1.0` (ratelimit module)

Our sync automation detects all these tags and applies the `-kong-N` suffix to **all of them**, creating:

- `v0.13.4-kong-1`
- `envoy/v1.32.3-kong-1`
- `contrib/v1.32.3-kong-1`
- `xdsmatcher/v0.13.4-kong-1`
- `ratelimit/v0.1.0-kong-1`

### Manual Releases

If you merge a custom change and want to create a release, you need to tag your commit with **all module tags** using the following pattern:

```text
<module-prefix><current-version>-kong-<incremented-number>
```

### Why use `+` instead of `-`?

- Per Semantic Versioning (https://semver.org/), anything after a hyphen (`-`) denotes a pre-release (e.g., `1.2.3-alpha.1`). Many tooling ecosystems interpret pre-releases as being "behind" the final release and will constantly flag that a newer non-pre-release version is available
- Anything after a plus (`+`) is build metadata. Build metadata does not affect version precedence. By using tags like `vX.Y.Z-kong-N`, we keep the base version identical to upstream and attach our fork-specific identifier without triggering dependency updaters to suggest moving to plain `vX.Y.Z`

#### Example

The current tags are:
- `v0.13.1-kong-1`
- `envoy/v1.32.2-kong-1`
- `xdsmatcher/v0.13.1-kong-1`

After merging your changes, you want to release a new version. You should tag your commit with **all** the following tags:

- `v0.13.1-kong-2`
- `envoy/v1.32.2-kong-2`
- `xdsmatcher/v0.13.1-kong-2`
- (and any other module tags present)

```bash
# Example: Creating all tags for a custom release
git tag v0.13.1-kong-2
git tag envoy/v1.32.2-kong-2
git tag xdsmatcher/v0.13.1-kong-2
git tag ratelimit/v0.1.0-kong-2
git push origin --tags
```

After tagging, push the commit to the repository to finalize the release.

## Conflict resolution

The automatic release job rebases Kong custom commits on top of new upstream versions. If a conflict occurs, manual resolution is needed.

### Understanding the Branch Structure

- **main branch**: Mirrors upstream, synced daily
- **release branch**: Contains upstream commits + Kong custom commits on top
- **Kong tags** (e.g., `v0.14.0-kong-1`): Point to the HEAD of release (custom commit), not the upstream commit

### Manual Rebase Instructions

1. Fetch current `main` and `release` branches with all tags

   ```bash
   git fetch origin --tags
   ```

2. Checkout and reset the release branch

   ```bash
   git checkout release
   git reset --hard origin/release
   git clean -fd
   ```

3. Detect new tags and check if rebase is needed

   ```bash
   ./.github/scripts/get-version-tags.sh
   ```

   **Example output:**

   ```
   === Detecting new upstream tags ===
   Finding latest Kong tags per module...
     root: v0.13.4
     envoy: envoy/v1.35.0
     ...

   Checking for new upstream tags...
     New tag in root: v0.14.0
     New tag in envoy: envoy/v1.36.0

   Current release root tag: v0.13.4
   Upstream root tag: v0.14.0
   Root version changed - will rebase preserving custom commits
   Current base commit: 2d07f5a1
   Found 2 custom Kong commit(s) on release branch

   New Kong tags to be created:
     - v0.14.0-kong-1
     - envoy/v1.36.0-kong-1

   === Detection complete ===
   ```

   If you see `"No new upstream tags found"`, no action is needed.

4. Use the base commit and new tag from the script output

   ```bash
   BASE_COMMIT="2d07f5a1"      # "Current base commit" from script output
   NEW_TAG="v0.14.0"           # "Upstream root tag" from script output
   ```

5. Rebase Kong custom commits onto the new upstream tag

   ```bash
   git rebase --onto "$NEW_TAG" "$BASE_COMMIT" release
   ```

   This command:
   - Takes all commits after `BASE_COMMIT` on the release branch (your custom Kong commits)
   - Rebases them onto `NEW_TAG` (the new upstream version)
   - Your custom commits remain on top

6. If you encounter conflicts:

   ```bash
   # Resolve conflicts in your editor
   git add <resolved-files>
   git rebase --continue

   # Or abort if needed
   # git rebase --abort
   ```

7. Tag the rebased HEAD with all module tags

   ```bash
   # Apply all Kong tags from the script output
   # Example:
   git tag v0.14.0-kong-1
   git tag envoy/v1.36.0-kong-1
   git tag contrib/v1.36.0-kong-1
   git tag xdsmatcher/v0.14.0-kong-1
   ```

8. Push changes to origin

   ```bash
   git push origin release --tags --force-with-lease
   ```

### Important Notes

- The Kong tags (`-kong-1`) always point to **your custom commit (HEAD of release)**, not the upstream commit
- Custom commits are always **on top** of the upstream base
- When upstream releases a new version, we rebase your commits to keep them on top
- All module tags (root, envoy, contrib, etc.) are created together on the same commit (HEAD)
