package types

import (
	"fmt"
)

type ClipAnimationEnterEffect string

const (
	ClipAnimationEnterEffectNone        ClipAnimationEnterEffect = "none"
	ClipAnimationEnterEffectFade        ClipAnimationEnterEffect = "fade"
	ClipAnimationEnterEffectFloat       ClipAnimationEnterEffect = "float"
	ClipAnimationEnterEffectGentleFloat ClipAnimationEnterEffect = "gentle_float"
	ClipAnimationEnterEffectZoomIn      ClipAnimationEnterEffect = "zoom_in"
	ClipAnimationEnterEffectDrop        ClipAnimationEnterEffect = "drop"
	ClipAnimationEnterEffectSlideLeft   ClipAnimationEnterEffect = "slide_left"
	ClipAnimationEnterEffectSlideRight  ClipAnimationEnterEffect = "slide_right"
	ClipAnimationEnterEffectSlideUp     ClipAnimationEnterEffect = "slide_up"
	ClipAnimationEnterEffectSlideDown   ClipAnimationEnterEffect = "slide_down"
	ClipAnimationEnterEffectPop         ClipAnimationEnterEffect = "pop"
	ClipAnimationEnterEffectBounce      ClipAnimationEnterEffect = "bounce"
	ClipAnimationEnterEffectSpin        ClipAnimationEnterEffect = "spin"
	ClipAnimationEnterEffectSlideBounce ClipAnimationEnterEffect = "slide_bounce"
)

func NewClipAnimationEnterEffectFromString(s string) (ClipAnimationEnterEffect, error) {
	switch s {
	case "none":
		return ClipAnimationEnterEffectNone, nil
	case "fade":
		return ClipAnimationEnterEffectFade, nil
	case "float":
		return ClipAnimationEnterEffectFloat, nil
	case "gentle_float":
		return ClipAnimationEnterEffectGentleFloat, nil
	case "zoom_in":
		return ClipAnimationEnterEffectZoomIn, nil
	case "drop":
		return ClipAnimationEnterEffectDrop, nil
	case "slide_left":
		return ClipAnimationEnterEffectSlideLeft, nil
	case "slide_right":
		return ClipAnimationEnterEffectSlideRight, nil
	case "slide_up":
		return ClipAnimationEnterEffectSlideUp, nil
	case "slide_down":
		return ClipAnimationEnterEffectSlideDown, nil
	case "pop":
		return ClipAnimationEnterEffectPop, nil
	case "bounce":
		return ClipAnimationEnterEffectBounce, nil
	case "spin":
		return ClipAnimationEnterEffectSpin, nil
	case "slide_bounce":
		return ClipAnimationEnterEffectSlideBounce, nil
	}
	var t ClipAnimationEnterEffect
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ClipAnimationEnterEffect) Ptr() *ClipAnimationEnterEffect {
	return &e
}
