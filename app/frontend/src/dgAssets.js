const assetModules = import.meta.glob('./assets/dg/*.{svg,png,webp}', {
  eager: true,
  import: 'default',
})

export const DG_SUPPORTED_STAGES = [0, 1, 2]
export const DG_STAGE_VISUAL_POSES = ['idle', 'wave', 'spark']

function assetFile(name) {
  return (
    assetModules[`./assets/dg/${name}.webp`] ||
    assetModules[`./assets/dg/${name}.png`] ||
    assetModules[`./assets/dg/${name}.svg`] ||
    null
  )
}

const baseAssets = {
  collapsed: assetFile('collapsed-badge'),
  idle: assetFile('idle'),
  wave: assetFile('wave'),
  nod: assetFile('nod'),
  think: assetFile('think'),
  rest: assetFile('rest'),
  spark: assetFile('spark'),
}

function normalizePose(pose) {
  const normalized = String(pose || 'idle').replace(/^pose-/, '')
  if (baseAssets[normalized]) {
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

function stageAssetName(pose, stage) {
  return `${pose}-stage${stage}`
}

export function resolveDGMascotAsset({ pose = 'idle', stage = 0, collapsed = false } = {}) {
  const normalizedStage = normalizeStage(stage)
  const normalizedPose = collapsed ? 'collapsed' : normalizePose(pose)

  if (collapsed) {
    return (
      assetFile(stageAssetName('collapsed-badge', normalizedStage)) ||
      baseAssets.collapsed ||
      baseAssets.idle
    )
  }

  return (
    assetFile(stageAssetName(normalizedPose, normalizedStage)) ||
    baseAssets[normalizedPose] ||
    baseAssets.idle
  )
}

export function resolveDGMascotStageClass(stage) {
  return `stage-${normalizeStage(stage)}`
}

export function dgStageVisualReady(pose) {
  return DG_STAGE_VISUAL_POSES.includes(normalizePose(pose))
}
