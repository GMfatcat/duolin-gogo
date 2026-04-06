# DG Mascot V1 Visual Spec

## Goal

Create the first real visual version of `DG`, replacing the current text-circle assistant badge with a small mascot that feels:

- playful
- slightly mischievous
- warm and encouraging
- compact enough to live inside the current top assistant bubble

The mascot should support the existing study flow and hidden pet-growth system without turning the product into a loud game UI.

## Product Role

`DG` is not a stat-heavy pet and not a full-screen character.

`DG` is:

- a study companion
- a click-bait-flavored guide
- a quiet mascot with more reactions over time
- a small top-of-screen presence that supports learning

`DG` should feel like:

- "a tiny clever helper"
- "a warm but slightly cheeky study spirit"
- "a mascot that gets more alive as the user keeps learning"

## Visual Direction

### Core Shape

Recommended shape:

- rounded body
- slightly droplet-like silhouette
- large readable eyes
- tiny expressive mouth
- one small accent detail on top

Good body silhouette references:

- blob spirit
- small tool sprite
- rounded ghost
- chubby desktop familiar

Avoid:

- realistic anatomy
- humanoid proportions
- detailed fingers/hands in V1
- anime-style complexity
- tiny thin-line features that disappear at small size

### Personality Keywords

Use these keywords when drawing or generating:

- round
- clever
- cheeky
- observant
- warm
- compact
- magical-tool spirit
- low-ego helper
- slightly smug in a funny way

Avoid these keywords:

- childish mascot
- hyper-cute toy
- aggressive gamer mascot
- robotic corporate assistant
- gloomy ghost

## Color Palette

Align the mascot to the current app shell.

### Primary Palette

- warm gold: `#E9C46A`
- cream text: `#F7F3E8`
- shell blue: `#10212F`
- deep navy: `#0B141C`
- muted steel blue: `#9FB3C3`
- soft light blue-gray: `#D4DDE5`

### Mascot Base Recommendation

Base body:

- deep teal-navy base close to `#183345`
- warm highlight accents using `#E9C46A`
- face details in `#F7F3E8`

Optional secondary accent:

- dusty blue accent close to `#7FA0B8`

### Contrast Guidance

The mascot must still read clearly when:

- shown on the dark shell
- shrunk to small avatar size
- collapsed into the compact DG badge state

Avoid overly dark outlines on dark fill. Prefer:

- slightly lighter outer edge
- clear face contrast
- one bright accent point

## Face Language

Face should do most of the work.

Recommended:

- large eyes with clear eyelid variation
- mouth shapes that can switch between:
  - neutral
  - smirk
  - smile
  - tiny "o" surprise
  - thinking line

Avoid:

- eyebrows that require too much detail
- tiny pupils that disappear at small sizes
- overly subtle mouth changes that won't read in a 40–64 px area

## Signature Detail

V1 should include one memorable top/side feature so DG is recognizable even without text.

Pick one:

1. tiny spark nub
2. bent antenna
3. bookmark-like fin
4. little flame curl

Recommended choice:

`tiny spark nub`

Why:

- connects naturally to "idea / click / spark"
- works with the `spark` pose
- easy to stylize in static SVG

## Pose Set

These poses map directly to the current frontend/backend pose system.

### 1. idle

Use for:

- default learn state
- neutral ambient presence

Look:

- upright
- soft eyes
- tiny half-smile or neutral mouth
- spark nub relaxed

Emotion:

- calm
- attentive

### 2. wave

Use for:

- return after break
- welcome back
- light click interaction

Look:

- body tilted slightly
- one stub/side lifted or implied wave shape
- eyes brighter / more open
- spark nub leaning outward

Emotion:

- friendly
- greeting

### 3. nod

Use for:

- correct answer
- encouraging reinforcement

Look:

- body slightly compressed downward
- happy or confident eyes
- small satisfied smile
- spark nub perked upward

Emotion:

- approving
- "yep, that one"

### 4. think

Use for:

- wrong answer
- caution / reflection

Look:

- slight lean
- one eye smaller or narrowed
- tiny tilted mouth
- spark nub curved in a questioning direction

Emotion:

- thoughtful
- "look once more"

### 5. rest

Use for:

- learn break
- pause between batches

Look:

- lower posture
- softer eyes
- relaxed mouth
- body slightly tucked

Emotion:

- cozy
- low-energy
- calm pause

### 6. spark

Use for:

- review completion
- special happy moments

Look:

- brighter expression
- widened eyes or proud smile
- spark nub glowing / emphasized
- optional tiny surrounding spark marks

Emotion:

- celebratory
- proud

## Stage-Based Growth Expression

Do not expose stages to the user, but let visuals slowly feel richer.

### Stage 0

- simplest facial variants
- more restrained mouth shapes
- fewer accent details

### Stage 1

- slightly more expressive mouth/eyes
- stronger tilt in `wave`, `nod`, `think`
- spark nub becomes more characterful

### Stage 2

- richer smiles
- more confident asymmetry
- optional tiny extra accent marks in `spark`
- slightly more animated silhouette feel

Important:

Stage changes should be subtle. Users should feel "DG seems more alive lately," not "the mascot visibly leveled up."

## Animation Guidance

V1 should not use heavy animation systems.

Use only lightweight CSS motion layered on top of static assets:

- small float
- tiny scale pulse
- slight tilt
- gentle bounce for `nod`
- soft settle for `rest`
- tiny glow emphasis for `spark`

Avoid:

- looping attention-seeking bounce
- large idle motion
- anything that makes reading the hint text harder

## Asset Format

Preferred format:

- `SVG`

Recommended files:

- `app/frontend/src/assets/dg/idle.svg`
- `app/frontend/src/assets/dg/wave.svg`
- `app/frontend/src/assets/dg/nod.svg`
- `app/frontend/src/assets/dg/think.svg`
- `app/frontend/src/assets/dg/rest.svg`
- `app/frontend/src/assets/dg/spark.svg`

Optional later:

- stage variants such as `idle-stage2.svg`

But V1 should first ship with one asset per pose.

## Size and Placement

DG must work in two layouts:

1. expanded assistant bubble
2. collapsed compact badge

Recommended avatar box:

- expanded: `44px` to `56px`
- collapsed: `36px` to `44px`

The silhouette must still read clearly at the smaller size.

## Prompt Keywords For Image Generation

If generating concept art with an image model, start from this prompt direction:

> A small rounded mascot spirit for a desktop learning app, slightly mischievous but warm, dark teal body with warm gold spark accent, large readable eyes, tiny expressive mouth, simple vector-friendly silhouette, compact assistant character, clean transparent background, six clear emotional poses, not humanoid, not corporate, not anime, charming and clever.

### Pose Prompt Modifiers

- idle: calm, observant, soft smile
- wave: welcoming, friendly tilt, bright eyes
- nod: pleased, approving, confident smile
- think: curious, reflective, slightly puzzled
- rest: cozy, relaxed, low-energy
- spark: proud, celebratory, bright accent glow

## Implementation Notes

Frontend should eventually map:

- `pose-idle` -> `idle.svg`
- `pose-wave` -> `wave.svg`
- `pose-nod` -> `nod.svg`
- `pose-think` -> `think.svg`
- `pose-rest` -> `rest.svg`
- `pose-spark` -> `spark.svg`

The existing hidden growth system should stay unchanged. V1 only swaps:

- visual carrier
- pose rendering
- subtle animation polish

## V1 Definition Of Done

Mascot V1 is ready when:

- the current `DG` text-circle is replaced by real mascot art
- all 6 current poses have matching assets
- pose switching is visually obvious at normal app size
- the mascot still reads clearly when collapsed
- the mascot does not add visual noise or pull attention away from the study card

## DG Pet V2-A: Stage Visuals

The next mascot phase should make hidden growth feel visible without surfacing any explicit pet stats.

### Goal

Let stage `0`, `1`, and `2` feel progressively more alive through subtle visual differences only.

Users should feel:

- "DG looks a bit brighter lately."
- "The expressions feel richer now."
- "This mascot seems more awake than before."

Users should not feel:

- "I just leveled up my pet."
- "This is a visible RPG progression system."

### V2-A Design Rule

Keep silhouette and identity fixed.

Stage visuals should change:

- glow strength
- gloss richness
- spark nub energy
- eye liveliness
- tiny accent marks

Stage visuals should not change:

- core body shape
- basic palette family
- camera angle
- character identity

### Recommended Pose Scope For First Pass

Start with:

- `idle`
- `wave`
- `spark`

These three poses are enough to make stage growth noticeable without redrawing the whole mascot pack at once.

### Frontend Wiring Rule

The frontend should resolve mascot art with a strict fallback chain:

1. collapsed badge stage asset if the bubble is collapsed
2. pose-specific stage asset such as `idle-stage1`
3. pose default such as `idle.svg`
4. final fallback to `idle.svg`

This makes V2-A safe to ship incrementally:

- start with `idle`, `wave`, and `spark`
- leave `nod`, `think`, and `rest` on the current V1 art
- keep the resolver stable while stage assets arrive over time

### Naming Convention

Recommended staged filenames:

- `idle-stage0.svg`
- `idle-stage1.svg`
- `idle-stage2.svg`
- `wave-stage0.svg`
- `wave-stage1.svg`
- `wave-stage2.svg`
- `spark-stage0.svg`
- `spark-stage1.svg`
- `spark-stage2.svg`

Optional later:

- `collapsed-badge-stage1.svg`
- `collapsed-badge-stage2.svg`

If a staged asset does not exist yet, the frontend should always fall back to the current pose asset.

### Current Frontend Constraint

Right now the frontend reliably knows the hidden pet stage during explicit reaction payloads.

For always-on stage visuals in the shell, a later V2-A implementation slice should also expose current pet stage through ordinary dashboard loading, not only through reaction responses.

### Stage 0

Feel:

- calm
- simple
- reserved

Visual cues:

- lower glow
- simpler gloss
- softer spark nub
- fewer accent spark marks
- gentler eye highlight

### Stage 1

Feel:

- warmer
- more engaged
- slightly more awake

Visual cues:

- stronger gloss on the body
- brighter spark nub
- slightly clearer face contrast
- more lively eye highlights
- slightly richer teal-gold separation

### Stage 2

Feel:

- confident
- familiar
- quietly magical

Visual cues:

- strongest glow within the same restrained palette
- crispest gloss
- spark nub most energized
- richer expression detail
- optional tiny extra accent sparkle in `spark`
- subtle sense of polish, not a costume change
