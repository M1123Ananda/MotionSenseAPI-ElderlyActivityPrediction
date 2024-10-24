# Motion Sense

## Setting Up TorchServe with Docker

Before proceeding with the installation, you need to set up Docker and create a TorchServe service. Follow these steps:

1. Ensure Docker is installed and running on your system.
2. Place your model file in the `/model-server` directory.
3. Run the following command to create a TorchServe service:

   ```sh
   docker run --rm -p 8555:8080 -p 8081:8081 --name ts \
      -v \$(pwd)/model-server:/home/model-server/model-store \
      pytorch/torchserve:latest \
      torchserve --start --model-store /home/model-server/model-store --models model-name=model-name.mar --disable-token-auth
   ```

    - You can change the default port `8555` in the `config.yaml` file if needed.
    - Replace `model-name` with the desired name for your model.

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

