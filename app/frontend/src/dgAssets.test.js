import { describe, expect, it } from 'vitest'
import {
  dgStageVisualReady,
  resolveDGMascotAsset,
  resolveDGMascotStageClass,
} from './dgAssets'

describe('dgAssets', () => {
  it('resolves stage-specific assets when they exist', () => {
    expect(resolveDGMascotAsset({ pose: 'idle', stage: 0 })).toContain('idle-stage0')
    expect(resolveDGMascotAsset({ pose: 'idle', stage: 1 })).toContain('idle-stage1')
    expect(resolveDGMascotAsset({ pose: 'idle', stage: 2 })).toContain('idle-stage2')
    expect(resolveDGMascotAsset({ pose: 'wave', stage: 1 })).toContain('wave-stage1')
    expect(resolveDGMascotAsset({ pose: 'spark', stage: 2 })).toContain('spark-stage2')
  })

  it('falls back to the base pose when a staged variant does not exist for that pose', () => {
    expect(resolveDGMascotAsset({ pose: 'think', stage: 2 })).toContain('/think.')
    expect(resolveDGMascotAsset({ pose: 'rest', stage: 1 })).toContain('/rest.')
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
