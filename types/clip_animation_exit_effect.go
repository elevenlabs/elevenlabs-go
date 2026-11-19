package types

import (
	"fmt"
)

type ClipAnimationExitEffect string

const (
	ClipAnimationExitEffectNone        ClipAnimationExitEffect = "none"
	ClipAnimationExitEffectFade        ClipAnimationExitEffect = "fade"
	ClipAnimationExitEffectFloat       ClipAnimationExitEffect = "float"
	ClipAnimationExitEffectGentleFloat ClipAnimationExitEffect = "gentle_float"
	ClipAnimationExitEffectZoomIn      ClipAnimationExitEffect = "zoom_in"
	ClipAnimationExitEffectDrop        ClipAnimationExitEffect = "drop"
	ClipAnimationExitEffectSlideLeft   ClipAnimationExitEffect = "slide_left"
	ClipAnimationExitEffectSlideRight  ClipAnimationExitEffect = "slide_right"
	ClipAnimationExitEffectSlideUp     ClipAnimationExitEffect = "slide_up"
	ClipAnimationExitEffectSlideDown   ClipAnimationExitEffect = "slide_down"
	ClipAnimationExitEffectPop         ClipAnimationExitEffect = "pop"
	ClipAnimationExitEffectBounce      ClipAnimationExitEffect = "bounce"
	ClipAnimationExitEffectSpin        ClipAnimationExitEffect = "spin"
	ClipAnimationExitEffectSlideBounce ClipAnimationExitEffect = "slide_bounce"
)

func NewClipAnimationExitEffectFromString(s string) (ClipAnimationExitEffect, error) {
	switch s {
	case "none":
		return ClipAnimationExitEffectNone, nil
	case "fade":
		return ClipAnimationExitEffectFade, nil
	case "float":
		return ClipAnimationExitEffectFloat, nil
	case "gentle_float":
		return ClipAnimationExitEffectGentleFloat, nil
	case "zoom_in":
		return ClipAnimationExitEffectZoomIn, nil
	case "drop":
		return ClipAnimationExitEffectDrop, nil
	case "slide_left":
		return ClipAnimationExitEffectSlideLeft, nil
	case "slide_right":
		return ClipAnimationExitEffectSlideRight, nil
	case "slide_up":
		return ClipAnimationExitEffectSlideUp, nil
	case "slide_down":
		return ClipAnimationExitEffectSlideDown, nil
	case "pop":
		return ClipAnimationExitEffectPop, nil
	case "bounce":
		return ClipAnimationExitEffectBounce, nil
	case "spin":
		return ClipAnimationExitEffectSpin, nil
	case "slide_bounce":
		return ClipAnimationExitEffectSlideBounce, nil
	}
	var t ClipAnimationExitEffect
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ClipAnimationExitEffect) Ptr() *ClipAnimationExitEffect {
	return &e
}
