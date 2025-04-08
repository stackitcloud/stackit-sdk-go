# Release

## Release cycle

A release of the whole SDK should be created at least every 2 weeks. 

## Release creation

### Single service

The SDK is split into all the different STACKIT [services](https://github.com/stackitcloud/stackit-sdk-go/tree/main/services), each of them having their own version numbers.

**Checklist before releasing a single service**

- [ ] Changelog entries were added to the `CHANGELOG.md` in the root directory of the repository (see [here](https://github.com/stackitcloud/stackit-sdk-go/blob/f1e375a38064f798821d22a951bc74ca8a9c8845/CHANGELOG.md))
- [ ] Changelog entries were added to the `CHANGELOG.md` file of the service to be released (see e.g. [here](https://github.com/stackitcloud/stackit-sdk-go/blob/f1e375a38064f798821d22a951bc74ca8a9c8845/services/dns/CHANGELOG.md))

**Releasing a single service**

1. Check out latest main branch on your machine
2. Create git tag: `git tag services/<SERVICE-NAME>/vX.X.X`
    - E.g. for the `sqlserverflex` service version `v1.0.1` the git tag would be named `services/sqlserverflex/v1.0.1`
3. Push the git tag: `git push origin --tags`

### Whole SDK

**Checklist before releasing the whole SDK**

- [ ] Date was set/updated in the `CHANGELOG.md` file in the root directory of the repository (see [here](https://github.com/stackitcloud/stackit-sdk-go/blob/f1e375a38064f798821d22a951bc74ca8a9c8845/CHANGELOG.md))

**Releasing the whole SDK**

> [!IMPORTANT]
> Consider informing / syncing with the team before creating a new release.

1. Check out latest main branch on your machine
2. Create git tag: `git tag release-YYYY-MM-DD`
3. Push the git tag: `git push origin --tags`
4. Copy the changelog entries for the new release from the `CHANGELOG.md` file in the root directory of the repository (see [here](https://github.com/stackitcloud/stackit-sdk-go/blob/f1e375a38064f798821d22a951bc74ca8a9c8845/CHANGELOG.md)) to your clipboard.
5. Go to the [releases page](https://github.com/stackitcloud/stackit-sdk-go/releases) on GitHub and create a new release. Select the git tag you just created.
6. Before creating the GitHub release, add the *Highlights* heading at the top of the markdown description and paste the changelog entries from your clipboard (see [previous releases](https://github.com/stackitcloud/stackit-sdk-go/releases/tag/release-2025-03-27) to see what it should look like). Then create and publish the GitHub release.

