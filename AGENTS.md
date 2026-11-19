# Agent instructions

## Design system

Follow the design system in [`DESIGN.md`](DESIGN.md) for **every** change to the interface
(`template/listing.gohtml`, any new page or component): colors, typography, spacing, shapes,
components, themes, and the accessibility rules (RGAA 4 / WCAG 2.1 AA).

- Read `DESIGN.md` before touching the UI. Use its tokens; do not invent colors, sizes or components.
- If the design needs to change, update `DESIGN.md` first, then the template, in the same change.
- Never break an accessibility rule of `DESIGN.md` to match a mockup: keep the contrasts, visible
  focus, labels and reduced-motion behavior, and say which rule forced a deviation.
- Both themes (light and dark) must keep working.
- Show only data the app really has (service ID, title, URL). Do not invent durations, ratings
  or descriptions.

## Project

- Go app for App Engine (`app.yaml`, runtime `go126`); the page is rendered from
  `template/listing.gohtml`. See [`README.MD`](README.MD) for running, deploying and environments.
- Before finishing: `go build ./... && go vet ./...`, and check the page in both themes.
