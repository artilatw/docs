# NPU Quick Start Guide

The **Matrix-800** is equipped with Arm Ethos-U65 microNPU for accelerated AI/ML inference. 

## Project Setup

```bash
// Create project
mkdir $PROJECT_NAME
cd $PROJECT_NAME

// Activate virtual environment
python3.12 -m venv .venv
source .venv/bin/activate

// Install libraries
pip install /opt/npu/wheels/tflite_runtime-2.18.0-cp312-cp312-linux_aarch64.whl

// Compile .tflite model using Vela to optimize model for execution on NPU
vela --accelerator-config ethos-u65-256 --output-dir . <MODEL>.tflite
```

> [!NOTE]
> Your .tflite model must be fully INT8 (input and output tensors, not just weights)

## Python Pipeline

```python
from tflite_runtime.interpreter import Interpreter, load_delegate

# Initialize NPU
delegate = load_delegate("/usr/local/lib/libethosu_delegate.so")
interpreter = Interpreter(model_path="<MODEL>.tflite", experimental_delegates=[delegate])
interpreter.allocate_tensors()

_input = interpreter.get_input_details()[0]
_output = interpreter.get_output_details()[0]

# quantize input → set_tensor → invoke → get_tensor → dequantize output
while True:
	npu_input = get_data()

	x = quantize(npu_input, _input)
	interpreter.set_tensor(_input["index"], x)
	interpreter.invoke()
	y = interpreter.get_tensor(_output["index"])
	npu_output = dequantize(y, _output)
	
	class_id = int(np.argmax(npu_output))
	confidence = float(npu_output[class_id])
```