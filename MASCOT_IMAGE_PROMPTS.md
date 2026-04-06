# DG Mascot V1 Image Prompts

## Recommendation First

Yes. Start with **one strong front-facing idle master image first**, then use remix / reference-image generation to derive the other poses.

That is the best V1 strategy if consistency matters.

### Why this is the right order

- it locks the mascot identity early
- it stabilizes silhouette, proportions, face, and color palette
- it reduces the chance that each pose looks like a different creature
- it makes later cleanup easier if you redraw or vectorize by hand

### Best workflow

1. Generate a **single polished master idle image**
2. Pick the best one and lock it as the visual anchor
3. Use that idle image as the reference/remix source for:
   - `wave`
   - `nod`
   - `think`
   - `rest`
   - `spark`
4. After that, clean or redraw the final set into SVG if needed

### What to keep fixed across every remix

- body silhouette
- face shape
- eye size and placement
- spark nub shape
- main palette
- overall camera angle

### What can change across remixes

- eye expression
- mouth shape
- slight body tilt
- accent emphasis
- tiny implied gesture shape

## Global Generation Rules

Use these rules for every image:

- transparent background
- centered composition
- front-facing or near-front-facing
- compact rounded tool spirit
- dark teal body
- warm gold spark nub
- simple clean illustration
- readable at small size
- not humanoid
- not anime
- not corporate mascot
- not hyper-rendered
- no props
- no text
- no scene background

## Global Negative Prompt

Use this negative prompt for every generation:

> humanoid, human body, detailed hands, realistic anatomy, anime character, corporate assistant mascot, heavy armor, robot, mechanical limbs, over-detailed shading, background scene, text, speech bubble, photo realism, horror, creepy eyes, complex accessories, clutter, side profile, extreme perspective

## Master Prompt: Idle Anchor

This is the most important image in the set.

### Prompt

> A small rounded mascot spirit for a desktop learning app, front-facing, shaped like a compact droplet-like tool sprite, dark teal body, warm gold spark nub on top, large readable eyes, tiny expressive mouth, clever and slightly mischievous but warm, calm and observant, simple vector-friendly silhouette, clean flat illustration, transparent background, centered composition, highly readable at small size, not humanoid, not anime, charming and compact.

### Optional style suffix

> soft flat shading, subtle clean outline separation, polished app mascot design, minimal but expressive face

### Goal

This image should define:

- final silhouette
- final face proportions
- final spark nub shape
- final palette
- final vibe

## Prompt 02: Wave

Use the chosen idle image as the reference/remix source.

### Prompt

> Same mascot character as the reference image, keep the exact same silhouette, face proportions, spark nub, and color palette. Show a welcoming wave pose with a slight body tilt, brighter eyes, a small friendly smile, and one side subtly lifted as if waving. Transparent background, centered composition, same camera angle, simple clean flat illustration, readable at small size.

### Extra guidance

- keep identity nearly identical to idle
- only change expression and pose energy
- do not redesign the character

## Prompt 03: Nod

Use the chosen idle image as the reference/remix source.

### Prompt

> Same mascot character as the reference image, preserve the exact silhouette, spark nub, facial proportions, and colors. Show an approving nod pose with slight downward compression, pleased eyes, and a confident small smile. Transparent background, centered composition, same camera angle, simple clean flat illustration, readable at small size.

### Extra guidance

- should feel approving, not overexcited
- keep the body compact

## Prompt 04: Think

Use the chosen idle image as the reference/remix source.

### Prompt

> Same mascot character as the reference image, preserve the exact silhouette, spark nub, facial proportions, and palette. Show a thoughtful pose with a slight lean, mildly asymmetrical eyes, a tiny puzzled mouth, and a subtle questioning energy. Transparent background, centered composition, same camera angle, simple clean flat illustration, readable at small size.

### Extra guidance

- thoughtful, not sad
- curious, not scolding

## Prompt 05: Rest

Use the chosen idle image as the reference/remix source.

### Prompt

> Same mascot character as the reference image, preserve the exact silhouette, spark nub, face proportions, and color palette. Show a calm rest pose with slightly lower posture, softer eyes, a relaxed mouth, and a cozy paused energy. Transparent background, centered composition, same camera angle, simple clean flat illustration, readable at small size.

### Extra guidance

- should look restful, not sleepy or sad
- keep the silhouette recognizable

## Prompt 06: Spark

Use the chosen idle image as the reference/remix source.

### Prompt

> Same mascot character as the reference image, preserve the exact silhouette, face proportions, spark nub identity, and palette. Show a celebratory spark pose with brighter expression, open posture, energized eyes, a proud smile, and slightly emphasized spark details. Transparent background, centered composition, same camera angle, simple clean flat illustration, readable at small size.

### Extra guidance

- special, but still suitable for a calm productivity app
- can include tiny surrounding spark accents if they stay minimal

## Optional Prompt 07: Collapsed Badge

If the full character reads poorly when the bubble is collapsed, generate a simplified badge version.

### Prompt

> Same mascot character as the reference image, simplified into a compact face-first badge icon, keeping the same eyes, spark nub, and character identity. Very clear and readable at tiny size, transparent background, centered composition, simple clean flat illustration.

## Prompting Notes For Better Consistency

### If your tool supports reference strength

Use:

- medium-high identity preservation
- low-medium pose deviation for most poses
- medium deviation for `spark`

### If your tool supports seed reuse

Reuse the same seed from the chosen idle image for all derivative poses when possible.

### If your tool supports character reference

Always use the selected idle anchor image as the single source of truth.

Do not mix multiple reference images in V1 unless the tool absolutely requires it.

## Evaluation Checklist

After generation, check each pose:

- does it still clearly look like the same mascot?
- does the spark nub stay recognizable?
- do the eyes keep the same identity?
- does the palette remain stable?
- can the emotion be understood at a small size?
- does it still fit the app's calm dark UI?

## Production Advice

Best practical approach:

1. generate `idle` until it is excellent
2. freeze it
3. generate all other poses by remix/reference
4. choose the best variant of each
5. clean them into final SVG assets later

This is safer than generating six unrelated prompts from scratch.

## DG Pet V2-A Image Prompts

These prompts are for stage-based visual growth variants.

Use the currently shipped mascot SVGs or their high-quality PNG exports as the identity anchor.

Do not redesign DG here. The job is to make the same DG feel subtly more alive across hidden stages.

### V2-A Workflow

1. Take the current pose asset as the base image
2. Generate stage `0`, `1`, and `2` variants for that same pose
3. Keep identity locked
4. Only change liveliness, glow, gloss, and expression richness

---

### Pose: idle

#### Stage 0

Base image:

- current `idle.svg` or a clean PNG export from it

Image prompt:

> Same mascot character as the base image, preserve the exact silhouette, face shape, spark nub identity, and overall palette. Create a stage 0 idle variant that feels calmer and simpler, with softer gloss, lower glow, gentler eye highlights, and a more reserved magical presence. Transparent background, centered composition, same camera angle, clean vector-friendly illustration.

Enhance tips:

- reduce extra shine
- keep the spark nub warm but not bright
- avoid making the face dull or sleepy

#### Stage 1

Base image:

- current `idle.svg`

Image prompt:

> Same mascot character as the base image, preserve the exact silhouette, face shape, spark nub identity, and overall palette. Create a stage 1 idle variant that feels slightly more awake and companion-like, with stronger gloss, brighter eye highlights, a mildly brighter spark nub, and a warmer polished presence. Transparent background, centered composition, same camera angle, clean vector-friendly illustration.

Enhance tips:

- increase gloss slightly, not dramatically
- let the eyes feel more alert
- keep the stage jump subtle compared with stage 0

#### Stage 2

Base image:

- current `idle.svg`

Image prompt:

> Same mascot character as the base image, preserve the exact silhouette, face shape, spark nub identity, and overall palette. Create a stage 2 idle variant that feels the most alive and quietly magical, with the richest gloss, the brightest controlled spark nub, slightly more expressive eyes, and a polished bonded presence. Transparent background, centered composition, same camera angle, clean vector-friendly illustration.

Enhance tips:

- strongest glow of the three, but still tasteful
- do not add props or new body parts
- keep it feeling like the same DG, only more alive

---

### Pose: wave

#### Stage 0

Base image:

- current `wave.svg`

Image prompt:

> Same mascot character as the base image, preserve the exact silhouette, face proportions, spark nub identity, and color palette. Create a stage 0 wave variant that feels friendly but modest, with a lighter greeting energy, softer gloss, and restrained glow. Transparent background, centered composition, same camera angle, clean vector-friendly illustration.

Enhance tips:

- keep the smile small
- keep the greeting visible without making it too excited
- avoid large sparkle accents

#### Stage 1

Base image:

- current `wave.svg`

Image prompt:

> Same mascot character as the base image, preserve the exact silhouette, face proportions, spark nub identity, and palette. Create a stage 1 wave variant that feels warmer and more familiar, with clearer gloss, brighter eyes, a livelier spark nub, and a more companion-like greeting energy. Transparent background, centered composition, same camera angle, clean vector-friendly illustration.

Enhance tips:

- let the expression feel more welcoming
- add a touch more polish to the body
- keep the shape identical to the base pose

#### Stage 2

Base image:

- current `wave.svg`

Image prompt:

> Same mascot character as the base image, preserve the exact silhouette, face proportions, spark nub identity, and palette. Create a stage 2 wave variant that feels most alive, cheerful, and familiar, with the richest gloss, strongest controlled spark-nub energy, and the clearest eye expression while still staying compact and calm. Transparent background, centered composition, same camera angle, clean vector-friendly illustration.

Enhance tips:

- greeting should feel confident, not loud
- add a very small extra magical warmth
- avoid turning it into a celebration pose

---

### Pose: spark

#### Stage 0

Base image:

- current `spark.svg`

Image prompt:

> Same mascot character as the base image, preserve the exact silhouette, face shape, spark nub identity, and palette. Create a stage 0 spark variant that feels celebratory but still relatively simple, with moderate glow, restrained accent sparks, and a clean bright smile. Transparent background, centered composition, same camera angle, clean vector-friendly illustration.

Enhance tips:

- keep effects minimal
- make the celebration read clearly without looking flashy
- keep accent spark marks sparse

#### Stage 1

Base image:

- current `spark.svg`

Image prompt:

> Same mascot character as the base image, preserve the exact silhouette, face shape, spark nub identity, and palette. Create a stage 1 spark variant that feels brighter and more rewarding, with stronger glow, richer gloss, more energized eye highlights, and slightly more expressive spark accents. Transparent background, centered composition, same camera angle, clean vector-friendly illustration.

Enhance tips:

- add glow before adding new shapes
- keep spark accents small and tidy
- make the face more delighted, not hyperactive

#### Stage 2

Base image:

- current `spark.svg`

Image prompt:

> Same mascot character as the base image, preserve the exact silhouette, face shape, spark nub identity, and palette. Create a stage 2 spark variant that feels the richest, brightest, and most bonded version of the same mascot, with the strongest tasteful glow, clearest gloss, most energized spark nub, and slightly richer celebratory accent spark marks. Transparent background, centered composition, same camera angle, clean vector-friendly illustration.

Enhance tips:

- this should be the most special version, but still belong to the same calm app
- avoid overwhelming the face with effects
- keep the silhouette unchanged and the identity fully intact
