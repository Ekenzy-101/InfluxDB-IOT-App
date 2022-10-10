# INFLUXDB IOT APP

This project provides a NextJS Web App and Golang server-side REST API that interacts with InfluxDB.
It is an adaptation of [InfluxData IoT Center](https://github.com/bonitoo-io/iot-center-v2).

## FEATURES

This application demonstrates how you can use InfluxDB to do the following:

- Create and manage InfluxDB authorizations (API tokens and permissions).
- Write and query device metadata in InfluxDB.
- Write and query telemetry data in InfluxDB.
- Generate data visualizations with the InfluxDB Giraffe library.

## SETUP

1. If you don't already have an InfluxDB instance, [create an InfluxDB Cloud account](https://www.influxdata.com/products/influxdb-cloud/) or [install InfluxDB OSS](https://www.influxdata.com/products/influxdb/).
2. Clone this repository to your machine.
3. Change to the API directory, enter the following code into the terminal:

   ```bash
   cd api
   ```

4. Install [Task](https://taskfile.dev/installation/) and [Golang](https://go.dev/doc/install).
5. Fill in the neccessary `.env` variables:
6. To start the API server, enter the following code into your terminal:

   ```bash
   task server
   ```

7. Change to the API directory, enter the following code into the terminal:

   ```bash
   cd api
   ```

8. Install [NodeJS](https://nodejs.org/) and [Yarn](https://yarnpkg.com/getting-started/install).

9. With `yarn` installed, enter the following code into your terminal to install the project dependencies:

   ```bash
   yarn
   ```

10. Add a `./.env.local` file that contains the following configuration variables:

    ```bash
    # Local environment secrets

    INFLUX_TOKEN=INFLUXDB_ALL_ACCESS_TOKEN
    INFLUX_ORG=INFLUXDB_ORG_ID
    ```

    Replace the following:

    - **`INFLUXDB_ALL_ACCESS_TOKEN`** with your InfluxDB **All Access** token.
    - **`INFLUXDB_ORG_ID`** with your InfluxDB organization ID.

11. To start the application in **development** mode, enter the following code into the terminal:

    ```bash
    yarn dev
    ```

12. In your browser, visit <http://localhost:3000/devices> to view the API output.
