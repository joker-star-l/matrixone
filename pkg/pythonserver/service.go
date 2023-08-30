// Copyright 2023 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pythonserver

import (
	"os"
	"os/exec"
	"path"
	"sync"
)

type service struct {
	cfg     Config
	process *os.Process
	mutex   sync.Mutex
}

func NewService(cfg Config) (PythonUdfServer, error) {
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}
	return &service{cfg: cfg}, nil
}

func (s *service) Start() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.process == nil {
		cmd := exec.Command(
			"/bin/bash", "-c",
			"python -u "+path.Join(s.cfg.Path, "server.py")+" --address="+s.cfg.Address+" > server.log 2>&1 &",
		)
		err := cmd.Run()
		if err != nil {
			return err
		}
		s.process = cmd.Process
	}
	return nil
}

func (s *service) Close() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.process != nil {
		err := s.process.Kill()
		if err != nil {
			return err
		}
		s.process = nil
	}
	return nil
}
