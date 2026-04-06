import dgCollapsedBadge from './assets/dg/collapsed-badge.svg'
import dgIdle from './assets/dg/idle.svg'
import dgNod from './assets/dg/nod.svg'
import dgRest from './assets/dg/rest.svg'
import dgSpark from './assets/dg/spark.svg'
import dgThink from './assets/dg/think.svg'
import dgWave from './assets/dg/wave.svg'

export const DG_SUPPORTED_STAGES = [0, 1, 2]
export const DG_STAGE_VISUAL_POSES = ['idle', 'wave', 'spark']

const assetManifest = {
  collapsed: {
    default: dgCollapsedBadge,
  },
  idle: {
    default: dgIdle,
  },
  wave: {
    default: dgWave,
  },
  nod: {
    default: dgNod,
  },
  think: {
    default: dgThink,
  },
  rest: {
    default: dgRest,
  },
  spark: {
    default: dgSpark,
  },
}

function normalizePose(pose) {
  const normalized = String(pose || 'idle').replace(/^pose-/, '')
  if (assetManifest[normalized]) {
    return normalized
  }
  return 'idle'
}

function normalizeStage(stage) {
  const numeric = Number(stage)
  if (Number.isNaN(numeric) || numeric < 0) return 0
  if (numeric > 2) return 2
  return Math.floor(numeric)
}

export function resolveDGMascotAsset({ pose = 'idle', stage = 0, collapsed = false } = {}) {
  const normalizedStage = normalizeStage(stage)
  const stageKey = `stage${normalizedStage}`

  if (collapsed) {
    return assetManifest.collapsed[stageKey] || assetManifest.collapsed.default
  }

  const normalizedPose = normalizePose(pose)
  const poseAssets = assetManifest[normalizedPose] || assetManifest.idle
  return poseAssets[stageKey] || poseAssets.default || assetManifest.idle.default
}

export function resolveDGMascotStageClass(stage) {
  return `stage-${normalizeStage(stage)}`
}

export function dgStageVisualReady(pose) {
  return DG_STAGE_VISUAL_POSES.includes(normalizePose(pose))
}
