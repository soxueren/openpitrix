// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package frontgate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"

	"openpitrix.io/openpitrix/pkg/pb/types"
)

type ConfigManager struct {
	path string
	cfg  *pbtypes.FrontgateConfig
	mu   sync.Mutex
}

func NewConfigManager(path string, cfg *pbtypes.FrontgateConfig, opts ...Options) *ConfigManager {
	if cfg != nil {
		cfg = proto.Clone(cfg).(*pbtypes.FrontgateConfig)
	} else {
		cfg = NewDefaultConfig()
	}

	for _, fn := range opts {
		fn(cfg)
	}

	return &ConfigManager{
		path: path,
		cfg:  cfg,
	}
}

func (p *ConfigManager) Get() (cfg *pbtypes.FrontgateConfig) {
	p.mu.Lock()
	defer p.mu.Unlock()

	cfg = proto.Clone(p.cfg).(*pbtypes.FrontgateConfig)
	return
}

func (p *ConfigManager) Set(cfg *pbtypes.FrontgateConfig) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if cfg.Id != "" && cfg.Id != p.cfg.Id {
		return fmt.Errorf("frontgate: config.Id is read only")
	}
	if cfg.ListenPort > 0 && cfg.ListenPort != p.cfg.ListenPort {
		return fmt.Errorf("frontgate: config.ListenPort is read only")
	}

	if cfg.PilotHost != "" && cfg.PilotHost != p.cfg.PilotHost {
		return fmt.Errorf("frontgate: config.PilotHost is read only")
	}
	if cfg.PilotPort > 0 && cfg.PilotPort != p.cfg.PilotPort {
		return fmt.Errorf("frontgate: config.PilotPort is read only")
	}

	cfg.Id = p.cfg.Id
	cfg.ListenPort = p.cfg.ListenPort

	cfg.PilotHost = p.cfg.PilotHost
	cfg.PilotPort = p.cfg.PilotPort

	p.cfg = proto.Clone(cfg).(*pbtypes.FrontgateConfig)
	return nil
}

func (p *ConfigManager) Save() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	data, err := json.MarshalIndent(p.cfg, "", "\t")
	if err != nil {
		return err
	}

	// backup old config
	bakpath := p.path + time.Now().Format(".20060102.bak")
	if err := os.Rename(p.path, bakpath); err != nil {
		return err
	}

	data = bytes.Replace(data, []byte("\n"), []byte("\r\n"), -1)
	err = ioutil.WriteFile(p.path, data, 0666)
	if err != nil {
		os.Rename(bakpath, p.path) // revert
		return err
	}

	return nil
}
