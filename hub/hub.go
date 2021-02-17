package hub

import (
	"time"

	"github.com/wencode/fastkb/draw"
)

type Notify int

type Module interface {
	Name() string
	Init() error
	Update(time.Duration) error
	draw.Drawing
	OnNotify(ntf Notify, arg0, arg1 int, arg interface{})
}

type ModuleFunc func(Module)

type listenKey struct {
	modName string
	ntf     Notify
}

var (
	mods      = make([]Module, 0)
	listenMap = make(map[listenKey][]Module)
)

func Register(mod Module) {
	mods = append(mods, mod)
}

func Listen(mod Module, srcModName string, ntf Notify) {
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

func ConcurrentExec(fn ModuleFunc) {
}

func SyncExec(fn ModuleFunc) {
	for _, mod := range mods {
		fn(mod)
	}
}

func LightNotify(srcMod Module, ntf Notify) {
	Notifying(srcMod, ntf, 0, 0, nil)
}

func Notifying(srcMod Module, ntf Notify, arg0, arg1 int, arg interface{}) {
	NotifyingDelegate(srcMod.Name(), ntf, arg0, arg1, arg)
}

func NotifyingDelegate(srcModeName string, ntf Notify, arg0, arg1 int, arg interface{}) {
	key := listenKey{srcModeName, ntf}
	mods, ok := listenMap[key]
	if !ok {
		return
	}
	for _, mod := range mods {
		mod.OnNotify(ntf, arg0, arg1, arg)
	}
}
