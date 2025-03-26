package fileService

import (
	"errors"
	"github.com/lxn/win"
	"syscall"
	"unsafe"
	"strings"
)

// ClipboardOperation 定义剪贴板操作类型
type ClipboardOperation int

const (
	Copy ClipboardOperation = iota
	Cut
)

// 将文件路径列表转换为Windows剪贴板格式
func pathsToDropFiles(paths []string) []byte {
	// DROPFILES 结构
	type DROPFILES struct {
		pFiles uint32 // offset of file list
		pt     struct {
			x, y int32
		}
		fNC      int32
		fWide    int32 // TRUE if wide char
	}

	// 计算所需总大小
	size := uint32(unsafe.Sizeof(DROPFILES{}))
	for _, path := range paths {
		size += uint32(len(path)*2 + 2) // UTF16 size + null terminator
	}
	size += 2 // 额外的null terminator

	// 创建缓冲区
	buffer := make([]byte, size)

	// 填充DROPFILES结构
	df := (*DROPFILES)(unsafe.Pointer(&buffer[0]))
	df.pFiles = uint32(unsafe.Sizeof(DROPFILES{}))
	df.fWide = 1 // 使用宽字符

	// 填充文件路径
	offset := int(df.pFiles)
	for _, path := range paths {
		pathPtr := syscall.StringToUTF16(path)
		for i := 0; i < len(pathPtr)-1; i++ {
			*(*uint16)(unsafe.Pointer(&buffer[offset+i*2])) = pathPtr[i]
		}
		offset += len(pathPtr) * 2
	}

	return buffer
}

// SetClipboard 设置系统剪贴板内容
func SetClipboard(paths []string, operation ClipboardOperation) error {
	if len(paths) == 0 {
		return errors.New("路径列表不能为空")
	}

	// 规范化路径
	for i, path := range paths {
		paths[i] = strings.ReplaceAll(path, "/", "\\")
	}

	// 打开剪贴板
	if !win.OpenClipboard(0) {
		return errors.New("无法打开系统剪贴板")
	}
	defer win.CloseClipboard()

	// 清空剪贴板
	win.EmptyClipboard()

	// 准备DROPFILES数据
	dropFilesData := pathsToDropFiles(paths)

	// 分配全局内存
	hGlobal := win.GlobalAlloc(win.GMEM_MOVEABLE, uint32(len(dropFilesData)))
	if hGlobal == 0 {
		return errors.New("内存分配失败")
	}

	// 锁定内存
	lpGlobal := win.GlobalLock(hGlobal)
	if lpGlobal == nil {
		win.GlobalFree(hGlobal)
		return errors.New("内存锁定失败")
	}

	// 复制数据到全局内存
	copy((*[1 << 30]byte)(unsafe.Pointer(lpGlobal))[:len(dropFilesData)], dropFilesData)

	// 解锁内存
	win.GlobalUnlock(hGlobal)

	// 设置剪贴板数据
	format := uint32(win.CF_HDROP)
	if win.SetClipboardData(format, win.HANDLE(hGlobal)) == 0 {
		win.GlobalFree(hGlobal)
		return errors.New("设置剪贴板数据失败")
	}

	// 设置Preferred DropEffect
	if operation == Cut {
		effect := uint32(2) // DROPEFFECT_MOVE
		effectData := []byte{byte(effect), 0, 0, 0}

		hGlobalEffect := win.GlobalAlloc(win.GMEM_MOVEABLE, 4)
		if hGlobalEffect == 0 {
			return errors.New("效果内存分配失败")
		}

		lpGlobalEffect := win.GlobalLock(hGlobalEffect)
		if lpGlobalEffect == nil {
			win.GlobalFree(hGlobalEffect)
			return errors.New("效果内存锁定失败")
		}

		copy((*[4]byte)(unsafe.Pointer(lpGlobalEffect))[:], effectData)
		win.GlobalUnlock(hGlobalEffect)

		formatEffect := win.RegisterClipboardFormat(syscall.StringToUTF16Ptr("Preferred DropEffect"))
		if win.SetClipboardData(uint32(formatEffect), win.HANDLE(hGlobalEffect)) == 0 {
			win.GlobalFree(hGlobalEffect)
			return errors.New("设置剪贴板效果失败")
		}
	}

	return nil
}

// ClearClipboard 清空系统剪贴板
func ClearClipboard() error {
	if !win.OpenClipboard(0) {
		return errors.New("无法打开系统剪贴板")
	}
	defer win.CloseClipboard()
	
	win.EmptyClipboard()
	return nil
}
