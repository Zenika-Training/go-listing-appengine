# go-listing-appengine

Home page of https://training-material.zenika.com/: it lists the App Engine services of the
`zen-formations` project (one service per training) as a grid of links.

- `/`: HTML page, rendered from `template/listing.gohtml`. The `default` service (this app) is hidden.
- `/api/services`: the same list in JSON (CORS enabled).

The app reads the list through the [Admin API](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services/list)
of the project given by `GOOGLE_CLOUD_PROJECT` (set automatically on App Engine). A new training
appears as soon as its service is deployed, and disappears when the service is deleted.

# The page

`template/listing.gohtml` is self-contained: no external font, image or library.

- Light and dark themes. It follows the system by default; a **System / Light / Dark** switch
  lets the user override it. The choice is kept in the browser (`localStorage`). Without
  JavaScript, the switch is hidden and the page follows the system.
- Accessibility: AA contrasts in both themes, red used as an accent only (never the only carrier
  of information), visible keyboard focus, text-size friendly units, reduced motion respected.
- Colors are CSS variables at the top of the `<style>` block, defined with `light-dark()`
  (recent browsers, 2024+).

# Run locally

Requires Go 1.26 (`go 1.26` in `go.mod`) and credentials able to list the App Engine services:

```
gcloud auth application-default login
GOOGLE_CLOUD_PROJECT=zen-formations go run listing.go
```

Then open http://localhost:8080 (`PORT` changes the port). Without credentials the page
answers `500`; to work on the design only, execute the template with a few fake services
(see `transformAndDisplay` in `listing.go` for the expected data: `ID`, `Title`, `URL`).

# Environments and deployment

The CircleCI job `deploy` (`.circleci/config.yml`) deploys the `default` service of the
`zen-formations` project on every push. It uses the `training-deploy-gae` context
(`GAE_KEY_FILE_CONTENT`), like the training repositories (see `training-material-orb`).

| Branch | Version | Promoted | URL |
| --- | --- | --- | --- |
| `master` | `go-listing-<sha>` | yes (production) | https://training-material.zenika.com/ |
| any other | `dev-<branch>` | no | `https://dev-<branch>-dot-zen-formations.appspot.com/` |

Workflow: create a branch, push, open the `dev-<branch>` URL (printed in the job log) to
validate, then merge into `master` to publish.

Good to know:

- The dev version uses the same project, so it lists the same trainings as production.
- The branch name is lowercased and cleaned (`[a-z0-9-]`, 40 characters max: a DNS label is
  limited to 63 characters).
- App Engine allows 210 versions for the whole `zen-formations` project (shared with all
  training repositories). After each deploy, the job purges the versions without traffic that are
  older than 7 days (`prune-versions` command of the orb). The version serving production is
  never deleted.
- `app.yaml` sets `automatic_scaling.max_instances` (3), runtime `go126`.

Manual deployment (needs the `appengine.deployer` role):

```
gcloud app deploy --project zen-formations ./app.yaml --version go-listing-XX --promote
```

Use `--no-promote` to publish without taking traffic.
