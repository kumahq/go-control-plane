# Contributing

We welcome contributions from the community. Please read the following guidelines carefully to
maximize the chances of your PR being merged.

## Communication

* Before starting work on a major feature, please reach out to us via GitHub, Slack,
  email, etc. We will make sure no one else is already working on it and ask you to open a
  GitHub issue.
* A "major feature" is defined as any change that is > 100 LOC altered (not including tests), or
  changes any user-facing behavior. We will use the GitHub issue to discuss the feature and come to
  agreement. This is to prevent your time being wasted, as well as ours. The GitHub review process
  for major features is also important so that [organizations with commit access](OWNERS.md) can
  come to agreement on design. If it is appropriate to write a design document, the document must
  be hosted either in the GitHub tracking issue, or linked to from the issue and hosted in a
  world-readable location.
* Small patches and bug fixes don't need prior communication.

## Inclusive language policy

The Envoy community has an explicit goal to be inclusive to all. As such, all PRs must adhere to the
following guidelines for all code, APIs, and documentation:

* The following words and phrases are not allowed:
  * *Whitelist*: use allowlist instead.
  * *Blacklist*: use denylist or blocklist instead.
  * *Master*: use primary instead.
  * *Slave*: use secondary or replica instead.
* Documentation should be written in an inclusive style. The [Google developer
  documentation](https://developers.google.com/style/inclusive-documentation) contains an excellent
  reference on this topic.
* The above policy is not considered definitive and may be amended in the future as industry best
  practices evolve. Additional comments on this topic may be provided by maintainers during code
  review.

## Submitting a PR

* Fork the repo.
* Create your PR.
* Tests will automatically run for you.
* We will **not** merge any PR that is not passing tests.
* PRs are expected to have 100% test coverage for added code. This can be verified with a coverage
  build. If your PR cannot have 100% coverage for some reason please clearly explain why when you
  open it.
* Any PR that changes user-facing behavior **must** have associated documentation in [docs](docs) as
  well as the [changelog](CHANGELOG.md).
* All code comments and documentation are expected to have proper English grammar and punctuation.
  If you are not a fluent English speaker (or a bad writer ;-)) please let us know and we will try
  to find some help but there are no guarantees.
* Your PR title should be descriptive, and generally start with a subsystem name followed by a
  colon. Examples:
  * "docs: fix grammar error"
  * "http conn man: add new feature"
* Your PR commit message will be used as the commit message when your PR is merged. You should
  update this field if your PR diverges during review.
* Your PR description should have details on what the PR does. If it fixes an existing issue it
  should end with "Fixes #XXX".
* If your PR is co-authored or based on an earlier PR from another contributor,
  please attribute them with `Co-authored-by: name <name@example.com>`. See
  GitHub's [multiple author
  guidance](https://help.github.com/en/github/committing-changes-to-your-project/creating-a-commit-with-multiple-authors)
  for further details.
* When all of the tests are passing and all other conditions described herein are satisfied, a
  maintainer will be assigned to review and merge the PR.
* Once you submit a PR, *please do not rebase it*. It's much easier to review if subsequent commits
  are new commits and/or merges. We squash rebase the final merged commit so the number of commits
  you have in the PR don't matter.
* We expect that once a PR is opened, it will be actively worked on until it is merged or closed.
  We reserve the right to close PRs that are not making progress. This is generally defined as no
  changes for 7 days. Obviously PRs that are closed due to lack of activity can be reopened later.
  Closing stale PRs helps us to keep on top of all of the work currently in flight.

## PR review policy for maintainers

* See [OWNERS.md](OWNERS.md) for the current list of maintainers.
* It is generally expected that a senior maintainer should review every PR.
* It is also generally expected that a "domain expert" for the code the PR touches should review the
  PR. This person does not necessarily need to have commit access.
* The previous two points generally mean that every PR should have two approvals. (Exceptions can
  be made by the senior maintainers).
* The above rules may be waived for PRs which only update docs or comments, or trivial changes to
  tests and tools (where trivial is decided by the maintainer in question).
* In general, we should also attempt to make sure that at least one of the approvals is *from an
  organization different from the PR author.* E.g., if Lyft authors a PR, at least one approver
  should be from an organization other than Lyft. This helps us make sure that we aren't putting
  organization specific shortcuts into the code.
* If there is a question on who should review a PR please discuss in Slack.
* Anyone is welcome to review any PR that they want, whether they are a maintainer or not.
* Please make sure that the PR title, commit message, and description are updated if the PR changes
  significantly during review.
* Please **clean up the title and body** before merging. By default, GitHub fills the squash merge
  title with the original title, and the commit body with every individual commit from the PR.
  The maintainer doing the merge should make sure the title follows the guidelines above and should
  overwrite the body with the original commit message from the PR (cleaning it up if necessary)
  while preserving the PR author's final DCO sign-off.

## DCO: Sign your work

Envoy ships commit hooks that allow you to auto-generate the DCO signoff line if
it doesn't exist when you run `git commit`. Simply navigate to the Envoy project
root and run:

```bash
./support/bootstrap
```

From here, simply commit as normal, and you will see the signoff at the bottom
of each commit.

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch. The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](https://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.

If you want this to be automatic you can set up some aliases:

```bash
git config --add alias.amend "commit -s --amend"
git config --add alias.c "commit -s"
```

## Fixing DCO

If your PR fails the DCO check, it's necessary to fix the entire commit history in the PR. Best
practice is to [squash](https://gitready.com/advanced/2009/02/10/squashing-commits-with-rebase.html)
the commit history to a single commit, append the DCO sign-off as described above, and [force
push](https://git-scm.com/docs/git-push#git-push---force). For example, if you have 2 commits in
your history:

```bash
git rebase -i HEAD^^
(interactive squash + DCO append)
git push origin -f
```

Note, that in general rewriting history in this way is a hindrance to the review process and this
should only be done to correct a DCO mistake.

## Triggering CI re-run without making changes

Sometimes tasks will be stuck in CI and won't be marked as failed, which means
the above command won't work. Should this happen, pushing an empty commit should
re-run all the CI tasks. Consider adding an alias into your `.gitconfig` file:

```
[alias]
    kick-ci = !"git commit -s --allow-empty -m 'Kick CI' && git push"
```

Once you add this alias you can issue the command `git kick-ci` and the PR
will be sent back for a retest.

# Releasing

The repository is split in multiple packages, themselves using distinct versions:
 - envoy API
    - `github.com/envoyproxy/go-control-plane/envoy`, including envoy API generated protobuf files
    - `github.com/envoyproxy/go-control-plane/contrib`, including envoy API contrib generated protobuf files
 - ratelimit
    - `github.com/envoyproxy/go-control-plane/ratelimit`, including generated files for envoy RLS (https://github.com/envoyproxy/ratelimit)
 - go-control-plane
    - `github.com/envoyproxy/go-control-plane`
    - `github.com/envoyproxy/go-control-plane/xdsmatcher`

To create a new release, from a clean branch:
 - update the version for the desired part in `versions.yaml` and commit
 - run the following command to update all mod files
```
make multimod/prerelease MODSET={part to release}
```
 - push the resulting branch to get merged in, then create tags with 
 ```
make multimod/push-tags MODSET={part to release}
 ```