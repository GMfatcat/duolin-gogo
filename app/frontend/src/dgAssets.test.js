import { describe, expect, it } from 'vitest'
import {
  dgStageVisualReady,
  resolveDGMascotAsset,
  resolveDGMascotStageClass,
} from './dgAssets'

describe('dgAssets', () => {
  it('falls back to base pose art when no stage asset exists yet', () => {
    const idleBase = resolveDGMascotAsset({ pose: 'idle', stage: 0 })

    expect(resolveDGMascotAsset({ pose: 'idle', stage: 2 })).toBe(idleBase)
    expect(resolveDGMascotAsset({ pose: 'spark', stage: 1 })).not.toBe('')
  })

  it('always uses the collapsed badge while collapsed', () => {
    const collapsed = resolveDGMascotAsset({ collapsed: true, pose: 'spark', stage: 2 })
    expect(collapsed).toBe(resolveDGMascotAsset({ collapsed: true, pose: 'idle', stage: 0 }))
  })

  it('clamps stages into the supported 0-2 range', () => {
    expect(resolveDGMascotStageClass(-2)).toBe('stage-0')
    expect(resolveDGMascotStageClass(1)).toBe('stage-1')
    expect(resolveDGMascotStageClass(99)).toBe('stage-2')
  })

  it('marks idle wave and spark as the first stage-visual poses', () => {
    expect(dgStageVisualReady('idle')).toBe(true)
    expect(dgStageVisualReady('wave')).toBe(true)
    expect(dgStageVisualReady('spark')).toBe(true)
    expect(dgStageVisualReady('rest')).toBe(false)
  })
})
