/*
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package calculations

import (
	"fmt"
	"magma/orc8r/cloud/go/services/analytics/protos"
	"magma/orc8r/cloud/go/services/analytics/query_api"

	"github.com/golang/glog"
)

//RawMetricsCalculation params for querying existing metrics
type RawMetricsCalculation struct {
	CalculationParams
	MetricExpr string
}

//Calculate queries for preexisting metric or provided promql expression and returns result
func (x *RawMetricsCalculation) Calculate(prometheusClient query_api.PrometheusAPI) ([]*protos.CalculationResult, error) {
	glog.V(10).Infof("Calculating Raw Metrics for %s", x.Name)
	vec, err := query_api.QueryPrometheusVector(prometheusClient, x.MetricExpr)
	if err != nil {
		return nil, fmt.Errorf("query error: %s", err)
	}
	results := MakeVectorResults(vec, x.Labels, x.Name)
	glog.V(10).Infof("Results %v", results)
	return results, nil
}
