# NPU Quick Start Guide

The **Matrix-800** is equipped with an Arm Ethos-U65 microNPU, enabling efficient on-device AI/ML inference with significantly lower latency and power consumption compared to CPU-only execution.

## Performance Comparison 

Benchmark model: [ssd_mobilenet_v2_coco_quant_postprocess_vela.tflite](https://github.com/pogJames/camera-demo/blob/basic/tflite_model/ssd_mobilenet_v2_coco_quant_postprocess.tflite)

| Metric | CPU | NPU | Improvement |
| --- | --- | --- | --- |
| Avg inference time | 199,791 µs | 11,605.1 µs | ~17.2x faster |
| FPS (1e6/avg) | ~5.0 FPS | ~86.2 FPS | ~17.2x |
| Min - Max time | 199.451ms - 200.067ms | 11.490ms - 11.982ms | ~17.4x faster (min), ~16.7x faster (max) |
| Model size | 6.221 MB | 4.579 MB | ~26% smaller |

## Project Setup

On-device inference is commonly performed using [TensorFlow Lite](https://www.tensorflow.org/lite/guide) on Python:

```bash
// Create project
mkdir $PROJECT_NAME
cd $PROJECT_NAME

// Create and activate virtual environment
python3.12 -m venv .venv
source .venv/bin/activate

// Install TFLite library
pip install /opt/npu/wheels/tflite_runtime-2.18.0-cp312-cp312-linux_aarch64.whl
```

> [!NOTE]
> TFLite has recently been rebranded as LiteRT, but libethosu_delegate.so was strictly built against TFLite 2.18.0 and Python 3.12.
> > Notice how we used `python3.12` when building `venv`

## Vela Compiler

All models must be compiled with Arm Vela before they can run on the Ethos-U65 NPU. Using an unoptimized model will result in an error:

```bash
// Compile .tflite model using Vela to optimize model for execution on NPU
vela --accelerator-config ethos-u65-256 --output-dir . <MODEL>.tflite
```

> [!NOTE]
> Your .tflite model must be fully INT8 quantized (both weights and activations, including input and output tensors).

## Python Pipeline

```python
from tflite_runtime.interpreter import Interpreter, load_delegate

# Initialize NPU
delegate = load_delegate("/usr/local/lib/libethosu_delegate.so")
interpreter = Interpreter(model_path="<MODEL>.tflite", experimental_delegates=[delegate])
interpreter.allocate_tensors()

# Get input/output tensor details
_input = interpreter.get_input_details()[0]
_output = interpreter.get_output_details()[0]

# Inference pipeline HERE
while True:
	# 1. Get input data
    input_data = get_data()
    # 2. Quantize input
    quantized_input = quantize(input_data, input_details)
    # 3. Set input tensor
    interpreter.set_tensor(input_details["index"], quantized_input)
    # 4. Run inference
    interpreter.invoke()
    # 5. Get and dequantize output
    quantized_output = interpreter.get_tensor(output_details["index"])
    output = dequantize(quantized_output, output_details)
    # 6. Post-process
    class_id = int(np.argmax(output))
    confidence = float(output[class_id])
```
