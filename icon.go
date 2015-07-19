package gowi

import (
	"errors"
	"image"
	"path/filepath"
	"syscall"
	"unsafe"
)

import (
	"github.com/nvsoft/win"
)

// Icon is a bitmap that supports transparency and combining multiple
// variants of an image in different resolutions.
type Icon struct {
	hIcon   win.HICON
	isStock bool
}

func IconApplication() *Icon {
	return &Icon{win.LoadIcon(0, win.MAKEINTRESOURCE(win.IDI_APPLICATION)), true}
}

func IconError() *Icon {
	return &Icon{win.LoadIcon(0, win.MAKEINTRESOURCE(win.IDI_ERROR)), true}
}

func IconQuestion() *Icon {
	return &Icon{win.LoadIcon(0, win.MAKEINTRESOURCE(win.IDI_QUESTION)), true}
}

func IconWarning() *Icon {
	return &Icon{win.LoadIcon(0, win.MAKEINTRESOURCE(win.IDI_WARNING)), true}
}

func IconInformation() *Icon {
	return &Icon{win.LoadIcon(0, win.MAKEINTRESOURCE(win.IDI_INFORMATION)), true}
}

func IconWinLogo() *Icon {
	return &Icon{win.LoadIcon(0, win.MAKEINTRESOURCE(win.IDI_WINLOGO)), true}
}

func IconShield() *Icon {
	return &Icon{win.LoadIcon(0, win.MAKEINTRESOURCE(win.IDI_SHIELD)), true}
}

// NewIconFromFile returns a new Icon, using the specified icon image file.
func NewIconFromFile(filePath string) (*Icon, error) {
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}

	hIcon := win.HICON(win.LoadImage(
		0,
		syscall.StringToUTF16Ptr(absFilePath),
		win.IMAGE_ICON,
		0,
		0,
		win.LR_DEFAULTSIZE|win.LR_LOADFROMFILE))
	if hIcon == 0 {
		return nil, errors.New("LoadImage")
	}

	return &Icon{hIcon: hIcon}, nil
}

// NewIconFromResource returns a new Icon, using the specified icon resource.
func NewIconFromResource(resName string) (ic *Icon, err error) {
	hInst := win.GetModuleHandle(nil)
	if hInst == 0 {
		err = errors.New("GetModuleHandle")
		return
	}
	if hIcon := win.LoadIcon(hInst, syscall.StringToUTF16Ptr(resName)); hIcon == 0 {
		err = errors.New("LoadIcon")
	} else {
		ic = &Icon{hIcon: hIcon}
	}
	return
}

func NewIconFromImage(im image.Image) (ic *Icon, err error) {
	hIcon, err := createAlphaCursorOrIconFromImage(im, image.Pt(0, 0), true)
	if err != nil {
		return nil, err
	}
	return &Icon{hIcon: hIcon}, nil
}

// Dispose releases the operating system resources associated with the Icon.
func (i *Icon) Dispose() error {
	if i.isStock || i.hIcon == 0 {
		return nil
	}

	if !win.DestroyIcon(i.hIcon) {
		return errors.New("DestroyIcon")
	}

	i.hIcon = 0

	return nil
}

func (i *Icon) Destroy() bool {
	return win.DestroyIcon(i.hIcon)
}

func (i *Icon) Handle() win.HICON {
	return i.hIcon
}

// create an Alpha Icon or Cursor from an Image
// http://support.microsoft.com/kb/318876
func createAlphaCursorOrIconFromImage(im image.Image, hotspot image.Point, fIcon bool) (win.HICON, error) {
	hBitmap, err := hBitmapFromImage(im)
	if err != nil {
		return 0, err
	}
	defer win.DeleteObject(win.HGDIOBJ(hBitmap))

	// Create an empty mask bitmap.
	hMonoBitmap := win.CreateBitmap(int32(im.Bounds().Dx()), int32(im.Bounds().Dy()), 1, 1, nil)
	if hMonoBitmap == 0 {
		return 0, errors.New("CreateBitmap failed")
	}
	defer win.DeleteObject(win.HGDIOBJ(hMonoBitmap))

	var ii win.ICONINFO
	if fIcon {
		ii.FIcon = win.TRUE
	}
	ii.XHotspot = uint32(hotspot.X)
	ii.YHotspot = uint32(hotspot.Y)
	ii.HbmMask = hMonoBitmap
	ii.HbmColor = hBitmap

	// Create the alpha cursor with the alpha DIB section.
	hIconOrCursor := win.CreateIconIndirect(&ii)

	return hIconOrCursor, nil
}

func hBitmapFromImage(im image.Image) (win.HBITMAP, error) {
	var bi win.BITMAPV5HEADER
	bi.BiSize = uint32(unsafe.Sizeof(bi))
	bi.BiWidth = int32(im.Bounds().Dx())
	bi.BiHeight = -int32(im.Bounds().Dy())
	bi.BiPlanes = 1
	bi.BiBitCount = 32
	bi.BiCompression = win.BI_BITFIELDS
	// The following mask specification specifies a supported 32 BPP
	// alpha format for Windows XP.
	bi.BV4RedMask = 0x00FF0000
	bi.BV4GreenMask = 0x0000FF00
	bi.BV4BlueMask = 0x000000FF
	bi.BV4AlphaMask = 0xFF000000

	hdc := win.GetDC(0)
	defer win.ReleaseDC(0, hdc)

	var lpBits unsafe.Pointer

	// Create the DIB section with an alpha channel.
	hBitmap := win.CreateDIBSection(hdc, &bi.BITMAPINFOHEADER, win.DIB_RGB_COLORS, &lpBits, 0, 0)
	switch hBitmap {
	case 0, win.ERROR_INVALID_PARAMETER:
		return 0, errors.New("CreateDIBSection failed")
	}

	// Fill the image
	bitmap_array := (*[1 << 30]byte)(unsafe.Pointer(lpBits))
	i := 0
	for y := im.Bounds().Min.Y; y != im.Bounds().Max.Y; y++ {
		for x := im.Bounds().Min.X; x != im.Bounds().Max.X; x++ {
			r, g, b, a := im.At(x, y).RGBA()
			bitmap_array[i+3] = byte(a >> 8)
			bitmap_array[i+2] = byte(r >> 8)
			bitmap_array[i+1] = byte(g >> 8)
			bitmap_array[i+0] = byte(b >> 8)
			i += 4
		}
	}

	return hBitmap, nil
}
