# DG Mascot V1 Asset Briefs

## Purpose

This document turns the visual direction in [MASCOT_SPEC.md](/D:/duolin-gogo/MASCOT_SPEC.md) into production-ready asset briefs for the first mascot pack.

The goal is to create a first usable set of mascot assets that can replace the current `DG` text badge in the UI.

## Global Art Rules

Apply these rules to every pose.

### Character Core

- character type: small rounded tool spirit
- silhouette: compact, droplet-like, readable at tiny size
- personality: clever, warm, slightly cheeky
- style: clean, vector-friendly, lightly whimsical
- body: soft rounded blob with one small spark nub on top
- face: large readable eyes, tiny expressive mouth
- limbs: optional implied stubs only, do not rely on detailed arms/hands

### Rendering Rules

- transparent background
- front-facing or slightly three-quarter view only
- clean silhouette first, detail second
- avoid clutter, props, scenery, text, or speech bubbles
- no heavy shading
- prefer flat or lightly layered shading
- no realistic textures

### Palette

- primary dark body: `#183345`
- deep shell support tone: `#10212F`
- warm accent: `#E9C46A`
- light face details: `#F7F3E8`
- optional soft secondary accent: `#7FA0B8`

### Output Target

- preferred format: SVG
- if raster is necessary: PNG with transparent background
- recommended artboard: square
- recommended working size: `512 x 512`
- keep composition centered
- mascot should occupy most of the frame without touching edges

## Shared Prompt Base

Use this as the base prompt for image generation or illustrator handoff.

> A small rounded mascot spirit for a desktop learning app, shaped like a compact droplet-like tool sprite, dark teal body, warm gold spark nub accent, large readable eyes, tiny expressive mouth, clever and slightly mischievous but warm, simple vector-friendly silhouette, transparent background, clean flat illustration, not humanoid, not anime, not corporate, highly readable at small size.

## Asset 01: idle

### File

- `app/frontend/src/assets/dg/idle.svg`

### UI Purpose

- default study state
- neutral learn state
- calm companion presence

### Visual Intent

DG is quietly attentive and ready.

### Pose Notes

- upright posture
- body balanced and centered
- spark nub relaxed, slightly tilted
- eyes open but soft
- tiny half-smile or neutral mouth

### Emotional Tone

- calm
- observant
- steady

### Pose Prompt Add-on

> Idle pose, upright and centered, soft eyes, tiny calm smile, relaxed spark nub, attentive but quiet presence.

### Must-Haves

- reads clearly at tiny size
- does not look sleepy
- does not look too excited

## Asset 02: wave

### File

- `app/frontend/src/assets/dg/wave.svg`

### UI Purpose

- user returns after break
- welcome-back state
- some click interactions

### Visual Intent

DG is greeting the user in a friendly, familiar way.

### Pose Notes

- slight side tilt
- one side of the body or stub lifted in a wave-like gesture
- brighter eyes
- spark nub more lifted and lively
- small friendly smile

### Emotional Tone

- welcoming
- light
- approachable

### Pose Prompt Add-on

> Friendly wave pose, slight tilt, bright eyes, welcoming smile, one side lifted like a small wave, spark nub lively.

### Must-Haves

- should feel like a greeting even without visible hands
- avoid overexcited motion

## Asset 03: nod

### File

- `app/frontend/src/assets/dg/nod.svg`

### UI Purpose

- correct answer
- positive reinforcement

### Visual Intent

DG is giving a confident, approving reaction.

### Pose Notes

- slight vertical compression as if dipping into a nod
- pleased eyes or subtly curved eyes
- satisfied smile
- spark nub perked upward

### Emotional Tone

- approving
- proud
- encouraging

### Pose Prompt Add-on

> Approval nod pose, slightly compressed as if dipping forward, pleased eyes, confident smile, spark nub perked up.

### Must-Haves

- should read as approval, not excitement
- compact enough to still look good in the bubble avatar slot

## Asset 04: think

### File

- `app/frontend/src/assets/dg/think.svg`

### UI Purpose

- wrong answer
- reflective prompt
- "look once more" moments

### Visual Intent

DG is thoughtful, not disappointed.

### Pose Notes

- slight lean
- one eye slightly narrower or asymmetrical
- tiny tilted or flat mouth
- spark nub curved or angled as if questioning

### Emotional Tone

- curious
- reflective
- gently skeptical

### Pose Prompt Add-on

> Thoughtful pose, slight lean, mildly asymmetrical eyes, tiny puzzled mouth, spark nub bent like a question cue, reflective not sad.

### Must-Haves

- must not look upset or scolding
- should feel like "let's look again"

## Asset 05: rest

### File

- `app/frontend/src/assets/dg/rest.svg`

### UI Purpose

- learn break
- pause between batches

### Visual Intent

DG is also taking a short pause with the user.

### Pose Notes

- lower posture
- body slightly tucked
- eyelids softer or lower
- relaxed mouth
- spark nub softened, less upright

### Emotional Tone

- cozy
- restful
- unhurried

### Pose Prompt Add-on

> Rest pose, lower tucked posture, softer eyes, relaxed mouth, spark nub softened, calm and cozy break-state.

### Must-Haves

- should read as a pause, not sadness
- avoid looking asleep unless very lightly implied

## Asset 06: spark

### File

- `app/frontend/src/assets/dg/spark.svg`

### UI Purpose

- review completion
- notable success moment
- special encouragement

### Visual Intent

DG is delighted and proud without becoming chaotic.

### Pose Notes

- more open posture
- brighter smile
- eyes more energized
- spark nub emphasized
- optional tiny accent spark marks around the head

### Emotional Tone

- celebratory
- proud
- bright

### Pose Prompt Add-on

> Celebration pose, open posture, bright smile, energized eyes, spark nub emphasized, tiny surrounding spark accents, proud and cheerful.

### Must-Haves

- should feel special compared with all other states
- still needs to fit a calm productivity app

## Optional Early Variant: collapsed badge

This is optional for V1, but can help if the normal pose art does not read well in the collapsed state.

### File

- `app/frontend/src/assets/dg/badge.svg`

### Purpose

- used only when the assistant bubble is collapsed

### Visual Intent

A simplified mascot head or face-only version that remains readable at tiny size.

### Prompt Add-on

> Simplified mascot badge icon, face-first crop, same character identity, minimal details, high readability at very small size.

## Acceptance Checklist

Each asset should pass these checks:

- looks like the same character as the other poses
- readable when shrunk down
- emotion is understandable without text
- silhouette remains clean on dark backgrounds
- accent gold does not overpower the face
- works beside a single-line hint sentence

## Recommended Production Order

Make assets in this order:

1. `idle`
2. `think`
3. `nod`
4. `wave`
5. `rest`
6. `spark`

Why:

- `idle`, `think`, and `nod` will be seen most often
- once those three are solid, the character identity is usually stable enough to finish the rest quickly

## Handoff Note

If these assets are generated first and cleaned later, prioritize:

- consistent silhouette
- face readability
- matching color palette

Perfect polish can come after the first in-app integration.

## DG Pet V2-A Asset Briefs

These briefs are for stage-based variants of the existing mascot, not for a redesign.

### Asset Strategy

Use the current shipped mascot as the identity anchor.

For V2-A:

- keep one shared character
- preserve the current pose set
- generate stage variants only for the first three high-value poses:
  - `idle`
  - `wave`
  - `spark`

Suggested filenames:

- `idle-stage0.svg`
- `idle-stage1.svg`
- `idle-stage2.svg`
- `wave-stage0.svg`
- `wave-stage1.svg`
- `wave-stage2.svg`
- `spark-stage0.svg`
- `spark-stage1.svg`
- `spark-stage2.svg`

### Stage 0 Brief

Character should feel:

- new
- calmer
- slightly simpler

Art cues:

- weakest glow
- simplest gloss
- least energetic spark nub
- lowest amount of tiny accent sparkle

### Stage 1 Brief

Character should feel:

- more awake
- more familiar
- more responsive

Art cues:

- stronger gloss
- more readable eye highlight
- slightly brighter spark nub
- a touch more warmth in gold accents

### Stage 2 Brief

Character should feel:

- most alive
- quietly proud
- polished and bonded

Art cues:

- strongest glow while staying tasteful
- clearest body highlights
- most energized spark nub
- optional tiny accent sparkle for `spark`
- most expressive but still compact face treatment

### Production Recommendation

Generate stage variants in this order:

1. `idle-stage0`, `idle-stage1`, `idle-stage2`
2. `wave-stage0`, `wave-stage1`, `wave-stage2`
3. `spark-stage0`, `spark-stage1`, `spark-stage2`

Why:

- `idle` will make stage differences visible most often
- `wave` makes welcome/greeting energy easier to read
- `spark` makes celebration-stage richness easiest to notice
