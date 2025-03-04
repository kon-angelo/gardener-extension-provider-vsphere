/*
 * Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *
 *  You may obtain a copy of the License at
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package utils

import (
	"github.com/go-logr/logr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	logzap "sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// ZapLogger is a Logger implementation.
// If development is true, a Zap development config will be used
// (stacktraces on warnings, no sampling), otherwise a Zap production
// config will be used (stacktraces on errors, sampling).
// Additionally, the time encoding is adjusted to `zapcore.ISO8601TimeEncoder`.
// This is used by extensions for historical reasons.
// TODO: consolidate this with NewZapLogger and make everything configurable in a harmonized way
func ZapLogger(development bool) logr.Logger {
	return logzap.New(func(o *logzap.Options) {
		var encCfg zapcore.EncoderConfig
		if development {
			encCfg = zap.NewDevelopmentEncoderConfig()
		} else {
			encCfg = zap.NewProductionEncoderConfig()
		}
		encCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encCfg.EncodeDuration = zapcore.StringDurationEncoder

		o.Encoder = zapcore.NewJSONEncoder(encCfg)
		o.Development = development
	})
}
