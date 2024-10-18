# Motion Sense

## Prerequisites

1. Put your model file (in `.onnx` format) into the `ML` folder.

## Configuration

2. Set the model name in `configs/config.yaml`:

   ```yaml
   model_name: dummy_model.onnx
   ```

   Replace `dummy_model.onnx` with the appropriate model name if different.

## Installation

3. At the root of the project, install the necessary dependencies:

   ```sh
   make install
   ```

## Running the Project

4. To run the project:

   ```sh
   make run
   ```

## Running Tests

5. To run tests:

   ```sh
   make test
   ```
