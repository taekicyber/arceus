/*
Copyright © 2021 zc2638 <zc2638@qq.com>.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package global

import (
	"github.com/pkgms/go/ctr"
	"github.com/sirupsen/logrus"
	"github.com/zc2638/arceus/pkg/util"
)

func Init() error {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
		FullTimestamp:          true,
		TimestampFormat:        "2006/01/02 15:04:05",
	})
	ctr.InitLog(logrus.StandardLogger())
	if err := util.MkdirAll(CustomResourcePath); err != nil {
		return err
	}
	if err := util.MkdirAll(TemplateResourcePath); err != nil {
		return err
	}
	return util.MkdirAll(RuleResourcePath)
}
