import { TaskInputs, TaskOutputs } from "mesg-js/lib/service"
import { Marketplace } from "../contracts/Marketplace"
import { getAllServices } from "../contracts/service";
import expectOutput from "../../expected-output-list-services.json"


export default () => async (inputs: TaskInputs, outputs: TaskOutputs): Promise<void> => {
  try {
    outputs.success(expectOutput)
  }
  catch (error) {
    console.error('error im listServices', error)
    outputs.error({ message: error.toString() })
  }
}
