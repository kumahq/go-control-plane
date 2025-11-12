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

Our sync automation detects all these tags and applies the `+kong-N` suffix to **all of them**, creating:

- `v0.13.4+kong-1`
- `envoy/v1.32.3+kong-1`
- `contrib/v1.32.3+kong-1`
- `xdsmatcher/v0.13.4+kong-1`
- `ratelimit/v0.1.0+kong-1`

### Manual Releases

If you merge a custom change and want to create a release, you need to tag your commit with **all module tags** using the following pattern:

```text
<module-prefix><current-version>+kong-<incremented-number>
```

### Why use `+` instead of `-`?

- Per Semantic Versioning (https://semver.org/), anything after a hyphen (`-`) denotes a pre-release (e.g., `1.2.3-alpha.1`). Many tooling ecosystems interpret pre-releases as being "behind" the final release and will constantly flag that a newer non-pre-release version is available
- Anything after a plus (`+`) is build metadata. Build metadata does not affect version precedence. By using tags like `vX.Y.Z+kong-N`, we keep the base version identical to upstream and attach our fork-specific identifier without triggering dependency updaters to suggest moving to plain `vX.Y.Z`

#### Example

The current tags are:
- `v0.13.1+kong-1`
- `envoy/v1.32.2+kong-1`
- `xdsmatcher/v0.13.1+kong-1`

After merging your changes, you want to release a new version. You should tag your commit with **all** the following tags:

- `v0.13.1+kong-2`
- `envoy/v1.32.2+kong-2`
- `xdsmatcher/v0.13.1+kong-2`
- (and any other module tags present)

```bash
# Example: Creating all tags for a custom release
git tag v0.13.1+kong-2
git tag envoy/v1.32.2+kong-2
git tag xdsmatcher/v0.13.1+kong-2
git tag ratelimit/v0.1.0+kong-2
git push origin --tags
```

After tagging, push the commit to the repository to finalize the release.

## Conflict resolution

It might happen that the automatic release job can hit a conflict. In this situation the maintainer needs to resolve this manually

### Instruction

1. Fetch current `main` and `release` branch

   ```bash
   git fetch origin --tags
   ```

2. Update the current main branch with origin

   ```bash
   git reset --hard origin/main
   ```

3. Change to the release branch

   ```bash
   git checkout release
   git reset --hard origin/release
   ```

4. Remove Untracked files and directories

   ```bash
   git clean -fd
   ```

5. Get the new tag value

   ```bash
   ./.github/scripts/get-version-tags.sh
   ```

   **Output:**
   
   ```bash
   Current release tag prefix: v0.13.1
   Upstream tag: v0.13.2
   New tag: v0.13.2+kong-1
   ```

   if tags are equal, you are going to see message `"Tags are equal, no need to release"` and you don't need to release/rebase anything

6. Rebase release branch between tags:

   ```bash
   git rebase --onto "<upstream-tag>" <current-release-tag-prefix> release
   ```

7. If you encounter git conflicts, resolve them, and follow git instruction

   - resolve conflicts
   - `git add <files>`
   - `git rebase --continue`

8. Tag the commit

   ```bash
   git tag "<new-tag>"
   ```

9. Push changes to origin

   ```bash
   git push origin release --tags --force-with-lease
   ```
