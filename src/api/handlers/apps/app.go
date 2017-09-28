// +-------------------------------------------------------------------------
// | Copyright (C) 2017 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package apps

import (
	"github.com/go-openapi/runtime/middleware"

	"apphub/src/api/swagger/restapi/operations"
	"apphub/src/api/swagger/restapi/operations/apps"
)

type AppsServiceHander interface {
	AppsGetAppsHandler(apps.GetAppsParams) middleware.Responder
	AppsPostAppsHandler(apps.PostAppsParams) middleware.Responder
	AppsGetAppsAppIDHandler(apps.GetAppsAppIDParams) middleware.Responder
	AppsDeleteAppsAppIDHandler(apps.DeleteAppsAppIDParams) middleware.Responder
}

func RegisterHanler(api *operations.AppHubAPI, hanlder AppsServiceHander) {
	api.AppsGetAppsHandler = apps.GetAppsHandlerFunc(hanlder.AppsGetAppsHandler)
	api.AppsPostAppsHandler = apps.PostAppsHandlerFunc(hanlder.AppsPostAppsHandler)

	api.AppsGetAppsAppIDHandler = apps.GetAppsAppIDHandlerFunc(hanlder.AppsGetAppsAppIDHandler)
	api.AppsDeleteAppsAppIDHandler = apps.DeleteAppsAppIDHandlerFunc(hanlder.AppsDeleteAppsAppIDHandler)
}
