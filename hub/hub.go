package hub

import (
	"time"

	"github.com/wencode/fastkb/draw"
)

type Module interface {
	Name() string
	Init() error
	Update(time.Duration) error
	// Draw()方法
	draw.Drawing // 不直接引入Draw()方法, 原因是方便更换底层引擎
	// 接收到Notifaction
	OnNotify(ntf Notif, arg0, arg1 int, arg interface{})
}

type listenKey struct {
	modName string
	ntf     Notif
}

var (
	mods      = make([]Module, 0)
	listenMap = make(map[listenKey][]Module)
)

func Register(mod Module) {
	mods = append(mods, mod)
}

func Listen(mod Module, srcModName string, ntf Notif) {
	key := listenKey{srcModName, ntf}
	mods, ok := listenMap[key]
	if !ok {
		mods = make([]Module, 0)
	}
	listenMap[key] = append(mods, mod)
}

func Update(delta time.Duration) (errModName string, err error) {
	for _, mod := range mods {
		if moderr := mod.Update(delta); moderr != nil {
			if err != nil {
				errModName = mod.Name()
				err = moderr
			}
		}
	}
	return
}

// LightNotify用于无参数时Notify
func LightNotify(srcMod Module, ntf Notif) {
	Notify(srcMod, ntf, 0, 0, nil)
}

//
func Notify(srcMod Module, ntf Notif, arg0, arg1 int, arg interface{}) {
	NotifyingByName(srcMod.Name(), ntf, arg0, arg1, arg)
}

func NotifyingByName(srcModeName string, ntf Notif, arg0, arg1 int, arg interface{}) {
	key := listenKey{srcModeName, ntf}
	mods, ok := listenMap[key]
	if !ok {
		return
	}
	for _, mod := range mods {
		mod.OnNotify(ntf, arg0, arg1, arg)
	}
}

type ModuleFunc func(Module)

func ConcurrentExec(fn ModuleFunc) {
}

func SyncExec(fn ModuleFunc) {
	for _, mod := range mods {
		fn(mod)
	}
}
