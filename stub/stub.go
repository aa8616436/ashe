package stub

import (
	"github.com/agiledragon/gomonkey"
	"github.com/helloteemo/ashe/core"
)

// GoMonkey 打桩工具
type GoMonkey struct {
	*core.Core
	*gomonkey.Patches
}

func New(t *core.Core) *GoMonkey {
	return &GoMonkey{
		Core:    t,
		Patches: gomonkey.NewPatches(),
	}
}

func (m *GoMonkey) GetStub() *gomonkey.Patches {
	if m.Patches == nil {
		m.Patches = gomonkey.NewPatches()
	}
	return m.Patches
}
