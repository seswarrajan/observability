package host

import (
	"strconv"
	"strings"

	"github.com/accuknox/observability/src/types"
	"github.com/accuknox/observability/utils/constants"
	"github.com/accuknox/observability/utils/database"
)

//FilterByHostName - To fetch the kubearmor aggregated logs based on Specific Host Name
func FilterByHostName(hostName []string, limit int) ([]types.KubeArmor, error) {

	var result []types.KubeArmor

	//query to fetch logs based on host name(s)
	query := constants.SELECT_HostName_KUBEARMOR
	if limit != 0 {
		//query to fetch logs based on host name(s) with limit
		query = query + constants.LIMIT + strconv.Itoa(limit)
	}

	//Fetch rows
	rows, err := database.ConnectDB().Query(query, strings.Join(hostName, ", "))
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var kubeArmor types.KubeArmor
		//Scan the record
		if err := rows.Scan(&kubeArmor.ClusterName, &kubeArmor.HostName,
			&kubeArmor.NamespaceName, &kubeArmor.PodName, &kubeArmor.ContainerID, &kubeArmor.ContainerName, &kubeArmor.Total); err != nil {
			return result, err
		}
		//append record
		result = append(result, kubeArmor)
	}
	return result, nil
}