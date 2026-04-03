```markdown
# Technical Precision Design System

## 1. Overview & Creative North Star: "The Monolithic Terminal"
This design system is engineered for high-velocity technical environments. It rejects the "soft" aesthetics of modern consumer apps in favor of **The Monolithic Terminal**—a North Star that prioritizes data density, optical precision, and zero-latency comprehension. 

The system moves beyond generic grids by utilizing **intentional asymmetry** and **brutalist architectural layering**. By removing the "clutter of the curve" (standardizing at a 0px radius), we create a visual language of absolute certainty. The layout mimics a high-end instrument cluster: high-contrast numerical readouts are paired with muted metadata to ensure the signal is never lost in the noise.

---

## 2. Colors & Surface Logic
The palette is a sophisticated interplay of deep architectural shadows and high-energy terminal light.

### Tonal Architecture
- **Primary Surface (`surface_container_lowest` / #0a0e14):** The foundation. Use this for the deepest background layers.
- **Data Surfaces (`surface_container` / #1c2026):** For primary modules and data blocks.
- **Interaction Layers (`surface_bright` / #353940):** Reserved for hover states or active utility zones.

### The "No-Line" Rule
To maintain a premium, integrated feel, **1px solid borders are strictly prohibited for sectioning.** 
- Boundary definition must be achieved through **Surface Hierarchy**. A module sitting on the background is defined by the shift from `#0a0e14` to `#1c2026`.
- **Nesting:** When a component requires an internal container, shift the tone one tier higher (`surface_container_high`). This creates "Tonal Elevation" without the visual friction of lines.

### Signature Textures
While we avoid glassmorphism, we achieve "soul" through **Linear Precision Gradients**. For active data bars or primary CTAs, use a subtle vertical gradient from `primary` (#c3f5ff) to `primary_container` (#00e5ff). This mimics the glow of a high-end CRT or laser-etched display.

---

## 3. Typography: The Signal Hierarchy
The typography system uses a dual-font strategy to balance technical grit with editorial authority.

- **The Data Voice (Space Grotesk):** Used for `display` and `headline` roles. This font’s geometric tension makes numerical data feel engineered rather than typed.
- **The Metadata Voice (Inter):** Used for `body` and `label` roles. High legibility at small scales (down to 11px) ensures that secondary technical specs remain readable in high-density views.

**Editorial Tip:** Use `label-sm` in `secondary` (#adcbda) for all units of measurement (e.g., "ms", "kb/s") immediately following a `title-lg` numerical value to create an "Instrumentation" look.

---

## 4. Elevation & Depth: Tonal Layering
Traditional shadows and rounded cards are replaced by a system of **Geometric Stacking**.

- **The Layering Principle:** Depth is conveyed by "stacking" surface tiers. A `surface_container_highest` element naturally feels closer to the user than a `surface_dim` background.
- **Ambient Glow (The "Warning" Lift):** For critical alerts using the `error` token (#ffb4ab), apply a diffused 4% opacity outer glow of the same color. This mimics light emission from a technical warning lamp.
- **The Ghost Border Fallback:** If a boundary is essential for accessibility in ultra-dense charts, use the `outline_variant` token (#3b494c) at **15% opacity**. It must feel like a "whisper" of a line, not a structural wall.

---

## 5. Components

### Utility Buttons (The Switchers)
- **Geometry:** 0px radius, compact padding (4px 8px).
- **State:** Active states use `primary_container` with `on_primary_container` text. Inactive states use `surface_container_high` with `secondary` text.
- **Interaction:** On hover, shift background to `surface_bright`. No transition animations; changes should be instantaneous (0ms) to reflect "fast-to-read" performance.

### Data Bars & Dense Charts
- **Primary Bars:** Use `primary_container`. For "over-budget" states, use a hard-stop color swap to `on_tertiary_container` (#b01522).
- **Grid Lines:** Forbid standard grid lines. Use `surface_container_highest` for the X/Y axis baseline only.

### Numerical Readouts
- **Primary Metric:** `headline-lg` in `primary` (#c3f5ff). 
- **Sub-Metric:** `label-md` in `secondary`. 
- **Alignment:** Always use tabular figures (monospaced numbers) to prevent "jumping" layouts during real-time data refreshes.

### Input Fields
- **Styling:** No borders. Use `surface_container_low` for the field background. 
- **Focus State:** A 2px bottom-bar in `primary_container`. Do not outline the entire box.

---

## 6. Do’s and Don’ts

### Do:
- **Do** maximize data density. If there is empty space, consider if more telemetry can be surfaced.
- **Do** use `tertiary_fixed_variant` (#930015) for background zones containing critical errors—this creates a "Red Alert" zone that is impossible to miss.
- **Do** align all elements to a strict 4px baseline grid to maintain the "engineered" feel.

### Don’t:
- **Don’t use rounded corners.** Everything is a hard 90-degree angle.
- **Don’t use drop shadows.** Hierarchy is color-driven, not light-driven.
- **Don’t use purple or soft pastels.** The palette must remain strictly within the Navy, Cyan, and Orange-Red spectrum to maintain its professional "Control Center" identity.
- **Don’t use dividers.** Use a 16px or 24px vertical gap from the Spacing Scale to separate logical groups.

---

## 7. Signature Layout: The "Asymmetric Split"
When designing a dashboard view, avoid the 50/50 split. Use an **80/20 "Technical Sidestrip."** The 80% zone contains high-density charts on `surface_container_low`, while the 20% sidestrip uses `surface_container_lowest` for a vertical stream of raw log data or system alerts. This creates a functional imbalance that directs the eye to the most critical data first.```