package send

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kansattica/mastodial/common"
)

func readQueue() ([]action, error) {
	var actions []action

	file, err := os.Open(common.QueueLocation)
	defer file.Close()

	if err != nil {
		if os.IsNotExist(err) {
			return actions, nil
		}
		return nil, fmt.Errorf("Could not open queue file %s. os.Open said %s.", common.QueueLocation, err)
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, fmt.Errorf("Could not open queue file. ReadAll reported " + err.Error())
	}

	if len(bytes) > 0 {
		err = json.Unmarshal(bytes, &actions)

		if err != nil {
			return nil, fmt.Errorf("Could not parse queue file. Unmarshal reported " + err.Error())
		}
	}

	return actions, nil
}

func saveQueue(actions []action) error {
	bytes, err := json.MarshalIndent(actions, "", "\t")

	if err != nil {
		fmt.Println("Failed to save queue to disk. Any queue you had before has not been changed. Marshal reported " + err.Error())
		return err
	}

	err = ioutil.WriteFile(common.QueueLocation, bytes, 0644)

	if err != nil {
		fmt.Println("Failed to write new queue file. WriteFile reported " + err.Error())
		return err
	}

	return nil

}
