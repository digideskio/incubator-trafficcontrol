/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package datareq

import (
	"net/http"
	"net/url"

	"github.com/apache/incubator-trafficcontrol/traffic_monitor_golang/traffic_monitor/cache"
	"github.com/apache/incubator-trafficcontrol/traffic_monitor_golang/traffic_monitor/peer"
	"github.com/apache/incubator-trafficcontrol/traffic_monitor_golang/traffic_monitor/threadsafe"
	todata "github.com/apache/incubator-trafficcontrol/traffic_monitor_golang/traffic_monitor/trafficopsdata"
)

func srvCacheStats(params url.Values, errorCount threadsafe.Uint, path string, toData todata.TODataThreadsafe, statResultHistory threadsafe.ResultStatHistory, statInfoHistory threadsafe.ResultInfoHistory, monitorConfig threadsafe.TrafficMonitorConfigMap, combinedStates peer.CRStatesThreadsafe, statMaxKbpses threadsafe.CacheKbpses) ([]byte, int) {
	filter, err := NewCacheStatFilter(path, params, toData.Get().ServerTypes)
	if err != nil {
		HandleErr(errorCount, path, err)
		return []byte(err.Error()), http.StatusBadRequest
	}
	bytes, err := cache.StatsMarshall(statResultHistory.Get(), statInfoHistory.Get(), combinedStates.Get(), monitorConfig.Get(), statMaxKbpses.Get(), filter, params)
	return WrapErrCode(errorCount, path, bytes, err)
}
