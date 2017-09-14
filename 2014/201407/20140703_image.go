// The Image interface is simple with only 3 elements. Each one
// of those elements has its own details though. Fortunately,
// they aren't very complex either. Here they are.
type Model interface {
    // Model can convert any Color to one from its
    // own color model. The conversion may be lossy.
    Convert(c Color) Color
}

// Stores two (x,y) coordinates called Points
// Those points mark the two corners of the rectangle
type Rectangle struct {
    Min, Max Point
}

// Color stores an RGB value and an Alpha(transparency)
type Color interface {
    // RGBA returns the alpha-premultiplied red, green, blue and alpha values
    // for the color. Each value ranges within [0, 0xFFFF], but is represented
    // by a uint32 so that multiplying by a blend factor up to 0xFFFF will not
    // overflow.
    RGBA() (r, g, b, a uint32)
}